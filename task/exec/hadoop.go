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
	"../../log"
	"../../yafl/ast"
	"fmt"
)

//===================================================================
// Public APIs
//===================================================================

func NewHadoopExec() Exec {
	return new(HadoopExec)
}

type HadoopExec struct {
	job    *ast.Job
	jar    []string
	file   []string
	input  []string
	output []string
}

func (this *HadoopExec) Run(job *ast.Job) (string, error) {
	this.job = job
	if !this.valid() {
		return ast.FAIL, fmt.Errorf("not valid job")
	}

	args := this.prepareArgList()

	LogArgList("hadoop", args...)

	exitcode, err := CmdExec(job.InstanceID, "hadoop", args...)
	if err != nil || exitcode != 0 {
		return ast.FAIL, err
	}
	return ast.DONE, nil
}

//===================================================================
// Private
//===================================================================

var hadoopPropNames []string = []string{
	"hadoop_home",
}

func (this *HadoopExec) valid() bool {
	for _, p := range hadoopPropNames {
		if !ExistProp(this.job.Prop, p) {
			log.Fatalf("%s not found or empty", p)
			return false
		}
	}
	return true
}

func (this *HadoopExec) prepareArgList() []string {
	var args []string

	args = append(args, PrepareArg(this.job.Prop, "jar", "jar")...)
	args = append(args, PrepareArg(this.job.Prop, "file", "-file")...)
	args = append(args, PrepareArg(this.job.Prop, "mapper", "-mapper")...)
	args = append(args, PrepareArg(this.job.Prop, "reducer", "-reducer")...)
	args = append(args, PrepareArg(this.job.Prop, "input", "-input")...)
	args = append(args, PrepareArg(this.job.Prop, "output", "-output")...)

	return args
}
