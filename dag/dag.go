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
	"github.com/crackcell/hpipe/dag/symbol/ast"
	"github.com/crackcell/hpipe/log"
	"io/ioutil"
	"regexp"
	"strings"
	stdtime "time"
)

//===================================================================
// Public APIs
//===================================================================

type DAG struct {
	Name      string
	Jobs      map[string]*Job
	InDegrees map[string]int
}

func NewDAG(name string) *DAG {
	return &DAG{
		Name:      name,
		Jobs:      make(map[string]*Job),
		InDegrees: make(map[string]int),
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
	d, err := NewDotLoader().LoadBytes(data)
	if err != nil {
		return nil, err
	}
	for _, job := range d.Jobs {
		vars := ""
		if v, ok := job.Attrs["vars"]; ok {
			vars = v
		}

		if resolved, err := symbol.Resolve(strings.Trim(vars, "\"'"), builtins); err != nil {
			return nil, err
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
				return nil, err
			}
			job.Attrs[k] = nval
		}
	}
	return d, nil
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

//===================================================================
// Private
//===================================================================

var varPattern = regexp.MustCompile("\\$\\{(.*?)\\}")

var builtins = map[string]*ast.Stmt{
	// Date
	"gmtdate": ast.NewDate(stdtime.Now(), "YYYYMMDD"),
	"bizdate": ast.NewDate(stdtime.Now().AddDate(0, 0, -1), "YYYYMMDD"),
	// Duration
	"year":   ast.NewDurationExt(1, 0, 0),
	"month":  ast.NewDurationExt(0, 1, 0),
	"day":    ast.NewDurationExt(0, 0, 1),
	"hour":   ast.NewDuration(stdtime.Hour),
	"minute": ast.NewDuration(stdtime.Minute),
	"second": ast.NewDuration(stdtime.Second),
}
