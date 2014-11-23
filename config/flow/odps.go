/***************************************************************
 *
 * Copyright (c) 2014, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 * ODPS Jobs
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

const (
	SQLJob = iota
	MapredJob
)

type ODPSJob struct {
	Name      string
	Var       map[string]string
	File      string
	AccessID  string
	AccessKey string
	Project   string
	Endpoint  string
}

func NewODPSJob() *ODPSJob {
	return &ODPSJob{Var: make(map[string]string)}
}

func (this *ODPSJob) SetName(n string) {
	this.Name = n
}

func (this *ODPSJob) GetName() string {
	return this.Name
}

func (this *ODPSJob) SetVar(m map[string]string) {
	this.Var = m
}

func (this *ODPSJob) GetVar() map[string]string {
	return this.Var
}

func (this *ODPSJob) SetFile(f string) {
	this.File = f
}

func (this *ODPSJob) GetFile() string {
	return this.File
}

func (this *ODPSJob) IsValid() bool {
	return util.IsInMap(
		[]string{"accessid", "accesskey", "project", "endpoint", "jobtype"},
		this.Var)
}

func (this *ODPSJob) DoJob() {}

func (this *ODPSJob) CheckStatus() int {
	return Done
}

func (this *ODPSJob) DebugString() string {
	return fmt.Sprintf("odps_job:{name:%s, file:%s, var:%v}",
		this.Name, this.File, this.Var)
}

//===================================================================
// Private
//===================================================================
