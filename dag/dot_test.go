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
	//"fmt"
	"testing"
)

func TestDotLoaderLoadFile(t *testing.T) {
	d := NewDotLoader()
	_, err := d.LoadFile("../examples/wordcount/wordcount.dot")
	if err != nil {
		t.Error(err)
		return
	}
	//fmt.Println(g)
}
