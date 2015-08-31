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
 * @file exec.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Tue Aug 25 18:09:11 2015
 *
 **/

package exec

import (
	"github.com/crackcell/hpipe/dag"
)

//===================================================================
// JobExec
//===================================================================

type Exec interface {
	Setup() error
	Run(job *dag.Job) error
	GetStatus(job *dag.Job) (dag.JobStatus, error)
}

//===================================================================
// Private
//===================================================================
