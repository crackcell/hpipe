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

			if orig, ok := p.Jobs[src]; !ok {
				n := dotNameToDAGJob(graph, src)
				n.Post = append(n.Post, dest)
				p.Jobs[src] = n
			} else {
				orig.Post = append(orig.Post, dest)
			}
			if orig, ok := p.Jobs[dest]; !ok {
				n := dotNameToDAGJob(graph, dest)
				n.Prev = append(n.Prev, src)
				p.Jobs[dest] = n
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

func dotToDAGJob(job *dot.Node) *Job {
	p := NewJob()
	p.Name = job.Name
	p.Attrs = dotToDAGAttrs(job.Attrs)
	p.Type = getJobTypeFromAttrs(p.Attrs)
	return p
}

func dotNameToDAGJob(graph *dot.Graph, name string) *Job {
	if dotJob, ok := graph.Nodes.Lookup[name]; !ok {
		panic("no corresponding job")
	} else {
		return dotToDAGJob(dotJob)
	}
}

func dotToDAGAttrs(attrs dot.Attrs) Attrs {
	p := NewAttrs()
	for k, v := range attrs {
		p.Set(k, v)
	}
	return p
}

func getJobTypeFromAttrs(attrs Attrs) JobType {
	if val, ok := attrs["type"]; !ok {
		return UnknownJob
	} else {
		return ParseJobType(val)
	}
}
