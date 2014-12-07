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
 * @file hadoop.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sat Dec  6 15:19:43 2014
 *
 **/

package exec

import (
	"../../../hpipe"
	_ "../../log"
	"../../yafl/ast"
	"fmt"
)

//===================================================================
// Public APIs
//===================================================================

func NewHadoopExec() Exec {
	return new(hadoopExec)
}

//===================================================================
// Private
//===================================================================

type hadoopExec struct {
	job *ast.Job
}

func (this *hadoopExec) Run(job *ast.Job) (string, error) {
	this.job = job
	if !ValidProp(this.job.Prop, hadoopPropNames) {
		return hpipe.FAIL, fmt.Errorf("not valid job")
	}

	args := PrepareArgList(this.job.Prop, hadoopArgs)
	LogArgList("hadoop", args...)

	exitcode, err := CmdExec(job.InstanceID, "hadoop", args...)
	if err != nil || exitcode != 0 {
		return hpipe.FAIL, err
	}

	return hpipe.DONE, nil
}

var hadoopPropNames []string = []string{
	"hadoop_home",
}

var hadoopArgs [][]string = [][]string{
	[]string{"jar", "jar", "s", "s"},
	[]string{"file", "-file", "s", "s"},
	[]string{"mapper", "-mapper", "s", "s"},
	[]string{"reducer", "-reducer", "s", "s"},
	[]string{"input", "-input", "s", "s"},
	[]string{"output", "-output", "s", "s"},
}
