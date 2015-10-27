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
 * @file flag.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sun Sep  6 23:26:06 2015
 *
 **/

package status

import (
	"github.com/crackcell/hpipe/dag"
)

//===================================================================
// Public APIs
//===================================================================

type Saver interface {
	GetFlag(job *dag.Job) (dag.JobStatus, error)
	SetFlag(job *dag.Job) error
	DeleteFlag(job *dag.Job, status dag.JobStatus) error
	ClearFlag(job *dag.Job) error
}

var FlagSuffix = map[dag.JobStatus]string{
	dag.Started:  ".hpipe.started",
	dag.Finished: ".hpipe.finished",
	dag.Failed:   ".hpipe.failed",
}
