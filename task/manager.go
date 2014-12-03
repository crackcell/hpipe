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
	"../log"
	"../yafl/ast"
	"./exec"
	_ "fmt"
)

//===================================================================
// Public APIs
//===================================================================

type TaskManager interface {
	Setup(flow *ast.Flow) error
	Run() error
}

func NewTaskManager() TaskManager {
	ret := new(taskMgr)
	ret.executors = make(map[string]exec.Executor)
	return ret
}

//===================================================================
// Private
//===================================================================

type taskMgr struct {
	flow      *ast.Flow
	executors map[string]exec.Executor
}

func (this *taskMgr) Setup(flow *ast.Flow) error {
	this.flow = flow
	return nil
}

func (this *taskMgr) Run() error {
	log.Debug("taskmgr runs")
	this.runStep(this.flow.Entry)
	return nil
}

func (this *taskMgr) runStep(s *ast.Step) error {
	log.Debugf("run step: %s", s.Name)
	for _, dep := range s.Dep {
		this.runStep(dep)
	}
	for _, do := range s.Do {
		this.runJob(do)
	}
	return nil
}

func (this *taskMgr) runJob(j *ast.Job) error {
	log.Debugf("run job: %s", j.Name)
	return nil
}
