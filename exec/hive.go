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
 * @file hive.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sun Sep  6 22:23:42 2015
 *
 **/

package exec

import (
	"fmt"
	"github.com/colinmarc/hdfs"
	"github.com/crackcell/hpipe/config"
	"github.com/crackcell/hpipe/dag"
	"github.com/crackcell/hpipe/log"
	"os"
	"strings"
)

//===================================================================
// Public APIs
//===================================================================

type HiveExec struct {
	HDFSClient *hdfs.Client
	Jar        string
	Output     string
	Script     string
}

func NewHiveExec() *HiveExec {
	return &HiveExec{}
}

func (this *HiveExec) Setup() error {
	if client, err := hdfs.New(config.NameNode); err != nil {
		msg := fmt.Sprintf("connect to hdfs namenode failed: %s", config.NameNode)
		log.Fatal(msg)
		return fmt.Errorf(msg)
	} else {
		this.HDFSClient = client
	}
	return nil
}

func (this *HiveExec) Run(job *dag.Job) error {
	if !checkJobAttr(job, []string{"script", "output"}) {
		return fmt.Errorf("invalid job")
	}

	// !!!VERY IMPORTANT!!!
	// Many other operations relay on this TrimRight.
	job.Attrs["output"] = strings.TrimRight(job.Attrs["output"], "/")

	this.deleteRemainFiles(job)
	this.createOutput(job)
	this.createStatusFile(job, dag.Started)
	defer this.deleteStatusFile(job, dag.Started)

	args := this.genCmdArgs(job)
	log.Debugf("CMD: hive %s", strings.Join(args, " "))
	retcode, err := cmdExec(job.Name, "hive", args...)
	if err != nil {
		job.Status = dag.Failed
		this.createStatusFile(job, dag.Failed)
		return err
	}
	if retcode != 0 {
		job.Status = dag.Failed
		this.createStatusFile(job, dag.Failed)
		return fmt.Errorf("script failed: %d", retcode)
	}
	job.Status = dag.Finished
	this.createStatusFile(job, dag.Finished)
	return nil
}

func (this *HiveExec) GetStatus(job *dag.Job) (dag.JobStatus, error) {
	output := job.Attrs["output"]
	for status, flag := range hdfsJobStatusFlags {
		if exist, err := this.isFileExist(output + flag); err != nil {
			return dag.UnknownStatus, err
		} else if exist {
			return status, err
		}
	}
	return dag.NotStarted, nil
}

//===================================================================
// Private
//===================================================================

var hdfsJobStatusFlags = map[dag.JobStatus]string{
	dag.Started:  ".hpipe.started",
	dag.Finished: ".hpipe.finished",
	dag.Failed:   ".hpipe.failed",
}

func (this *HiveExec) isFileExist(path string) (bool, error) {
	_, err := this.HDFSClient.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			log.Debugf("path not exist: %s", path)
			return false, nil
		}
		return false, err
	}
	log.Debugf("path exist: %s", path)
	return true, nil
}

func (this *HiveExec) createOutput(job *dag.Job) error {
	tokens := strings.Split(job.Attrs["output"], "/")
	if len(tokens) <= 1 {
		return nil
	}
	return this.hdfsMkdirp(strings.Join(tokens[:len(tokens)-1], "/"))
}

func (this *HiveExec) createStatusFile(job *dag.Job, status dag.JobStatus) {
	output := job.Attrs["output"]
	flag, ok := hdfsJobStatusFlags[status]
	if ok {
		this.hdfsTouch(output + flag)
	}
}

func (this *HiveExec) deleteStatusFile(job *dag.Job, status dag.JobStatus) {
	output := job.Attrs["output"]
	flag, ok := hdfsJobStatusFlags[status]
	if ok {
		this.hdfsRm(output + flag)
	}
}

func (this *HiveExec) deleteRemainFiles(job *dag.Job) {
	output := job.Attrs["output"]
	for _, flag := range hdfsJobStatusFlags {
		this.hdfsRm(output + flag)
	}
	this.hdfsRm(output)
}

func (this *HiveExec) genCmdArgs(job *dag.Job) []string {
	args := []string{}
	args = append(args, "f")
	args = append(args, config.WorkPath+"/"+this.Script)

	for k, v := range job.Attrs {
		if _, ok := dag.JobReservedAttrs[k]; ok {
			continue
		}
		args = append(args, "-D")
		args = append(args, fmt.Sprintf("%s=%s", k, v))
	}

	return args
}

func (this *HiveExec) hdfsMkdirp(path string) error {
	log.Debugf("mkdir -p %s", path)
	return this.HDFSClient.MkdirAll(path, 0755)
}

func (this *HiveExec) hdfsRm(path string) error {
	log.Debugf("remove file: %s", path)
	return this.HDFSClient.Remove(path)
}

func (this *HiveExec) hdfsTouch(path string) error {
	log.Debugf("touch file: %s", path)
	return this.HDFSClient.CreateEmptyFile(path)
}
