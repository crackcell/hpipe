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
	"./ast"
	"fmt"
	"testing"
	"time"
)

func TestAll(t *testing.T) {
	c := NewTit()
	c.AddStmt("date1",
		ast.NewDate(time.Date(2014, time.November, 11, 0, 0, 0, 0, time.UTC)))
	c.AddStmt("num1", ast.NewInt64(int64(123)))
	src := "date2=date1+1"
	c.AddSrc(src)
	m, err := c.DoEval()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("src:%v\n", src)
	fmt.Println("output:")
	for k, v := range m {
		fmt.Printf("%s=%v\n", k, v)
	}
}
