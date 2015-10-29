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
	"github.com/crackcell/hpipe/log"
	"github.com/crackcell/hpipe/storage"
	"github.com/crackcell/hpipe/util"
	"strings"
)

//===================================================================
// Public APIs
//===================================================================

type HadoopExec struct {
	hdfs *storage.HDFS
	jar  string
}

func NewHadoopExec() *HadoopExec {
	return &HadoopExec{}
}

func (this *HadoopExec) Setup() error {
	if len(config.HadoopStreamingJar) == 0 {
		msg := fmt.Sprintf("hadoop streaming jar not set")
		log.Errorf(msg)
		return fmt.Errorf(msg)
	} else {
		this.jar = config.HadoopStreamingJar
	}

	if fs, err := storage.NewHDFS(config.NameNode); err != nil {
		msg := fmt.Sprintf("connect to hdfs namenode failed: %s", config.NameNode)
		log.Fatal(msg)
		return fmt.Errorf(msg)
	} else {
		this.hdfs = fs.(*storage.HDFS)
	}

	return nil
}

func (this *HadoopExec) Run(job *dag.Job) error {
	if !job.ValidateAttr([]string{"mapper", "input", "output"}) {
		msg := "invalid job"
		log.Error(msg)
		return fmt.Errorf(msg)
	}

	// !!!VERY IMPORTANT!!!
	// Many other operations relay on this TrimRight.
	job.Attrs["output"] = strings.TrimRight(job.Attrs["output"], "/")

	this.hdfs.Rm(job.Attrs["output"])
	this.createOutput(job)

	args := this.genCmdArgs(job)
	log.Debugf("CMD: hadoop %v", args)
	retcode, err := util.ExecCmd(job.Name, "hadoop", args...)
	if err != nil {
		job.Status = dag.Failed
		return err
	}
	if retcode != 0 {
		job.Status = dag.Failed
		return fmt.Errorf("script failed: %d", retcode)
	}
	job.Status = dag.Finished
	return nil
}

//===================================================================
// Private
//===================================================================

func (this *HadoopExec) createOutput(job *dag.Job) error {
	tokens := strings.Split(job.Attrs["output"], "/")
	if len(tokens) <= 1 {
		return nil
	}
	return this.hdfs.MkdirP(strings.Join(tokens[:len(tokens)-1], "/"))
}

func (this *HadoopExec) genCmdArgs(job *dag.Job) []string {
	args := []string{}
	args = append(args, "jar")
	args = append(args, this.jar)

	args = append(args, "-D")
	args = append(args, fmt.Sprintf("mapred.job.name=%s", job.Name))

	for k, v := range job.Attrs {
		if _, ok := dag.JobReservedAttrs[k]; ok {
			continue
		}
		args = append(args, "-D")
		args = append(args, fmt.Sprintf("%s=%s", k, v))
	}

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
