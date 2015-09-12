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
// JobType
//===================================================================

type JobType int

const (
	DummyJob JobType = iota
	HadoopJob
	HiveJob
	OdpsJob
	ShellJob
	UnknownJob
)

func (typ JobType) String() string {
	switch typ {
	case DummyJob:
		return "dummy"
	case HadoopJob:
		return "hadoop"
	case HiveJob:
		return "hive"
	case OdpsJob:
		return "odps"
	case ShellJob:
		return "shell"
	default:
		return "unknown"
	}
}

func ParseJobType(typ string) JobType {
	switch typ {
	case "dummy":
		return DummyJob
	case "hadoop":
		return HadoopJob
	case "hive":
		return HiveJob
	case "odps":
		return OdpsJob
	case "shell":
		return ShellJob
	default:
		return UnknownJob
	}
}

//===================================================================
// JobStatus
//===================================================================

type JobStatus int

const (
	NotStarted = iota
	Started
	Finished
	Failed
	UnknownStatus
)

func (status JobStatus) String() string {
	switch status {
	case NotStarted:
		return "not_started"
	case Started:
		return "started"
	case Finished:
		return "finished"
	case Failed:
		return "failed"
	default:
		return "unknown"
	}
}

func ParseJobStatus(status string) JobStatus {
	switch status {
	case "not_started":
		return NotStarted
	case "started":
		return Started
	case "finished":
		return Finished
	case "failed":
		return Failed
	default:
		return UnknownStatus
	}
}

//===================================================================
// Job
//===================================================================

var JobReservedAttrs = map[string]int{
	"name":    1,
	"type":    1,
	"vars":    1,
	"input":   1,
	"output":  1,
	"mapper":  1,
	"reducer": 1,
	"script":  1,
	"hql":     1,
}

type Job struct {
	Name   string
	Type   JobType
	Status JobStatus
	Attrs  Attrs
	Prev   []string
	Post   []string
	Vars   map[string]string
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
	this.Status = job.Status
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
