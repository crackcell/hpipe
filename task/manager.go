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
	"../config"
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
	ret.todo = make(map[string]*ast.Job)
	ret.db = storage.NewSqliteDB(config.MetaPath + "/meta.db")
	ret.exec = map[string]exec.Exec{
		"odps": exec.NewODPSExec(),
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
	todo map[string]*ast.Job
}

func (this *taskManager) Run(flow *ast.Flow) error {
	log.Debug("start to run")

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
		//log.Debugf("-------------------> todo len: %d", len(this.todo))
		for k, v := range this.todo {
			log.Debugf("todo: %s - %v", k, v)
		}

		jobCount := 0
		ch := make(chan []string)
		for _, job := range this.todo {
			//log.Debugf("------------> run: %s", job.Name)
			e, ok := this.exec[job.Type]
			if !ok {
				log.Fatalf("no exec for %s", job.Type)
				continue
			}

			jobCount++

			go func(j *ast.Job, e exec.Exec) {
				//log.Debugf("------------> run in go: %s", j.Name)
				status, err := e.Run(j)
				if err != nil {
					log.Fatalf("job failed: %v", err)
					j.Status = ast.FAIL
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
				job, ok := this.todo[ret[0]]
				if !ok {
					panic(fmt.Errorf("no job"))
				}
				job.Status = ret[1]
				jobCount--
				//log.Debugf("%s <--- %s", job.Name, job.Status)
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
	if j.Status == "todo" {
		this.todo[j.Name] = j
		log.Debugf("ready to go: %s", j.Name)
	}
}
