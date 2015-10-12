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
	"github.com/crackcell/hpipe/log"
	"github.com/crackcell/hpipe/util"
	"regexp"
	"strings"
)

//===================================================================
// Public APIs
//===================================================================

type ShellExec struct {
}

func NewShellExec() *ShellExec {
	return &ShellExec{}
}

func (this *ShellExec) Setup() error {
	return nil
}

func (this *ShellExec) Run(job *dag.Job) error {
	if !checkJobAttr(job, []string{"output"}) {
		msg := "invalid job"
		log.Error(msg)
		return fmt.Errorf(msg)
	}

	// !!!VERY IMPORTANT!!!
	// Many other operations relay on this TrimRight.
	job.Attrs["output"] = strings.TrimRight(job.Attrs["output"], "/")

	args := this.genCmdArgs(job)
	log.Debugf("CMD: bash %s", strings.Join(args, " "))
	retcode, err := util.ExecCmd(job.Name, "bash", args...)
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

var paramPattern = regexp.MustCompile("'[\\w\\s\\._-]*'|[\\w\\._-]+")

func (this *ShellExec) genCmdArgs(job *dag.Job) []string {
	args := []string{}

	if v, ok := job.Attrs["cmd"]; ok {
		args = append(args, "-c")
		args = append(args, v)
	} else if v, ok := job.Attrs["script"]; ok {
		params := paramPattern.FindAllStringSubmatch(v, -1)
		if len(params) > 0 {
			args = append(args, config.WorkPath+"/"+params[0][0])
			for i := 1; i < len(params); i++ {
				args = append(args, strings.Trim(params[i][0], "'"))
			}
		}
	}

	return args
}
