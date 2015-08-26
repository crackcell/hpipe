/***************************************************************
 *
 * Copyright (c) 2015, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 *
 *
 * @file node.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Tue Aug 18 04:25:23 2015
 *
 **/

package dag

import (
	"fmt"
)

//===================================================================
// Public APIs
//===================================================================

type JobType int

func ParseJobType(typ string) JobType {
	switch typ {
	case "dummy":
		return DummyJob
	case "hadoop_streaming":
		return HadoopStreamingJob
	case "shell":
		return ShellJob
	default:
		return UnknownJob
	}
}

const (
	DummyJob JobType = iota
	HadoopStreamingJob
	ShellJob
	JobCount
	UnknownJob
)

type Job struct {
	Name     string
	Type     JobType
	Attrs    Attrs
	Prev     []string
	Post     []string
	Vars     map[string]string
	Finished bool
}

func NewJob() *Job {
	return &Job{
		Attrs: NewAttrs(),
		Prev:  []string{},
		Post:  []string{},
		Vars:  make(map[string]string),
	}
}

func (this *Job) Assign(job *Job) {
	this.Name = job.Name
	this.Type = job.Type
	this.Attrs = job.Attrs
	this.Prev = job.Prev
	this.Post = job.Post
}

func (this *Job) String() string {
	return fmt.Sprintf("Job{name=%s,attrs=%v,prev=%v,post=%v}",
		this.Name, this.Attrs, this.Prev, this.Post)
}

//===================================================================
// Private
//===================================================================
