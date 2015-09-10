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
 * @file file.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sun Sep  6 23:26:06 2015
 *
 **/

package status

import (
	"fmt"
	"github.com/crackcell/hpipe/dag"
	"github.com/crackcell/hpipe/storage"
)

//===================================================================
// Public APIs
//===================================================================

type FileKeeper struct {
	fs storage.Storage
}

func NewFileKeeper(fs storage.Storage) *FileKeeper {
	return &FileKeeper{
		fs: fs,
	}
}

func (this *FileKeeper) GetStatus(job *dag.Job) (dag.JobStatus, error) {
	output := job.Attrs["output"]
	for status, flag := range JobStatusFlags {
		if exist, err := this.fs.IsExist(output + flag); err != nil {
			return dag.UnknownStatus, err
		} else if exist {
			return status, err
		}
	}
	return dag.NotStarted, nil
}

func (this *FileKeeper) SetStatus(job *dag.Job, status dag.JobStatus) error {
	if err := this.ClearStatus(job); err == nil {
		return err
	}
	output := job.Attrs["output"]
	if flag, ok := JobStatusFlags[status]; ok {
		return this.fs.Touch(output + flag)
	} else {
		return fmt.Errorf("invalid status: %s", status)
	}
}

func (this *FileKeeper) DeleteStatus(job *dag.Job, status dag.JobStatus) error {
	output := job.Attrs["output"]
	if flag, ok := JobStatusFlags[status]; ok {
		return this.fs.Rm(output + flag)
	} else {
		return fmt.Errorf("invalid status: %s", status)
	}
}

func (this *FileKeeper) ClearStatus(job *dag.Job) error {
	output := job.Attrs["output"]
	for _, flag := range JobStatusFlags {
		if err := this.fs.Rm(output + flag); err != nil {
			return err
		}
	}
	return nil
}

//===================================================================
// Private
//===================================================================
