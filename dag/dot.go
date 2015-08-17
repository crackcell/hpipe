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
	"github.com/awalterschulze/gographviz"
)

//===================================================================
// Public APIs
//===================================================================

type DotLoader struct {
}

func (this *DotLoader) LoadFile(path string) (*DAG, error) {
	return nil, nil
}

func (this *DotLoader) LoadBytes(data []byte) (*DAG, error) {
	dot, err := gographviz.Read(data)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

//===================================================================
// Private
//===================================================================
