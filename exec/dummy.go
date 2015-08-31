/***************************************************************
 *
 * Copyright (c) 2015, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 *
 *
 * @file hadoop.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Tue Aug 25 18:28:05 2015
 *
 **/

package exec

import (
	//"fmt"
	//"github.com/crackcell/hpipe/config"
	"github.com/crackcell/hpipe/dag"
)

//===================================================================
// Public APIs
//===================================================================

type DummyExec struct {
}

func NewDummyExec() *DummyExec {
	return &DummyExec{}
}

func (this *DummyExec) Setup() error {
	return nil
}

func (this *DummyExec) Run(job *dag.Job) error {
	return nil
}

func (this *DummyExec) GetStatus(job *dag.Job) (dag.JobStatus, error) {
	return dag.UnknownStatus, nil
}

//===================================================================
// Private
//===================================================================
