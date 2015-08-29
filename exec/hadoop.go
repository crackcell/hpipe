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
		log.Errorf(msg)
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
	retcode, err := cmdExec(job.Name, "hadoop", this.genCmdArgs(job)...)
	if err != nil {
		return err
	}
	if retcode != 0 {
		return fmt.Errorf("script failed: %d", retcode)
	}
	return nil
}

func (this *HadoopExec) GetStatus(job *dag.Job) (dag.JobStatus, error) {
	output := strings.TrimRight(job.Attrs["output"], "/")
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
