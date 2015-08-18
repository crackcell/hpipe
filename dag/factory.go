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
)

//===================================================================
// Public APIs
//===================================================================

type DAGFactory struct {
	loader Loader
}

func NewDAGFactory(loader Loader) *DAGFactory {
	return &DAGFactory{
		loader: loader,
	}
}

func (this *DAGFactory) CreateDAGFromFile(path string) (*DAG, error) {
	d, err := this.loader.LoadFile(path)
	if err != nil {
		return nil, err
	}
	// TODO: parse variables
	return d, nil
}

func (this *DAGFactory) CreateDAGFromBytes(data []byte) (*DAG, error) {
	d, err := this.loader.LoadBytes(data)
	if err != nil {
		return nil, err
	}
	// TODO: parse variables
	return d, nil
}

//===================================================================
// Private
//===================================================================
