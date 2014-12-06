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
	"../../log"
	"../../yafl/ast"
	"fmt"
	_ "time"
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
	if !this.valid() {
		return ast.FAIL, fmt.Errorf("not valid job")
	}
	cmdstr := "odpscmd -u " + GetProp(this.job.Prop, "access_id") +
		" -p " + GetProp(this.job.Prop, "access_key") +
		" --project=" + GetProp(this.job.Prop, "project") +
		" --endpoint=" + GetProp(this.job.Prop, "endpoint") +
		" -e " + GetProp(this.job.Prop, "cmd")
	log.Debugf("cmd: %s", cmdstr)
	exitcode, err := CmdExec(job.InstanceID, "odpscmd",
		"-u ", GetProp(this.job.Prop, "access_id"),
		"-p ", GetProp(this.job.Prop, "access_key"),
		"--project="+GetProp(this.job.Prop, "project"),
		"--endpoint="+GetProp(this.job.Prop, "endpoint"),
		"-e", GetProp(this.job.Prop, "cmd"))
	if err != nil || exitcode != 0 {
		return ast.FAIL, err
	}
	return ast.DONE, nil
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

func (this *ODPSExec) valid() bool {
	for _, p := range odpsPropNames {
		if !ExistProp(this.job.Prop, p) {
			log.Fatalf("%s not found or empty", p)
			return false
		}
	}
	return true
}
