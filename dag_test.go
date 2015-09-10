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
 * @file dag_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Wed Aug 19 00:45:32 2015
 *
 **/

package main

import (
	//"fmt"
	"github.com/crackcell/hpipe/dag"
	"io/ioutil"
	"testing"
)

func TestDAGFactoryCreateDAGFromFile(t *testing.T) {
	f := dag.NewFactory()
	_, err := f.CreateDAGFromFile("examples/wordcount/wordcount.dot")
	if err != nil {
		t.Error(err)
		return
	}
	//fmt.Println(d)
}

func TestDAGFactoryCreateDAGFromBytes(t *testing.T) {
	data, err := ioutil.ReadFile("examples/wordcount/wordcount.dot")
	if err != nil {
		t.Error(err)
		return
	}
	f := dag.NewFactory()
	_, err = f.CreateDAGFromBytes(data)
	if err != nil {
		t.Error(err)
		return
	}
	//fmt.Println(d)
}
