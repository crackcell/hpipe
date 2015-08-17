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
 * @file dot_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Mon Aug 17 21:17:34 2015
 *
 **/

package dag

import (
	"fmt"
	"github.com/awalterschulze/gographviz"
	"testing"
)

func TestAll(t *testing.T) {
	g, err := gographviz.Read([]byte(`
digraph "example - wordcount" {
  wordcount1 [
    name="wordcount1"
    input="/hpipe/examples/wordcount/input",
    output="/hpipe/examples/wordcount/output1"
  ]
  wordcount2 [
    name="wordcount2"
    input="/hpipe/examples/wordcount/input",
    output="/hpipe/examples/wordcount/output2"
  ]
  wordcount1 -> wordcount2
}
`))
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(g.String())
}
