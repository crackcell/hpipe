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
	"testing"
)

func TestSymbolResolveAll(t *testing.T) {
	res, err := Resolve("$gmtdate-1*$day")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(res)
}
