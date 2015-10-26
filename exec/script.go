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
	"strings"
)

//===================================================================
// Public APIs
//===================================================================

type ScriptExec struct {
}

func NewScriptExec() *ScriptExec {
	return &ScriptExec{}
}

func (this *ScriptExec) Setup() error {
	return nil
}

func (this *ScriptExec) Run(job *dag.Job) error {
	if !checkJobAttr(job, []string{"output", "script", "interpreter"}) {
		msg := "invalid job: missing output, script or interpreter"
		log.Error(msg)
		return fmt.Errorf(msg)
	}

	// !!!VERY IMPORTANT!!!
	// Many other operations relay on this TrimRight.
	job.Attrs["output"] = strings.TrimRight(job.Attrs["output"], "/")

	args := this.genCmdArgs(job)
	intepreter := getProp("interpreter", job.Attrs)
	log.Debugf("CMD: %s %v", intepreter, args)
	retcode, err := util.ExecCmd(job.Name, intepreter, args...)
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

//var paramPattern = regexp.MustCompile("'[\\w\\s\\._-\\/]*'|[\\w\\._-\\/]+")

func (this *ScriptExec) genCmdArgs(job *dag.Job) []string {
	args := []string{}
	v := getProp("script", job.Attrs)
	params := strings.Split(v, " ")
	if len(params) > 0 {
		args = append(args, config.WorkPath+"/"+params[0])
		for i := 1; i < len(params); i++ {
			args = append(args, params[i])
		}
	}
	return args
}
