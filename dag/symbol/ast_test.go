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
	src := "$res.name=1+2"
	//fmt.Printf("src: %s\n", src)
	p := parser.NewParser()
	l := lexer.NewLexer([]byte(src))
	a, err := p.Parse(l)
	if err != nil {
		t.Error(err)
		return
	}
	//fmt.Printf("src: %s\n", a)
	if a.([]*ast.Stmt)[0].Type != ast.Operator {
		t.Error(fmt.Errorf("type error"))
	}
}

func TestASTMinus(t *testing.T) {
	src := "$res=1-2"
	//fmt.Printf("src: %s\n", src)
	p := parser.NewParser()
	l := lexer.NewLexer([]byte(src))
	a, err := p.Parse(l)
	if err != nil {
		t.Error(err)
		return
	}
	//fmt.Printf("res: %v\n", a)
	if a.([]*ast.Stmt)[0].Type != ast.Operator {
		t.Error(fmt.Errorf("type error"))
	}
}

func TestASTTimes(t *testing.T) {
	src := "$res=3*9"
	//fmt.Printf("src: %s\n", src)
	p := parser.NewParser()
	l := lexer.NewLexer([]byte(src))
	a, err := p.Parse(l)
	if err != nil {
		t.Error(err)
		return
	}
	//fmt.Printf("res: %v\n", a)
	if a.([]*ast.Stmt)[0].Type != ast.Operator {
		t.Error(fmt.Errorf("type error"))
	}
}

func TestASTDateFormat1(t *testing.T) {
	src := "$res=${date:YYYYMMDD}"
	//fmt.Printf("src: %s\n", src)
	p := parser.NewParser()
	l := lexer.NewLexer([]byte(src))
	a, err := p.Parse(l)
	if err != nil {
		t.Error(err)
		return
	}
	//fmt.Printf("res: %v\n", a)
	if a.([]*ast.Stmt)[0].Type != ast.Operator {
		t.Error(fmt.Errorf("type error"))
	}
}

func TestASTDateFormat2(t *testing.T) {
	src := "$res=${date:YYYYMMDD hh:mm:ss}"
	p := parser.NewParser()
	l := lexer.NewLexer([]byte(src))
	//fmt.Printf("src: %s\n", src)
	a, err := p.Parse(l)
	if err != nil {
		t.Error(err)
		return
	}
	//fmt.Printf("res: %v\n", a)
	if a.([]*ast.Stmt)[0].Type != ast.Operator {
		t.Error(fmt.Errorf("type error"))
	}
}

func TestASTVarAdd(t *testing.T) {
	src := "$res=$var0+1"
	p := parser.NewParser()
	l := lexer.NewLexer([]byte(src))
	a, err := p.Parse(l)
	if err != nil {
		t.Error(err)
		return
	}
	//fmt.Printf("src: %s\nres: %v\n", src, a)
	if a.([]*ast.Stmt)[0].Type != ast.Operator {
		t.Error(fmt.Errorf("type error"))
	}
}

func TestASTVarAddDate(t *testing.T) {
	src := "$res=$var0+${date:YYYYYMMDD hh:mm:ss}"
	p := parser.NewParser()
	l := lexer.NewLexer([]byte(src))
	//fmt.Printf("src: %s\n", src)
	a, err := p.Parse(l)
	if err != nil {
		t.Error(err)
		return
	}
	//fmt.Printf("src: %s\nres: %v\n", src, a)
	if a.([]*ast.Stmt)[0].Type != ast.Operator {
		t.Error(fmt.Errorf("type error"))
	}
}

func TestASTDuration(t *testing.T) {
	src := "$res=10*$day"
	p := parser.NewParser()
	l := lexer.NewLexer([]byte(src))
	//fmt.Printf("src: %s\n", src)
	a, err := p.Parse(l)
	if err != nil {
		t.Error(err)
		return
	}
	//fmt.Printf("res: %v\n", a)
	if a.([]*ast.Stmt)[0].Type != ast.Operator {
		t.Error(fmt.Errorf("type error"))
	}
}

func TestASTString(t *testing.T) {
	src := "$res=\"wordcount1\""
	p := parser.NewParser()
	l := lexer.NewLexer([]byte(src))
	//fmt.Printf("src: %s\n", src)
	a, err := p.Parse(l)
	if err != nil {
		t.Error(err)
		return
	}
	if a.([]*ast.Stmt)[0].Type != ast.Operator {
		t.Error(fmt.Errorf("type error"))
		return
	}
	//fmt.Printf("res: %v\n", a)
}

func TestASTEnv(t *testing.T) {
	src := "$res=${env:HOME}"
	p := parser.NewParser()
	l := lexer.NewLexer([]byte(src))
	//fmt.Printf("src: %s\n", src)
	a, err := p.Parse(l)
	if err != nil {
		t.Error(err)
		return
	}
	if a.([]*ast.Stmt)[0].Type != ast.Operator {
		t.Error(fmt.Errorf("type error"))
		return
	}
	//fmt.Printf("res: %v\n", a)
}
