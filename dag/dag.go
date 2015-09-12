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

		if resolved, err := symbol.Resolve(strings.Trim(vars, "\"'")); err != nil {
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
			for _, token := range varPattern.FindAllStringSubmatch(v, -1) {
				name := token[1]
				if val, ok := job.Vars[name]; ok {
					nval := strings.Replace(job.Attrs[k], "${"+name+"}", val, -1)
					log.Debugf("resolve: %s => %s", v, nval)
					job.Attrs[k] = nval
				} else {
					msg := fmt.Sprintf("variable %s is not defined", name)
					log.Error(msg)
					return nil, fmt.Errorf(msg)
				}
			}
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

//===================================================================
// Private
//===================================================================

var varPattern = regexp.MustCompile("\\$\\{(.*?)\\}")
