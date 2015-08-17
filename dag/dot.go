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
	//"fmt"
	dot "github.com/awalterschulze/gographviz"
	dotparser "github.com/awalterschulze/gographviz/parser"
	"io/ioutil"
)

//===================================================================
// Public APIs
//===================================================================

type DotLoader struct{}

func NewDotLoader() *DotLoader {
	p := new(DotLoader)
	return p
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

	d := NewDAG(graph.Name)
	return d, nil
}

//===================================================================
// Private
//===================================================================
