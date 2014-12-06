/***************************************************************
 *
 * Copyright (c) 2014, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 *
 *
 * @file odps.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Fri Nov 28 16:57:33 2014
 *
 **/

package exec

import (
	"../../../hpipe"
	"../../yafl/ast"
	"fmt"
)

//===================================================================
// Public APIs
//===================================================================

func NewODPSExec() Exec {
	return new(ODPSExec)
}

type ODPSExec struct {
	job *ast.Job
}

func (this *ODPSExec) Run(job *ast.Job) (string, error) {
	this.job = job
	if !ValidProp(this.job.Prop, odpsPropNames) {
		return hpipe.FAIL, fmt.Errorf("not valid job")
	}

	args := PrepareArgList(this.job.Prop, odpsArgs)
	LogArgList("odpscmd", args...)

	exitcode, err := CmdExec(job.InstanceID, "odpscmd", args...)
	if err != nil || exitcode != 0 {
		return hpipe.FAIL, err
	}

	return hpipe.DONE, nil
}

//===================================================================
// Private
//===================================================================

var odpsPropNames []string = []string{
	"access_id",
	"access_key",
	"project",
	"endpoint",
	"cmd",
}
var odpsArgs [][]string = [][]string{
	[]string{"access_id", "-u", "s", "n"},
	[]string{"access_key", "-p", "s", "n"},
	[]string{"project", "--project=", "c", "n"},
	[]string{"endpoint", "--endpoint=", "c", "n"},
	[]string{"cmd", "-e", "s", "n"},
}
