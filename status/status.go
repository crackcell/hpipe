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
 * @file status.go
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

type StatusKeeper interface {
	GetStatus(job *dag.Job) (dag.JobStatus, error)
	SetStatus(job *dag.Job, status dag.JobStatus) error
	DeleteStatus(job *dag.Job, status dag.JobStatus) error
	ClearStatus(job *dag.Job) error
}

var JobStatusFlags = map[dag.JobStatus]string{
	dag.Started:  ".hpipe.started",
	dag.Finished: ".hpipe.finished",
	dag.Failed:   ".hpipe.failed",
}
