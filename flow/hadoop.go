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
	"fmt"
)

//===================================================================
// Public APIs
//===================================================================

type HadoopJob struct {
	Name string
	Var  map[string]string
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

func (this *HadoopJob) IsValid() bool {
	return IsInMap([]string{"mapred.job.name"}, this.Var)
}

func (this *HadoopJob) DoJob() {}

func (this *HadoopJob) CheckStatus() int {
	return Done
}

func (this *HadoopJob) DebugString() string {
	return fmt.Sprintf("hadoop_job:{name:%s,var:%v}", this.Name, this.Var)
}

//===================================================================
// Private
//===================================================================
