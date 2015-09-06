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

package exec

import (
	"fmt"
	"github.com/crackcell/hpipe/dag"
	"github.com/crackcell/hpipe/exec/filesystem"
)

//===================================================================
// Public APIs
//===================================================================

type StatusKeeper struct {
	fs filesystem.Filesystem
}

func NewStatusKeeper(fs filesystem.Filesystem) *StatusKeeper {
	return &StatusKeeper{
		fs: fs,
	}
}

func (this *StatusKeeper) GetStatus(job *dag.Job) (dag.JobStatus, error) {
	output := job.Attrs["output"]
	for status, flag := range jobStatusFlags {
		if exist, err := this.fs.IsExist(output + flag); err != nil {
			return dag.UnknownStatus, err
		} else if exist {
			return status, err
		}
	}
	return dag.NotStarted, nil
}

func (this *StatusKeeper) SetStatus(job *dag.Job, status dag.JobStatus) error {
	output := job.Attrs["output"]
	if flag, ok := jobStatusFlags[status]; ok {
		return this.fs.Touch(output + flag)
	} else {
		return fmt.Errorf("invalid status: %s", status)
	}
}

func (this *StatusKeeper) DeleteStatus(job *dag.Job, status dag.JobStatus) error {
	output := job.Attrs["output"]
	if flag, ok := jobStatusFlags[status]; ok {
		return this.fs.Rm(output + flag)
	} else {
		return fmt.Errorf("invalid status: %s", status)
	}
}

func (this *StatusKeeper) ClearStatus(job *dag.Job) error {
	output := job.Attrs["output"]
	for _, flag := range jobStatusFlags {
		if err := this.fs.Rm(output + flag); err != nil {
			return err
		}
	}
	return nil
}

//===================================================================
// Private
//===================================================================

var jobStatusFlags = map[dag.JobStatus]string{
	dag.Started:  ".hpipe.started",
	dag.Finished: ".hpipe.finished",
	dag.Failed:   ".hpipe.failed",
}
