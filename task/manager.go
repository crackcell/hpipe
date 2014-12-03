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
)

//===================================================================
// Public APIs
//===================================================================

type TaskManager interface {
	Setup(flow *ast.Flow) error
	Run() error
}

func NewTaskManager() (TaskManager, error) {
	ret := new(taskManager)
	ret.executors = make(map[string]exec.Executor)
	ret.db = storage.NewSqliteDB(config.MetaPath + "/meta.db")
	if err := ret.db.Open(); err != nil {
		return nil, err
	}
	return ret, nil
}

//===================================================================
// Private
//===================================================================

type taskManager struct {
	flow      *ast.Flow
	executors map[string]exec.Executor // type:Exector
	db        storage.DB
	todo      []*ast.Job
}

func (this *taskManager) Setup(flow *ast.Flow) error {
	this.flow = flow
	return nil
}

func (this *taskManager) Run() error {
	log.Debug("taskmgr runs")

	f, err := this.db.RestoreFlow(this.flow)
	if err != nil {
		return err
	}
	this.flow = f

	exe := exec.NewODPSExecutor()

	this.scanStep(this.flow.Entry)
	for len(this.todo) != 0 {
		var wg sync.WaitGroup
		for _, j := range this.todo {
			switch {
			case j.Types == "odps":

			default:
				log.Fatal("unknown job type")
				return fmt.Errorf("unknown job type")
			}
		}
	}

	if err := this.db.SaveFlow(this.flow); err != nil {
		return err
	}
	return nil
}

func (this *taskManager) scanStep(s *ast.Step) error {
	log.Debugf("run step: %s", s.Name)
	for _, dep := range s.Dep {
		this.scanStep(dep)
	}
	for _, do := range s.Do {
		this.scanJob(do)
	}
	return nil
}

func (this *taskManager) scanJob(j *ast.Job) error {
	//log.Debugf("run job: %s", j.Name)
	if j.Status == "todo" {
		this.todo = append(this.todo, j)
		log.Debugf("ready to run: %s", j.Name)
	}
	return nil
}
