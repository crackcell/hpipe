/***************************************************************
 *
 * Copyright (c) 2014, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 *
 *
 * @file manager.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Thu Nov 27 09:23:19 2014
 *
 **/

package task

import (
	"../etc"
	"../hpipe"
	"../log"
	"../yafl/ast"
	"./exec"
	"./storage"
	"fmt"
)

//===================================================================
// Public APIs
//===================================================================

type TaskManager interface {
	Run(flow *ast.Flow) error
}

func NewTaskManager() (TaskManager, error) {
	ret := new(taskManager)
	ret.trys = make(map[string]int)
	ret.todo = make(map[string]*ast.Job)
	ret.db = storage.NewSqliteDB(etc.MetaPath + "/meta.db")
	ret.exec = map[string]exec.Exec{
		"odps":   exec.NewOdpsExec(),
		"hadoop": exec.NewHadoopExec(),
	}
	return ret, nil
}

//===================================================================
// Private
//===================================================================

type taskManager struct {
	flow *ast.Flow
	exec map[string]exec.Exec
	db   storage.DB
	trys map[string]int
	todo map[string]*ast.Job
}

func (this *taskManager) Run(flow *ast.Flow) error {
	this.flow = flow

	if err := this.db.Open(); err != nil {
		return err
	}
	defer this.db.Close()

	if err := this.db.RestoreFlow(this.flow); err != nil {
		return err
	}

	this.todo = make(map[string]*ast.Job)
	this.scanStep(this.flow.Entry)
	for len(this.todo) != 0 {
		jobCount := 0
		ch := make(chan []string)
		for _, job := range this.todo {
			e, ok := this.exec[job.Type]
			if !ok {
				log.Fatalf("no exec for %s", job.Type)
				continue
			}

			jobCount++

			if trys, ok := this.trys[job.InstanceID]; ok {
				this.trys[job.InstanceID] = trys + 1
			} // no need to put init value which is done in scanjob

			go func(j *ast.Job, e exec.Exec) {
				status, err := e.Run(j)
				if err != nil {
					log.Fatalf("job failed: %v", err)
					j.Status = hpipe.FAIL
				}
				ch <- []string{j.Name, status}
			}(job, e)
		}

		for jobCount > 0 {
			select {
			case ret, ok := <-ch:
				if len(ret) != 2 {
					panic(fmt.Errorf("invalid return"))
				}
				name := ret[0]
				status := ret[1]
				job, ok := this.todo[name]
				if !ok {
					panic(fmt.Errorf("no job"))
				}
				job.Status = status
				jobCount--
			}
		}

		this.db.SaveFlow(this.flow)
		this.todo = make(map[string]*ast.Job)
		this.scanStep(this.flow.Entry)
	}
	return nil
}

func (this *taskManager) scanStep(s *ast.Step) {
	for _, dep := range s.Dep {
		this.scanStep(dep)
	}
	for _, do := range s.Do {
		this.scanJob(do)
	}
}

func (this *taskManager) scanJob(j *ast.Job) {
	trys, ok := this.trys[j.InstanceID]
	if ok {
		if trys >= int(etc.MaxRetry) {
			log.Fatalf("<%s> reaches max retrys %v",
				j.InstanceID, etc.MaxRetry)
			return
		}
	} else {
		this.trys[j.InstanceID] = 0
		trys = 0
	}

	switch {
	case j.Status == "todo" || j.Status == "fail" ||
		(j.Status == "done" && etc.Rerun && trys == 0):
		this.todo[j.Name] = j
	}
}
