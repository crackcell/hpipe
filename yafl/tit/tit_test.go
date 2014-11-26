/***************************************************************
 *
 * Copyright (c) 2014, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 * Test for tit.go
 *
 * @file tit_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Tue Nov 18 13:01:42 2014
 *
 **/

package tit

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	c := NewTit()
	c.AddVar("date1", "20141111")
	c.AddVar("num1", "10")
	src := "date1=date1+1"
	c.AddPiece(src)
	m, err := c.DoEval()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("src:\t%v\n", src)
	fmt.Printf("output:\t%v\n", m)
}
