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
 * @file eval_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Fri Aug 20 21:50:22 2015
 *
 **/

package symbol

import (
	"fmt"
	"github.com/crackcell/hpipe/dag/symbol/ast"
	"github.com/crackcell/hpipe/dag/symbol/eval"
	"github.com/crackcell/hpipe/dag/symbol/lexer"
	"github.com/crackcell/hpipe/dag/symbol/parser"
	"testing"
)

func getTestEvalResult(src string, t *testing.T) *ast.Expr {
	p := parser.NewParser()
	l := lexer.NewLexer([]byte(src))
	a, err := p.Parse(l)
	if err != nil {
		t.Error(err)
		return nil
	}
	e := eval.NewEval()
	res, err := e.Evaluate(a.(*ast.Expr))
	if err != nil {
		t.Error(err)
		return nil
	}
	return res
}

func TestEvalAdd(t *testing.T) {
	src := "1+2"
	res := getTestEvalResult(src, t)
	if res == nil {
		return
	}
	if res.Value.(int) != 3 {
		t.Error(fmt.Errorf("%s=%d", src, res.Value.(int)))
		return
	}
	fmt.Printf("src: %s\nres: %v", src, res)
}

func TestEvalMinus(t *testing.T) {
	src := "1-2"
	res := getTestEvalResult(src, t)
	if res == nil {
		return
	}
	if res.Value.(int) != -1 {
		t.Error(fmt.Errorf("%s=%d", src, res.Value.(int)))
		return
	}
	fmt.Printf("src: %s\nres: %v", src, res)
}

func TestEvalTimes(t *testing.T) {
	src := "3*9"
	res := getTestEvalResult(src, t)
	if res == nil {
		return
	}
	if res.Value.(int) != 27 {
		t.Error(fmt.Errorf("%s=%d", src, res.Value.(int)))
		return
	}
	fmt.Printf("src: %s\nres: %v", src, res)
}

func TestEvalDate(t *testing.T) {
	src := "${YYYYMMDD hh:mm:ss}"
	res := getTestEvalResult(src, t)
	if res == nil {
		return
	}
	fmt.Printf("src: %s\nres: %v", src, res)
}

func TestEvalDateDurationAAdd(t *testing.T) {
	src := "${YYYYMMDD hh:mm:ss}+2*$hour"
	fmt.Printf("src: %s\n", src)
	res := getTestEvalResult(src, t)
	if res == nil {
		return
	}
	fmt.Printf("res: %v", res)
}

func TestEvalDateDurationAMinus(t *testing.T) {
	src := "${YYYYMMDD hh:mm:ss}-2*$hour"
	fmt.Printf("src: %s\n", src)
	res := getTestEvalResult(src, t)
	if res == nil {
		return
	}
	fmt.Printf("res: %v", res)
}

func TestEvalDateDurationBAdd(t *testing.T) {
	src := "${YYYYMMDD}+2*$day"
	fmt.Printf("src: %s\n", src)
	res := getTestEvalResult(src, t)
	if res == nil {
		return
	}
	fmt.Printf("res: %v", res)
}

func TestEvalDateDurationBMinus(t *testing.T) {
	src := "${YYYYMMDD}-2*$day"
	fmt.Printf("src: %s\n", src)
	res := getTestEvalResult(src, t)
	if res == nil {
		return
	}
	fmt.Printf("res: %v", res)
}
