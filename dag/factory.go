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
 * @file factory.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Wed Aug 19 00:40:20 2015
 *
 **/

package dag

import (
	//"fmt"
	"github.com/crackcell/hpipe/dag/symbol"
	"io/ioutil"
	"strings"
)

//===================================================================
// Public APIs
//===================================================================

type DAGFactory struct {
	loader Loader
}

func NewFactory() *DAGFactory {
	return &DAGFactory{
		loader: NewDotLoader(),
	}
}

func (this *DAGFactory) CreateDAGFromFile(path string) (*DAG, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return this.CreateDAGFromBytes(data)
}

func (this *DAGFactory) CreateDAGFromBytes(data []byte) (*DAG, error) {
	d, err := this.loader.LoadBytes(data)
	if err != nil {
		return nil, err
	}
	for _, node := range d.Nodes {
		for attr, value := range node.Attrs {
			//fmt.Printf("old: %s\n", value)
			if resolved, err := symbol.Resolve(strings.Trim(value, "\"'")); err != nil {
				return nil, err
			} else {
				node.Attrs[attr] = resolved.Value.(string)
			}
			//fmt.Println("new:", node.Attrs[attr])
		}
	}
	return d, nil
}

//===================================================================
// Private
//===================================================================
