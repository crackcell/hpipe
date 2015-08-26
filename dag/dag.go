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
