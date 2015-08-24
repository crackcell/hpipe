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
	"github.com/crackcell/hpipe/dag/symbol/ast"
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
	fmt.Printf("src: %s\nres: %v\n", src, a)
}

func TestASTMinus(t *testing.T) {
	src := "1-2"
	p := parser.NewParser()
	l := lexer.NewLexer([]byte(src))
	a, err := p.Parse(l)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("src: %s\nres: %v\n", src, a)
}

func TestASTTimes(t *testing.T) {
	src := "3*9"
	p := parser.NewParser()
	l := lexer.NewLexer([]byte(src))
	a, err := p.Parse(l)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("src: %s\nres: %v\n", src, a)
}

func TestASTDateFormat1(t *testing.T) {
	src := "${YYYYMMDD}"
	p := parser.NewParser()
	l := lexer.NewLexer([]byte(src))
	fmt.Printf("src: %s\n", src)
	a, err := p.Parse(l)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("res: %v\n", a)
	if a.(*ast.Expr).Type != ast.Date {
		t.Error(fmt.Errorf("type error"))
	}
}

func TestASTDateFormat2(t *testing.T) {
	src := "${YYYYMMDD hh:mm:ss}"
	p := parser.NewParser()
	l := lexer.NewLexer([]byte(src))
	fmt.Printf("src: %s\n", src)
	a, err := p.Parse(l)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("res: %v\n", a)
	if a.(*ast.Expr).Type != ast.Date {
		t.Error(fmt.Errorf("type error"))
	}
}

func TestASTVarAdd(t *testing.T) {
	src := "$var0+1"
	p := parser.NewParser()
	l := lexer.NewLexer([]byte(src))
	a, err := p.Parse(l)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("src: %s\nres: %v\n", src, a)
}

func TestASTVarAddDate(t *testing.T) {
	src := "$var0+${YYYYYMMDD hh:mm:ss}"
	p := parser.NewParser()
	l := lexer.NewLexer([]byte(src))
	a, err := p.Parse(l)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("src: %s\nres: %v\n", src, a)
}

func TestASTDuration(t *testing.T) {
	src := "10*$day"
	p := parser.NewParser()
	l := lexer.NewLexer([]byte(src))
	fmt.Printf("src: %s\n", src)
	a, err := p.Parse(l)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("res: %v\n", a)
}
