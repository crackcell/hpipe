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
 * @file ast_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Thu Aug 20 20:25:26 2015
 *
 **/

package symbol

import (
	"fmt"
	"github.com/crackcell/hpipe/dag/symbol/lexer"
	"github.com/crackcell/hpipe/dag/symbol/parser"
	"testing"
)

func TestASTAdd(t *testing.T) {
	src := "1+2"
	p := parser.NewParser()
	l := lexer.NewLexer([]byte(src))
	a, err := p.Parse(l)
	if err != nil {
		t.Error(t)
		return
	}
	fmt.Println(a)
}

func TestASTVarMinus(t *testing.T) {
	src := "$num-10"
	p := parser.NewParser()
	l := lexer.NewLexer([]byte(src))
	a, err := p.Parse(l)
	if err != nil {
		t.Error(t)
		return
	}
	fmt.Println(a)
}

func TestASTDateAdd(t *testing.T) {
	src := "${yyyymmdd}-10*$day"
	p := parser.NewParser()
	l := lexer.NewLexer([]byte(src))
	a, err := p.Parse(l)
	if err != nil {
		t.Error(t)
		return
	}
	fmt.Println(a)
}
