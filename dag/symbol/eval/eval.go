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
)

//===================================================================
// Public APIs
//===================================================================

type Eval struct {
	Root    *ast.Expr
	Context map[string]*ast.Expr
}

func NexEval() *Eval {
	return &Eval{
		Context: make(map[string]*ast.Expr),
	}
}

func (this *Eval) SetContext(context map[string]*ast.Expr) {
	this.Context = context
}

func (this *Eval) Evaluate(expr *ast.Expr) error {
	return this.evalExpr(expr)
}

//===================================================================
// Private
//===================================================================

func (this *Expr) evalExpr(expr *ast.Expr) (*Expr, error) {
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
		return this.evalDate(expr.Children[0])
	case ast.Var:
		return this.evalVar(expr.Children[0])
	default:
		return fmt.Errorf("unknown expression type: %d", expr.Type)
	}
}

func (this *Expr) evalOperator(op1 *ast.Expr, op2 *ast.Expr,
	opstr string) (*Expr, error) {

	if op1.Type == ast.Int64 && op2.Type == ast.Int64 {
		return &Expr{
			Type:  ast.Int64,
			Value: op1.Value.(int64) + op2.Value.(int64),
		}, nil
	}

	return nil
}

func (this *Expr) evalDate(expr *ast.Expr) (*Expr, error) {
	fmt.Println("date:", expr.Value.(string))
	return nil, nil
}

func (this *Expr) evalVar(expr *ast.Expr) (*Expr, error) {
	fmt.Println("var:", expr.Value.(string))
	return nil, nil
}
