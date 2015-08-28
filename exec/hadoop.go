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
	"fmt"
	"github.com/crackcell/hpipe/config"
	"github.com/crackcell/hpipe/dag"
)

//===================================================================
// Public APIs
//===================================================================

type HadoopExec struct {
	HadoopPrefix string
	Input        string
	Output       string
	Mapper       string
	Reducer      string
}

func NewHadoopExec() *HadoopExec {
	return &HadoopExec{}
}

func (this *HadoopExec) Setup() error {
	return nil
}

func (this *HadoopExec) Run(job *dag.Job) error {
	retcode, err := cmdExec(job.Name, "bash",
		config.WorkPath+"/"+job.Attrs["script"])
	if err != nil {
		return err
	}
	return fmt.Errorf("script failed: %d", retcode)
}

func (this *HadoopExec) GetJobStatus(job *dag.Job) dag.JobStatus {
	// TODO
	return dag.Finished
}

func (this *HadoopExec) CheckJobAttrs(job *dag.Job) bool {
	return checkJobAttr(job, []string{"script", "input", "output"})
}

//===================================================================
// Private
//===================================================================
