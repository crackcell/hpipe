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
	"github.com/crackcell/hpipe/dag/symbol/ast"
	"github.com/crackcell/hpipe/dag/symbol/eval"
	"github.com/crackcell/hpipe/dag/symbol/lexer"
	"github.com/crackcell/hpipe/dag/symbol/parser"
	"github.com/crackcell/hpipe/util"
	"time"
)

//===================================================================
// Public APIs
//===================================================================

var Builtins map[string]*ast.Expr = map[string]*ast.Expr{
	"gmtdate": &ast.Expr{
		Type:  ast.Date,
		Value: util.DateT(time.Now(), "YYYYMMDD"),
		Prop: map[string]interface{}{
			"format": "YYYYMMDD",
		},
	},
	"bizdate": &ast.Expr{
		Type:  ast.Date,
		Value: util.DateT(time.Now().AddDate(0, 0, -1), "YYYYMMDD"),
		Prop: map[string]interface{}{
			"format": "YYYYMMDD",
		},
	},
}

type Resolver struct {
}

func NewResolver() *Resolver {
	return &Resolver{}
}

func (this *Resolver) Resolve(src string) (*ast.Expr, error) {
	p := parser.NewParser()
	l := lexer.NewLexer([]byte(src))
	a, err := p.Parse(l)
	if err != nil {
		return nil, err
	}
	e := eval.NewEval()
	e.Builtins = Builtins
	return e.Evaluate(a.(*ast.Expr))
}

//===================================================================
// Private
//===================================================================
