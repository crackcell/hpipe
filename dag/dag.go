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
	return fmt.Sprintf(
		"DAG{name=%s,nodes=%v,indegrees=%v}",
		this.Name, this.Nodes, this.InDegrees)
}

//===================================================================
// Private
//===================================================================
