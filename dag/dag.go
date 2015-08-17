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
	Name  string
	Nodes *Nodes
	Edges *Edges
}

func NewDAG(name string) *DAG {
	p := new(DAG)
	p.Name = name
	return p
}

func (this *DAG) String() string {
	return fmt.Sprintf("DAG{name=%s}", this.Name)
}

//===================================================================
// Private
//===================================================================
