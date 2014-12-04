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
	"sync"
	"time"
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
	if err := ret.db.Open(); err != nil {
		return nil, err
	}
	ret.exec = map[string]exec.Exec{
		"odps": exec.NewODPSExec(),
	}
	for _, e := range ret.exec {
		if err := e.Setup(config.Env); err != nil {
			return nil, err
		}
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
	this.scanStep(this.flow.Entry)
	for len(this.todo) != 0 {
		var wg sync.WaitGroup
		for _, job := range this.todo {
			e, ok := this.exec[job.Type]
			if !ok {
				log.Fatalf("no exec for %s", job.Type)
				continue
			}
			wg.Add(1)
			go func(j *ast.Job, e exec.Exec) {
				defer wg.Done()
				/*
					status, err := e.Run(job)
					if err != nil {
						log.Fatalf("job failed: %v", err)
						job.Status = ast.FAIL
					}
					job.Status = status
				*/
				time.Sleep(10000)
				status := "done"
				log.Debugf("status: %s", status)
			}(job, e)
		}
		wg.Wait()
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
