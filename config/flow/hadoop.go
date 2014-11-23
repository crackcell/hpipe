/***************************************************************
 *
 * Copyright (c) 2014, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 * Hadoop Jobs
 *
 * @file odps.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Fri Nov 14 00:43:03 2014
 *
 **/

package flow

import (
	"../../util"
	"fmt"
)

//===================================================================
// Public APIs
//===================================================================

type HadoopJob struct {
	Name string
	Var  map[string]string
	File string
}

func NewHadoopJob() *HadoopJob {
	return &HadoopJob{Var: make(map[string]string)}
}

func (this *HadoopJob) SetName(n string) {
	this.Name = n
}

func (this *HadoopJob) GetName() string {
	return this.Name
}

func (this *HadoopJob) SetVar(m map[string]string) {
	this.Var = m
}

func (this *HadoopJob) GetVar() map[string]string {
	return this.Var
}

func (this *HadoopJob) SetFile(f string) {
	this.File = f
}

func (this *HadoopJob) GetFile() string {
	return this.File
}

func (this *HadoopJob) IsValid() bool {
	return util.IsInMap([]string{"mapred.job.name"}, this.Var)
}

func (this *HadoopJob) DoJob() {}

func (this *HadoopJob) CheckStatus() int {
	return Done
}

func (this *HadoopJob) DebugString() string {
	return fmt.Sprintf("hadoop_job:{name:%s, fileL%s, var:%v}",
		this.File, this.Name, this.Var)
}

//===================================================================
// Private
//===================================================================
