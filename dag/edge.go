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
 * @file edge.go
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

type Edge struct {
	Src  *Node
	Dest *Node
}

type Edges struct {
	SrcToDest map[*Node]map[*Node]*Edge
	DestToSrc map[*Node]map[*Node]*Edge
	Edges     []*Edge
}

func NewEdges() *Edges {
	p := new(Edges)
	p.SrcToDest = make(map[*Node]map[*Node]*Edge)
	p.DestToSrc = make(map[*Node]map[*Node]*Edge)
	return p
}

func (this *Edges) Add(edge *Edge) {
	if _, ok := this.SrcToDest[edge.Src]; !ok {
		this.SrcToDest[edge.Src] = make(map[*Node]*Edge)
	}
	this.SrcToDest[edge.Src][edge.Dest] = edge

	if _, ok := this.DestToSrc[edge.Dest]; !ok {
		this.DestToSrc[edge.Dest] = make(map[*Node]*Edge)
	}
	this.DestToSrc[edge.Src][edge.Dest] = edge
}

func NewEdge(src *Node, dest *Node) *Edge {
	p := new(Edge)
	p.Src = src
	p.Dest = dest
	return p
}

func (this *Edge) String() string {
	return fmt.Sprintf("%s->%s", this.Src.Name, this.Dest.Name)
}

//===================================================================
// Private
//===================================================================
