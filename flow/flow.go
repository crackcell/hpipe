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
 * @file flow.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Thu Nov 13 12:00:17 2014
 *
 **/

package flow

import (
	"fmt"
	"strings"
)

//===================================================================
// Public APIs
//===================================================================

type Step struct {
	Name string
	Dep  []*Step
	Do   []Job
	Var  map[string]string
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
	DoJob()
	IsValid() bool
	CheckStatus() int
	DebugString() string
}

func IsInMap(keys []string, m map[string]string) bool {
	for _, k := range keys {
		if _, ok := m[k]; !ok {
			return false
		}
	}
	return true
}

//===================================================================
// Private
//===================================================================
