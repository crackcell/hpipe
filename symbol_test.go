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
 * @file symbol_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Fri Aug 20 22:11:12 2015
 *
 **/

package main

import (
	"fmt"
	"github.com/crackcell/hpipe/dag/symbol"
	"github.com/crackcell/hpipe/dag/symbol/ast"
	"testing"
	stdtime "time"
)

func TestSymbolResolveAll(t *testing.T) {
	var builtins = map[string]*ast.Stmt{
		// Date
		"gmtdate": ast.NewDate(stdtime.Now(), "YYYYMMDD"),
		"bizdate": ast.NewDate(stdtime.Now().AddDate(0, 0, -1), "YYYYMMDD"),
		// Duration
		"year":   ast.NewDurationExt(1, 0, 0),
		"month":  ast.NewDurationExt(0, 1, 0),
		"day":    ast.NewDurationExt(0, 0, 1),
		"hour":   ast.NewDuration(stdtime.Hour),
		"minute": ast.NewDuration(stdtime.Minute),
		"second": ast.NewDuration(stdtime.Second),
	}

	src := "$res=$gmtdate-1*$day"
	ret, err := symbol.Resolve(src, builtins)
	if err != nil {
		t.Error(err)
		return
	}
	res := ret[0]
	check := ast.NewLeftID("res",
		ast.NewDate(stdtime.Now().AddDate(0, 0, -1), "YYYYMMDD"))
	if !res.Equals(check) {
		t.Error(fmt.Errorf("%s=%d", src, res.Value.(string)))
		return
	}
	//fmt.Println(res)
}
