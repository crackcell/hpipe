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

type Flow struct {
	Entry *Step
}

type Step struct {
	Name     string
	Dep      []*Step
	Do       []Job
	Var      map[string]string
	Resource string
}

func NewFlow() *Flow {
	return &Flow{}
}

func (this *Flow) DebugString() string {
	return this.Entry.DebugString()
}

func NewStep() *Step {
	return &Step{Var: make(map[string]string)}
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

const (
	Todo = iota
	Doing
	Done
	Fail
)

type Job interface {
	SetName(n string)
	GetName() string
	SetVar(m map[string]string)
	GetVar() map[string]string
	SetFile(f string)
	GetFile() string
	AddProp(key, value string)
	SetProp(p map[string]string)
	GetProp() map[string]string
	DebugString() string
}

//===================================================================
// ODPS Job

type ODPSJob struct {
	Name       string
	Var        map[string]string
	File       string
	AccessId   string
	AccessKey  string
	Project    string
	Endpoint   string
	Properties map[string]string
}

func NewODPSJob() *ODPSJob {
	return &ODPSJob{
		Var:        make(map[string]string),
		Properties: make(map[string]string)}
}

func (this *ODPSJob) SetName(n string)            { this.Name = n }
func (this *ODPSJob) GetName() string             { return this.Name }
func (this *ODPSJob) SetVar(m map[string]string)  { this.Var = m }
func (this *ODPSJob) GetVar() map[string]string   { return this.Var }
func (this *ODPSJob) SetFile(f string)            { this.File = f }
func (this *ODPSJob) GetFile() string             { return this.File }
func (this *ODPSJob) AddProp(k, v string)         { this.Properties[k] = v }
func (this *ODPSJob) SetProp(p map[string]string) { this.Properties = p }
func (this *ODPSJob) GetProp() map[string]string  { return this.Properties }

func (this *ODPSJob) DebugString() string {
	return fmt.Sprintf("odps_job:{name:%s, file:%s, var:%v, prop:%v}",
		this.Name, this.File, this.Var, this.Properties)
}

//===================================================================
// Hadoop Job

type HadoopJob struct {
	Name       string
	Var        map[string]string
	File       string
	Properties map[string]string
}

func NewHadoopJob() *HadoopJob {
	return &HadoopJob{
		Var:        make(map[string]string),
		Properties: make(map[string]string)}
}
func (this *HadoopJob) SetName(n string)            { this.Name = n }
func (this *HadoopJob) GetName() string             { return this.Name }
func (this *HadoopJob) SetVar(m map[string]string)  { this.Var = m }
func (this *HadoopJob) GetVar() map[string]string   { return this.Var }
func (this *HadoopJob) SetFile(f string)            { this.File = f }
func (this *HadoopJob) GetFile() string             { return this.File }
func (this *HadoopJob) AddProp(k, v string)         { this.Properties[k] = v }
func (this *HadoopJob) SetProp(p map[string]string) { this.Properties = p }
func (this *HadoopJob) GetProp() map[string]string  { return this.Properties }

func (this *HadoopJob) DebugString() string {
	return fmt.Sprintf("hadoop_job:{name:%s, fileL%s, var:%v, prop:%v}",
		this.File, this.Name, this.Var, this.Properties)
}

//===================================================================
// Private
//===================================================================
