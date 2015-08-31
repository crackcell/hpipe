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

type HadoopExec struct {
	HDFSClient *hdfs.Client
	Jar        string
	Input      string
	Output     string
	Mapper     string
	Reducer    string
}

func NewHadoopExec() *HadoopExec {
	return &HadoopExec{}
}

func (this *HadoopExec) Setup() error {
	if len(config.HadoopStreamingJar) == 0 {
		msg := fmt.Sprintf("invalid hadoop streaming jar: %s", config.HadoopStreamingJar)
		log.Errorf(msg)
		return fmt.Errorf(msg)
	} else {
		this.Jar = config.HadoopStreamingJar
	}
	if client, err := hdfs.New(config.NameNode); err != nil {
		msg := fmt.Sprintf("connect to hdfs namenode failed: %s", config.NameNode)
		log.Fatal(msg)
		return fmt.Errorf(msg)
	} else {
		this.HDFSClient = client
	}
	return nil
}

func (this *HadoopExec) Run(job *dag.Job) error {
	if !checkJobAttr(job, []string{"mapper", "input", "output"}) {
		return fmt.Errorf("invalid job")
	}

	// !!!VERY IMPORTANT!!!
	// Many other operations relay on this TrimRight.
	job.Attrs["output"] = strings.TrimRight(job.Attrs["output"], "/")

	this.deleteRemainFiles(job)
	this.createOutput(job)
	this.createStatusFile(job, dag.Started)
	defer this.deleteStatusFile(job, dag.Started)

	retcode, err := cmdExec(job.Name, "hadoop", this.genCmdArgs(job)...)
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

func (this *HadoopExec) GetStatus(job *dag.Job) (dag.JobStatus, error) {
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

func (this *HadoopExec) isFileExist(path string) (bool, error) {
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

func (this *HadoopExec) createOutput(job *dag.Job) error {
	tokens := strings.Split(job.Attrs["output"], "/")
	if len(tokens) <= 1 {
		return nil
	}
	return this.hdfsMkdirp(strings.Join(tokens[:len(tokens)-1], "/"))
}

func (this *HadoopExec) createStatusFile(job *dag.Job, status dag.JobStatus) {
	output := job.Attrs["output"]
	flag, ok := hdfsJobStatusFlags[status]
	if ok {
		this.hdfsTouch(output + flag)
	}
}

func (this *HadoopExec) deleteStatusFile(job *dag.Job, status dag.JobStatus) {
	output := job.Attrs["output"]
	flag, ok := hdfsJobStatusFlags[status]
	if ok {
		this.hdfsRm(output + flag)
	}
}

func (this *HadoopExec) deleteRemainFiles(job *dag.Job) {
	output := job.Attrs["output"]
	for _, flag := range hdfsJobStatusFlags {
		this.hdfsRm(output + flag)
	}
	this.hdfsRm(output)
}

func (this *HadoopExec) genCmdArgs(job *dag.Job) []string {
	args := []string{}
	args = append(args, "jar")
	args = append(args, this.Jar)

	args = append(args, "-input")
	args = append(args, "\""+job.Attrs["input"]+"\"")

	args = append(args, "-output")
	args = append(args, "\""+job.Attrs["output"]+"\"")

	args = append(args, "-mapper")
	args = append(args, "\""+job.Attrs["mapper"]+"\"")

	if val, ok := job.Attrs["reducer"]; ok {
		args = append(args, "-reducer")
		args = append(args, "\""+val+"\"")
	}
	return args
}

func (this *HadoopExec) hdfsMkdirp(path string) error {
	log.Debugf("mkdir -p %s", path)
	return this.HDFSClient.MkdirAll(path, 0755)
}

func (this *HadoopExec) hdfsRm(path string) error {
	log.Debugf("remove file: %s", path)
	return this.HDFSClient.Remove(path)
}

func (this *HadoopExec) hdfsTouch(path string) error {
	log.Debugf("touch file: %s", path)
	return this.HDFSClient.CreateEmptyFile(path)
}
