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
	cmd string
}

func (this *ODPSExec) Run(job *ast.Job) (string, error) {
	if !validateJob(job) {
		return ast.FAIL, fmt.Errorf("not valid job")
	}
	this.job = job
	//log.Debug(job.DebugString())
	cmdstr := "odpscmd -u " + this.getProp("access_id") +
		" -p " + this.getProp("access_key") +
		" --project=" + this.getProp("project") +
		" --endpoint=" + this.getProp("endpoint") +
		" -e " + this.getProp("cmd")
	log.Debugf("cmd: %s", cmdstr)
	exitcode, err := CmdExec(job.InstanceID, "odpscmd",
		"-u ", this.getProp("access_id"),
		"-p ", this.getProp("access_key"),
		"--project="+this.getProp("project"),
		"--endpoint="+this.getProp("endpoint"),
		"-e", this.getProp("cmd"))
	if err != nil {
		return ast.FAIL, err
	}
	if exitcode != 0 {
		return ast.FAIL, nil
	}
	return ast.DONE, nil
}

//===================================================================
// Private
//===================================================================

var propNames []string = []string{
	"access_id",
	"access_key",
	"project",
	"endpoint",
	"cmd",
}

func validateJob(job *ast.Job) bool {
	for _, p := range propNames {
		if v, ok := job.Prop[p]; !ok || len(v) == 0 {
			log.Fatalf("%s not found or empty", p)
			return false
		}
	}
	return true
}

func (this *ODPSExec) getProp(key string) string {
	v, ok := this.job.Prop[key]
	if !ok {
		panic(fmt.Errorf("no prop for %s", key))
	}
	return v
}
