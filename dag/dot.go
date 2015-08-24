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
 * @file dot.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Mon Aug 17 22:01:26 2015
 *
 **/

package dag

import (
	dot "github.com/awalterschulze/gographviz"
	dotparser "github.com/awalterschulze/gographviz/parser"
	"io/ioutil"
)

//===================================================================
// Public APIs
//===================================================================

type DotLoader struct{}

func NewDotLoader() *DotLoader {
	return &DotLoader{}
}

func (this *DotLoader) LoadFile(path string) (*DAG, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return this.LoadBytes(data)
}

func (this *DotLoader) LoadBytes(data []byte) (*DAG, error) {
	ast, err := dotparser.ParseBytes(data)
	if err != nil {
		return nil, err
	}
	graph := dot.NewGraph()
	dot.Analyse(ast, graph)

	p := NewDAG(graph.Name)

	for src, dests := range graph.Edges.SrcToDsts {
		for dest, _ := range dests {

			if orig, ok := p.Nodes[src]; !ok {
				n := dotNameToDAGNode(graph, src)
				n.Post = append(n.Post, dest)
				p.Nodes[src] = n
			} else {
				orig.Post = append(orig.Post, dest)
			}
			if orig, ok := p.Nodes[dest]; !ok {
				n := dotNameToDAGNode(graph, dest)
				n.Prev = append(n.Prev, src)
				p.Nodes[dest] = n
			} else {
				orig.Prev = append(orig.Prev, src)
			}

			if _, ok := p.InDegrees[src]; !ok {
				p.InDegrees[src] = 0
			}
			if orig, ok := p.InDegrees[dest]; !ok {
				p.InDegrees[dest] = 1
			} else {
				p.InDegrees[dest] = orig + 1
			}

		}
	}

	return p, nil
}

//===================================================================
// Private
//===================================================================

func dotToDAGNode(node *dot.Node) *Node {
	p := NewNode()
	p.Name = node.Name
	p.Attrs = dotToDAGAttrs(node.Attrs)
	p.Type = getNodeTypeFromAttrs(p.Attrs)
	return p
}

func dotNameToDAGNode(graph *dot.Graph, name string) *Node {
	if dotNode, ok := graph.Nodes.Lookup[name]; !ok {
		panic("no corresponding node")
	} else {
		return dotToDAGNode(dotNode)
	}
}

func dotToDAGAttrs(attrs dot.Attrs) Attrs {
	p := NewAttrs()
	for k, v := range attrs {
		p.Set(k, v)
	}
	return p
}

func getNodeTypeFromAttrs(attrs Attrs) NodeType {
	if val, ok := attrs["type"]; !ok {
		return UnknownNode
	} else {
		return ParseNodeType(val)
	}
}
