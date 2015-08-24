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
	"time"
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
	//fmt.Printf("src: %s\n", src)
	res := getTestEvalResult(src, t)
	if res == nil {
		return
	}
	check := ast.NewInt(3)
	if !res.Equals(check) {
		t.Error(fmt.Errorf("%s=%d", src, res.Value.(int)))
		return
	}
	//fmt.Printf("res: %v\n", res)
}

func TestEvalMinus(t *testing.T) {
	src := "1-2"
	//fmt.Printf("src: %s\n", src)
	res := getTestEvalResult(src, t)
	if res == nil {
		return
	}
	check := ast.NewInt(-1)
	if !res.Equals(check) {
		t.Error(fmt.Errorf("%s=%d", src, res.Value.(int)))
		return
	}
	//fmt.Printf("res: %v\n", res)
}

func TestEvalTimes(t *testing.T) {
	src := "3*9"
	//fmt.Printf("src: %s\n", src)
	res := getTestEvalResult(src, t)
	if res == nil {
		return
	}
	check := ast.NewInt(27)
	if !res.Equals(check) {
		t.Error(fmt.Errorf("%s=%d", src, res.Value.(int)))
		return
	}
	//fmt.Printf("res: %v\n", res)
}

func TestEvalDate(t *testing.T) {
	src := "${YYYYMMDD}"
	//fmt.Printf("src: %s\n", src)
	res := getTestEvalResult(src, t)
	if res == nil {
		return
	}
	check := ast.NewDate(time.Now(), "YYYYMMDD")
	if !res.Equals(check) {
		t.Error(fmt.Errorf("%v=%v", res.Value, check.Value))
		return
	}
	//fmt.Printf("res: %v\n", res)
}

func TestEvalDateDurationExtAdd(t *testing.T) {
	src := "${YYYYMMDD}+2*$day"
	//fmt.Printf("src: %s\n", src)
	res := getTestEvalResult(src, t)
	if res == nil {
		return
	}
	check := ast.NewDate(time.Now().AddDate(0, 0, 2), "YYYYMMDD")
	if !res.Equals(check) {
		t.Error(fmt.Errorf("%v=%v", res.Value, check.Value))
		return
	}
	//fmt.Printf("res: %v\n", res)
}

func TestEvalDateDurationExtMinus(t *testing.T) {
	src := "${YYYYMMDD}-2*$day"
	//fmt.Printf("src: %s\n", src)
	res := getTestEvalResult(src, t)
	if res == nil {
		return
	}
	check := ast.NewDate(time.Now().AddDate(0, 0, -2), "YYYYMMDD")
	if !res.Equals(check) {
		t.Error(fmt.Errorf("%v=%v", res.Value, check.Value))
		return
	}
	//fmt.Printf("res: %v\n", res)
}
