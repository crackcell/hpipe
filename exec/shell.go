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
 * @date Tue Aug 25 18:28:05 2015
 *
 **/

package exec

import (
	"fmt"
	"github.com/crackcell/hpipe/config"
	"github.com/crackcell/hpipe/dag"
	"github.com/crackcell/hpipe/exec/filesystem"
	"github.com/crackcell/hpipe/log"
	"strings"
)

//===================================================================
// Public APIs
//===================================================================

type ShellExec struct {
	status *StatusKeeper
	hdfs   *filesystem.HDFS
}

func NewShellExec() *ShellExec {
	return &ShellExec{}
}

func (this *ShellExec) Setup() error {
	if len(config.HadoopStreamingJar) == 0 {
		msg := fmt.Sprintf("invalid hadoop streaming jar: %s", config.HadoopStreamingJar)
		log.Errorf(msg)
		return fmt.Errorf(msg)
	}
	fs, err := filesystem.NewHDFS(config.NameNode)
	if err != nil {
		msg := fmt.Sprintf("connect to hdfs namenode failed: %s", config.NameNode)
		log.Fatal(msg)
		return fmt.Errorf(msg)
	}
	this.status = NewStatusKeeper(fs)
	this.hdfs = fs.(*filesystem.HDFS)
	return nil
}

func (this *ShellExec) Run(job *dag.Job) error {
	if !checkJobAttr(job, []string{"script", "output"}) {
		return fmt.Errorf("invalid job")
	}

	// !!!VERY IMPORTANT!!!
	// Many other operations relay on this TrimRight.
	job.Attrs["output"] = strings.TrimRight(job.Attrs["output"], "/")

	this.status.ClearStatus(job)
	this.hdfs.Rm(job.Attrs["output"])
	this.createOutput(job)
	this.status.SetStatus(job, dag.Started)
	defer this.status.DeleteStatus(job, dag.Started)

	args := this.genCmdArgs(job)
	log.Debugf("CMD: bash %s", strings.Join(args, " "))
	retcode, err := cmdExec(job.Name, "bash", args...)
	if err != nil {
		job.Status = dag.Failed
		this.status.SetStatus(job, dag.Failed)
		return err
	}
	if retcode != 0 {
		job.Status = dag.Failed
		this.status.SetStatus(job, dag.Failed)
		return fmt.Errorf("script failed: %d", retcode)
	}
	job.Status = dag.Finished
	this.status.SetStatus(job, dag.Finished)
	return nil
}

//===================================================================
// Private
//===================================================================

func (this *ShellExec) createOutput(job *dag.Job) error {
	tokens := strings.Split(job.Attrs["output"], "/")
	if len(tokens) <= 1 {
		return nil
	}
	return this.hdfs.MkdirP(strings.Join(tokens[:len(tokens)-1], "/"))
}

func (this *ShellExec) genCmdArgs(job *dag.Job) []string {
	args := []string{}

	args = append(args, job.Attrs["script"])

	for k, v := range job.Attrs {
		if _, ok := dag.JobReservedAttrs[k]; ok {
			continue
		}
		args = append(args, "-D")
		args = append(args, fmt.Sprintf("%s=%s", k, v))
	}

	return args
}
