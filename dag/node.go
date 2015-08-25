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

type NodeType int

func ParseNodeType(typ string) NodeType {
	switch typ {
	case "dummy":
		return DummyNode
	case "hadoop_streaming":
		return HadoopStreamingNode
	default:
		return UnknownNode
	}
}

const (
	DummyNode NodeType = iota
	HadoopStreamingNode
	NodeCount
	UnknownNode
)

type Node struct {
	Name  string
	Type  NodeType
	Attrs Attrs
	Prev  []string
	Post  []string
	Vars  map[string]string
}

func NewNode() *Node {
	return &Node{
		Attrs: NewAttrs(),
		Prev:  []string{},
		Post:  []string{},
		Vars:  make(map[string]string),
	}
}

func (this *Node) Assign(node *Node) {
	this.Name = node.Name
	this.Type = node.Type
	this.Attrs = node.Attrs
	this.Prev = node.Prev
	this.Post = node.Post
}

func (this *Node) String() string {
	return fmt.Sprintf("Node{name=%s,attrs=%v,prev=%v,post=%v}",
		this.Name, this.Attrs, this.Prev, this.Post)
}

//===================================================================
// Private
//===================================================================
