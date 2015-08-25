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
	Nodes     map[string]*Node
	InDegrees map[string]int
}

func NewDAG(name string) *DAG {
	return &DAG{
		Name:      name,
		Nodes:     make(map[string]*Node),
		InDegrees: make(map[string]int),
	}
}

func (this *DAG) String() string {
	str := fmt.Sprintf("dag{\n\tname=%s,\n", this.Name)
	for name, node := range this.Nodes {
		str += fmt.Sprintf("\tnode{\n")
		str += fmt.Sprintf("\t\tname=%s,\n", name)
		str += fmt.Sprintf("\t\tattrs{\n")
		for attr, value := range node.Attrs {
			str += fmt.Sprintf("\t\t\t%s=%s,\n", attr, value)
		}
		str += fmt.Sprintf("\t\t}\n")
		str += fmt.Sprintf("\t\tvars{\n")
		for v, value := range node.Vars {
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
