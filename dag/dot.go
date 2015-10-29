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
	"fmt"
	dot "github.com/awalterschulze/gographviz"
	dotparser "github.com/awalterschulze/gographviz/parser"
	"github.com/crackcell/hpipe/log"
	"github.com/crackcell/hpipe/util"
	"strings"
)

//===================================================================
// Public APIs
//===================================================================

type DotSerializer struct{}

func NewDotSerializer() *DotSerializer {
	return &DotSerializer{}
}

func (this *DotSerializer) Serialize(dag *DAG) ([]byte, error) {
	// TODO
	return nil, nil
}

func (this *DotSerializer) Deserialize(data []byte) (*DAG, error) {
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

	for _, edge := range graph.Edges.Edges {
		if v, ok := edge.Attrs["nonstrict"]; ok {
			nonstrict, err := util.StringToBool(strings.Trim(v, "\""))
			if err != nil {
				log.Error(err)
				return nil, err
			}
			if len(edge.Src) != 0 && len(edge.Dst) != 0 {
				r := NewRelation()
				r.NonStrict = nonstrict
				if _, ok := p.Relations[edge.Src]; !ok {
					p.Relations[edge.Src] = make(map[string]*Relation)
				}
				p.Relations[edge.Src][edge.Dst] = r
			}
		}
	}

	return p, nil
}

//===================================================================
// Private
//===================================================================

func dotNameToDAGJob(graph *dot.Graph, name string) *Job {
	dotJob, ok := graph.Nodes.Lookup[name]
	if !ok {
		panic(fmt.Errorf("no corresponding node"))
	}

	job, err := dotToDAGJob(dotJob)
	if err != nil {
		panic(err)
	}
	return job
}

func dotToDAGJob(node *dot.Node) (*Job, error) {
	job := NewJob()
	job.Name = node.Name

	job.Attrs = dotToDAGAttrs(node.Attrs)
	if !job.ValidateAttr([]string{"type"}) {
		return nil, log.ErrorErrf(
			"invalid job '%s', attributes missing: %v",
			job.Name,
			[]string{"type"})
	}

	job.Type = ParseJobType(job.Attrs["type"])
	if job.Type == UnknownJob {
		return nil, log.ErrorErrf("unknown job type: %s for %s", job.Attrs["type"], job.Name)
	}

	return job, nil
}

func dotToDAGAttrs(attrs dot.Attrs) Attrs {
	p := NewAttrs()
	for k, v := range attrs {
		p.Set(k, strings.Trim(v, "\""))
	}
	return p
}
