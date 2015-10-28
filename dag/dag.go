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
 * @file dag.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Mon Aug 17 21:27:58 2015
 *
 **/

package dag

import (
	"fmt"
	"github.com/crackcell/hpipe/dag/symbol"
	"github.com/crackcell/hpipe/log"
	"io/ioutil"
	"regexp"
	"strings"
)

//===================================================================
// Public APIs
//===================================================================

type Relation struct {
	NonStrict bool
}

func NewRelation() *Relation {
	return &Relation{
		NonStrict: false,
	}
}

type DAG struct {
	Name      string
	Builtins  *Builtins
	Jobs      map[string]*Job
	InDegrees map[string]int
	Relations map[string]map[string]*Relation
}

type Serializer interface {
	Serialize(dag *DAG) ([]byte, error)
	Deserialize(data []byte) (*DAG, error)
}

func NewDAG(name string) *DAG {
	return &DAG{
		Name:      name,
		Builtins:  NewBuiltins(),
		Jobs:      make(map[string]*Job),
		InDegrees: make(map[string]int),
		Relations: make(map[string]map[string]*Relation),
	}
}

func LoadFromFile(path string) (*DAG, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return LoadFromBytes(data)
}

func LoadFromBytes(data []byte) (*DAG, error) {
	d, err := NewDotSerializer().Deserialize(data)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (this *DAG) ResolveJob(job *Job) error {
	vars := ""
	if v, ok := job.Attrs["vars"]; ok {
		vars = v
	}

	if resolved, err := symbol.Resolve(strings.Trim(vars, "\"'"), this.Builtins.builtins); err != nil {
		return err
	} else {
		for _, stmt := range resolved {
			job.Vars[stmt.Value.(string)] = stmt.Children[0].Value.(string)
		}
	}

	for k, v := range job.Attrs {
		if k == "vars" {
			continue
		}
		nval, err := ApplyVarToString(v, job.Vars)
		if err != nil {
			log.Fatal(err)
			return err
		}
		job.Attrs[k] = nval
	}

	return nil
}

func ApplyVarToString(str string, vars map[string]string) (string, error) {
	for _, token := range varPattern.FindAllStringSubmatch(str, -1) {
		name := token[1]
		if val, ok := vars[name]; ok {
			nval := strings.Replace(str, "${"+name+"}", val, 1)
			log.Debugf("resolve: %s => %s", str, nval)
			str = nval
		} else {
			return "", fmt.Errorf("variable %s is not defined", name)
		}
	}
	return str, nil
}

func (this *DAG) AddJob(job *Job) {
	if _, ok := this.Jobs[job.Name]; ok {
		return
	}
	for _, n := range job.Post {
		if _, ok := this.Jobs[n]; !ok {
			panic(fmt.Errorf("no post job: %s", n))
		}
		this.InDegrees[n] = this.InDegrees[n] + 1
	}
	for _, n := range job.Prev {
		if _, ok := this.Jobs[n]; !ok {
			panic(fmt.Errorf("no prev job: %s", n))
		}
	}
	this.InDegrees[job.Name] = len(job.Prev)
}

func (this *DAG) String() string {
	str := fmt.Sprintf("dag{\n\tname=%s,\n", this.Name)
	for name, job := range this.Jobs {
		str += fmt.Sprintf("\tjob{\n")
		str += fmt.Sprintf("\t\tname=%s,\n", name)
		str += fmt.Sprintf("\t\tattrs{\n")
		for attr, value := range job.Attrs {
			str += fmt.Sprintf("\t\t\t%s=%s,\n", attr, value)
		}
		str += fmt.Sprintf("\t\t}\n")
		str += fmt.Sprintf("\t\tvars{\n")
		for v, value := range job.Vars {
			str += fmt.Sprintf("\t\t\t%s=%s,\n", v, value)
		}
		str += fmt.Sprintf("\t\t}\n")
		str += fmt.Sprintf("\t}\n")
	}
	str += fmt.Sprintf("}\n")
	return str
}

//===================================================================
// Private
//===================================================================

var varPattern = regexp.MustCompile("\\$\\{(.*?)\\}")
