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
)

//===================================================================
// Public APIs
//===================================================================

func Resolve(src string, builtins map[string]*ast.Stmt) ([]*ast.Stmt, error) {
	e := eval.NewEval()
	if len(src) == 0 {
		return e.Evaluate([]*ast.Stmt{}, builtins)
	}
	p := parser.NewParser()
	l := lexer.NewLexer([]byte(src))
	a, err := p.Parse(l)
	if err != nil {
		return nil, err
	}
	return e.Evaluate(a.([]*ast.Stmt), builtins)
}

//===================================================================
// Private
//===================================================================
