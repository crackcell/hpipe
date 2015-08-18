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
	Name           string
	LookupNode     map[string]*Node
	LookupIndegree map[string]int
}

func NewDAG(name string) *DAG {
	return &DAG{
		Name:           name,
		LookupNode:     make(map[string]*Node),
		LookupIndegree: make(map[string]int),
	}
}

func (this *DAG) String() string {
	return fmt.Sprintf(
		"DAG{name=%s,node=%v,indegree=%v}",
		this.Name, this.LookupNode, this.LookupIndegree)
}

//===================================================================
// Private
//===================================================================
