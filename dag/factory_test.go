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
 * @file factory_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Wed Aug 19 00:45:32 2015
 *
 **/

package dag

import (
	"fmt"
	"testing"
)

func TestDAGFactoryCreateDAGFromFile(t *testing.T) {
	f := NewDAGFactory(NewDotLoader())
	d, err := f.CreateDAGFromFile("./test.dot")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(d)
}

func TestDAGFactoryCreateDAGFromBytes(t *testing.T) {
	f := NewDAGFactory(NewDotLoader())
	d, err := f.CreateDAGFromBytes([]byte(`
digraph wordcount_example {
  wordcount1 [
    name="wordcount1"
  ]
  wordcount2 [
    name="wordcount2"
  ]
  wordcount1 -> wordcount2 -> wordcount3
  wordcount1 -> wordcount3
}
`))
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(d)
}
