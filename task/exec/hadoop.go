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
	"../../config"
	"../../yafl/ast"
	"fmt"
	"strings"
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

	files, _ := this.job.Prop["file"]
	var str string
	for _, file := range strings.Split(files, ",") {
		file = strings.Trim(file, " ")
		if len(file) == 0 {
			continue
		}
		if string(file[0]) != "/" {
			str += config.WorkPath + "/" + file + ","
		} else {
			str += file + ","
		}
	}
	this.job.Prop["file"] = str

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
	[]string{"jar", "jar", "s"},
	[]string{"file", "-file", "s"},
	[]string{"mapper", "-mapper", "s"},
	[]string{"reducer", "-reducer", "s"},
	[]string{"input", "-input", "s"},
	[]string{"output", "-output", "s"},
}