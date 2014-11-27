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
 * @file taskmgr.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Thu Nov 27 09:23:19 2014
 *
 **/

package taskmanager

import (
	"../log"
	"../yafl/ast"
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
	return new(taskMgr)
}

//===================================================================
// Private
//===================================================================

type taskMgr struct {
	flow *ast.Flow
}

func (this *taskMgr) Setup(flow *ast.Flow) error {
	this.flow = flow
	return nil
}

func (this *taskMgr) Run() error {
	log.Debug("taskmgr runs")
	return nil
}
