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
 * @file filesaver.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sun Sep  6 23:26:06 2015
 *
 **/

package status

import (
	"fmt"
	"github.com/crackcell/hpipe/dag"
	"github.com/crackcell/hpipe/storage"
	"sync"
)

//===================================================================
// Public APIs
//===================================================================

type FileSaver struct {
	fs   storage.Storage
	lock sync.Mutex
}

func NewHDFSSaver(namenode string) (*FileSaver, error) {
	if fs, err := storage.NewHDFS(namenode); err != nil {
		return nil, err
	} else {
		return NewFileSaver(fs), nil
	}
}

func NewFileSaver(fs storage.Storage) *FileSaver {
	return &FileSaver{
		fs: fs,
	}
}

func (this *FileSaver) GetFlag(job *dag.Job) (dag.JobStatus, error) {
	this.lock.Lock()
	defer this.lock.Unlock()

	output := job.Attrs["output"]
	for status, flag := range FlagSuffix {
		if exist, err := this.fs.IsExist(output + flag); err != nil {
			return dag.UnknownStatus, err
		} else if exist {
			return status, err
		}
	}
	return dag.NotStarted, nil
}

func (this *FileSaver) SetFlag(job *dag.Job, status dag.JobStatus) error {
	this.lock.Lock()
	defer this.lock.Unlock()

	if err := this.ClearFlag(job); err == nil {
		return err
	}
	output := job.Attrs["output"]
	if flag, ok := FlagSuffix[status]; ok {
		return this.fs.Touch(output + flag)
	} else {
		return fmt.Errorf("invalid status: %s", status)
	}
}

func (this *FileSaver) DeleteFlag(job *dag.Job, status dag.JobStatus) error {
	this.lock.Lock()
	defer this.lock.Unlock()

	output := job.Attrs["output"]
	if flag, ok := FlagSuffix[status]; ok {
		return this.fs.Rm(output + flag)
	} else {
		return fmt.Errorf("invalid status: %s", status)
	}
}

func (this *FileSaver) ClearFlag(job *dag.Job) error {
	this.lock.Lock()
	defer this.lock.Unlock()

	output := job.Attrs["output"]
	for _, flag := range FlagSuffix {
		if err := this.fs.Rm(output + flag); err != nil {
			return err
		}
	}
	return nil
}

//===================================================================
// Private
//===================================================================
