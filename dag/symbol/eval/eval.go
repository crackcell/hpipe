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
 * @file eval.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Thu Aug 20 19:32:45 2015
 *
 **/

package eval

import (
	"fmt"
	"github.com/crackcell/hpipe/dag/symbol/ast"
	"github.com/crackcell/hpipe/util"
	"time"
)

//===================================================================
// Public APIs
//===================================================================

type Eval struct {
	Root     *ast.Expr
	Builtins map[string]*ast.Expr
}

func NewEval() *Eval {
	return &Eval{
		Builtins: make(map[string]*ast.Expr),
	}
}

func (this *Eval) Evaluate(expr *ast.Expr) (*ast.Expr, error) {
	return this.evalExpr(expr)
}

//===================================================================
// Private
//===================================================================

func (this *Eval) evalExpr(expr *ast.Expr) (*ast.Expr, error) {
	switch expr.Type {
	case ast.Operator:
		op1, err := this.evalExpr(expr.Children[0])
		if err != nil {
			return nil, err
		}
		op2, err := this.evalExpr(expr.Children[1])
		if err != nil {
			return nil, err
		}
		return this.evalOperator(op1, op2, expr.Value.(string))
	case ast.Int64:
		return expr, nil
	case ast.Date:
		return this.evalDate(expr)
	case ast.Var:
		return this.evalVar(expr)
	default:
		return nil, fmt.Errorf("unknown expression type: %d", expr.Type)
	}
}

func (this *Eval) evalOperator(op1 *ast.Expr, op2 *ast.Expr, opstr string) (*ast.Expr, error) {

	if op1.Type == ast.Int64 && op2.Type == ast.Int64 {
		var res int64
		a := op1.Value.(int64)
		b := op2.Value.(int64)
		switch opstr {
		case "+":
			res = a + b
		case "-":
			res = a - b
		case "*":
			res = a * b
		case "/":
			res = a / b
		default:
			return nil, fmt.Errorf("invalid operator: %s", opstr)
		}
		return &ast.Expr{
			Type:  ast.Int64,
			Value: res,
		}, nil
	}

	return nil, nil
}

func (this *Eval) evalDate(expr *ast.Expr) (*ast.Expr, error) {
	format := expr.Value.(string)
	expr.Value = util.DateT(time.Now(), format)
	expr.Prop["format"] = format
	return expr, nil
}

func (this *Eval) evalVar(expr *ast.Expr) (*ast.Expr, error) {
	name := expr.Value.(string)
	expr.Prop["name"] = name
	if v, ok := this.Builtins[name]; !ok {
		return nil, fmt.Errorf("invalid var: %s", name)
	} else {
		return v, nil
	}
}
