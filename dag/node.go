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
 * @file node.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Tue Aug 18 04:25:23 2015
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

type Node struct {
	Name string
	Type int
	Attr map[string]string
}

type Nodes struct {
	Lookup map[string]*Node
	Nodes  []*Node
}

func NewNodes() *Nodes {
	p := new(Nodes)
	p.Lookup = make(map[string]*Node)
	p.Nodes = make([]*Node, 0)
	return p
}

func (this *Nodes) Add(node *Node) {
	if n, ok := this.Lookup[node.Name]; ok {
		n.Assign(node)
		return
	}
	this.Lookup[node.Name] = node
	this.Nodes = append(this.Nodes, node)
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

func (this *Node) Assign(node *Node) {
	this.Name = node.Name
	this.Type = node.Type
	this.Attr = node.Attr
}

func (this *Node) String() string {
	return fmt.Sprintf("Node{name=%s,attr=%v}", this.Name, this.Attr)
}

//===================================================================
// Private
//===================================================================
