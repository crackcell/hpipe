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
 * @file var.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Wed Aug 19 11:28:39 2015
 *
 **/

package symbol

import (
//"fmt"
)

//===================================================================
// Public APIs
//===================================================================

var BuiltinSymbols map[string]string = map[string]string{
	"gmtdate": "yyyymmdd",
	"bizdate": "yyyymmdd-1",
}

type SymbolParser struct {
}

func NewSymbolParser() *SymbolParser {
	return &SymbolParser{}
}

//===================================================================
// Private
//===================================================================
