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

type StatusTracker struct {
	saver  Saver
	status map[string]dag.JobStatus
}

func NewStatusTracker(saver Saver) *StatusTracker {
	return &StatusTracker{
		saver:  saver,
		status: make(map[string]dag.JobStatus),
	}
}

func (this *StatusTracker) GetStatus(job *dag.Job) (dag.JobStatus, error) {
	return this.saver.GetFlag(job)
}

func (this *StatusTracker) SetStatus(job *dag.Job, status dag.JobStatus) error {
	err := this.saver.SetFlag(job, status)
	if err == nil {
		this.status[job.Name] = status
	}
	return err
}
