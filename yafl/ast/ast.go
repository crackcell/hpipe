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
 * @file ast.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Thu Nov 13 12:00:17 2014
 *
 **/

package ast

import (
	"fmt"
	"strings"
)

//===================================================================
// Public APIs
//===================================================================

const (
	TODO = "todo"
	RUN  = "run"
	DONE = "done"
	FAIL = "fail"
)

type Flow struct {
	Entry *Step
}

type Step struct {
	Name     string
	Dep      []*Step
	Do       []*Job
	Var      map[string]string
	Resource string
	Status   string
}

type Job struct {
	Name       string
	InstanceID string
	Type       string
	Var        map[string]string
	File       string
	Prop       map[string]string
	Status     string
}

func NewFlow() *Flow {
	return &Flow{}
}

func (this *Flow) DebugString() string {
	return this.Entry.DebugString()
}

func NewStep() *Step {
	return &Step{
		Var:    make(map[string]string),
		Status: TODO,
	}
}

func (this *Step) DebugString() string {
	return this.debugString(0)
}

func (this *Step) debugString(depth int) string {
	indent := strings.Repeat("\t", depth)
	str := fmt.Sprintf("%s%s:{\n", indent, this.Name)

	str += fmt.Sprintf("%s\tvar:{", indent)
	if len(this.Var) == 0 {
		str += "}\n"
	} else {
		str += "\n"
		for k, v := range this.Var {
			str += fmt.Sprintf("%s\t\t%s=%s\n", indent, k, v)
		}
		str += indent + "\t}\n"
	}

	str += indent + "\tdep:{"
	if len(this.Dep) == 0 {
		str += "}\n"
	} else {
		str += "\n"
		for _, dep := range this.Dep {
			str += dep.debugString(depth+2) + "\n"
		}
		str += indent + "\t}\n"
	}

	str += indent + "\tdo:{"
	if len(this.Do) == 0 {
		str += "}\n"
	} else {
		str += "\n"
		for _, do := range this.Do {
			str += fmt.Sprintf("%s\t\t%s\n", indent, do.DebugString())
		}
		str += indent + "\t}\n"
	}

	str += indent + "}"
	return str
}

func NewJob() *Job {
	return &Job{
		Var:    make(map[string]string),
		Prop:   make(map[string]string),
		Status: TODO,
	}
}

func (this *Job) DebugString() string {
	return fmt.Sprintf("job:{name:%s,type:%s,id:%s,status:%s,file:%s,var:%v,prop:%v}",
		this.Name, this.Type, this.InstanceID, this.Status, this.File, this.Var,
		this.Prop)
}

//===================================================================
// Private
//===================================================================
