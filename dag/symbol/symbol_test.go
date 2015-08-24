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

package symbol

import (
	"fmt"
	"github.com/crackcell/hpipe/dag/symbol/ast"
	"testing"
	"time"
)

func TestSymbolResolveAll(t *testing.T) {
	src := "$gmtdate-1*$day"
	res, err := Resolve(src)
	if err != nil {
		t.Error(err)
		return
	}
	check := ast.NewDate(time.Now().AddDate(0, 0, -1), "YYYYMMDD")
	if !res.Equals(check) {
		t.Error(fmt.Errorf("%s=%d", src, res.Value.(string)))
		return
	}
	//fmt.Println(res)
}
