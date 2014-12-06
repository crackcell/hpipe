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
	"strings"
)

//===================================================================
// Public APIs
//===================================================================

func NewHadoopExec() Exec {
	return new(HadoopExec)
}

type HadoopExec struct {
	job  *ast.Job
	jar  []string
	file []string
}

func (this *HadoopExec) Run(job *ast.Job) (string, error) {
	this.job = job
	if !this.valid() {
		return ast.FAIL, fmt.Errorf("not valid job")
	}
	if err := this.setup(); err != nil {
		return ast.FAIL, err
	}
	cmdstr := "hadoop" + "jar" +
		"/home/a/lib/hadoop/default/hadoop-examples-*.jar" +
		"grep input output 'dfs[a-z.]+'"
	log.Debugf("cmd: %s", cmdstr)
	exitcode, err := CmdExec(job.InstanceID, "hadoop")
	if err != nil || exitcode != 0 {
		return ast.FAIL, err
	}
	return ast.DONE, nil
}

//===================================================================
// Private
//===================================================================

func (this *HadoopExec) setup() error {
	for _, file := range strings.Split(GetProp(this.job.Prop, "file"), ",") {
		file = strings.Trim(file, " ")
		this.file = append(this.file, file)
	}
	for _, jar := range strings.Split(GetProp(this.job.Prop, "jar"), ",") {
		jar = strings.Trim(jar, " ")
		this.file = append(this.jar, jar)
	}
	return nil
}

var hadoopPropNames []string = []string{
	"hadoop_home",
}

func (this *HadoopExec) valid() bool {
	for _, p := range hadoopPropNames {
		if v, ok := this.job.Prop[p]; !ok || len(v) == 0 {
			log.Fatalf("%s not found or empty", p)
			return false
		}
	}
	return true
}
