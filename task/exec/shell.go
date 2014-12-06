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
 * @file shell.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sat Dec  6 21:12:40 2014
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

func NewShellExec() Exec {
	return new(shellExec)
}

//===================================================================
// Private
//===================================================================

type shellExec struct {
	job *ast.Job
}

func (this *shellExec) Run(job *ast.Job) (string, error) {
	this.job = job
	if !ValidProp(this.job.Prop, shellPropNames) {
		return hpipe.FAIL, fmt.Errorf("not valid job")
	}
	return hpipe.FAIL, nil
}

var shellPropNames []string = []string{
	"cmd",
}
