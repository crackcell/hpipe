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
	_ "fmt"
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
	executors map[string]exec.Executor
	db        storage.DB
	todo      []*ast.Job
}

func (this *taskManager) Setup(flow *ast.Flow) error {
	this.flow = flow
	return nil
}

func (this *taskManager) Run() error {
	log.Debug("start to run")

	this.scanStep(this.flow.Entry)
	for len(this.todo) != 0 {
		var wg sync.WaitGroup

		retchan := make(chan string, len(this.todo))
		for j := range this.todo {
			// TODO
		}
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
		this.todo = append(this.todo, j)
		log.Debugf("ready to go: %s", j.Name)
	}
}
