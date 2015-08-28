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
 * @file shell.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Tue Aug 25 18:32:30 2015
 *
 **/

package exec

import (
	//"fmt"
	"github.com/crackcell/hpipe/dag"
)

//===================================================================
// Public APIs
//===================================================================

type ShellExec struct {
}

func NewShellExec() *ShellExec {
	return &ShellExec{}
}

func (this *ShellExec) Setup() error {
	return nil
}

func (this *ShellExec) Run(job *dag.Job) error {
	return nil
}

func (this *ShellExec) GetJobStatus(job *dag.Job) dag.JobStatus {
	// TODO
	return dag.Finished
}

func (this *ShellExec) CheckJobAttrs(job *dag.Job) bool {
	return true
}

//===================================================================
// Private
//===================================================================
