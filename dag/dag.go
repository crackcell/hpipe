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

const (
	DUMMY_NODE = iota
	HADOOP_STREAMING_NODE
	DAG_NODE_COUNT
)

type DAG struct {
	Name string
	Root *Node
}

type Node struct {
	Name string
	Type int
	Attr map[string]string
}

func NewDAG(name string) *DAG {
	p := new(DAG)
	return p
}

func (this *DAG) String() string {
	return fmt.Sprintf("DAG{name=%s}")
}

func NewNode(name string, typ int) (*Node, error) {
	if typ < DUMMY_NODE || typ >= DAG_NODE_COUNT {
		return nil, InvalidNodeType
	}
	p := new(Node)
	p.Name = name
	p.Type = typ
	return p, nil
}

func (this *Node) String() string {
	return fmt.Sprintf("Node{name=%s,attr=%v}", this.Name, this.Attr)
}

//===================================================================
// Private
//===================================================================
