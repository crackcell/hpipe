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
	"github.com/crackcell/hpipe/util/time"
	stdtime "time"
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
		Builtins: map[string]*ast.Expr{
			// Date
			"gmtdate": ast.NewDate(stdtime.Now(), "YYYYMMDD"),
			"bizdate": ast.NewDate(stdtime.Now().AddDate(0, 0, -1), "YYYYMMDD"),
			// Duration
			"year":   ast.NewDurationB(1, 0, 0),
			"month":  ast.NewDurationB(0, 1, 0),
			"day":    ast.NewDurationB(0, 0, 1),
			"hour":   ast.NewDurationA(stdtime.Hour),
			"minute": ast.NewDurationA(stdtime.Minute),
			"second": ast.NewDurationA(stdtime.Second),
		},
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
	case ast.Int:
		return expr, nil
	case ast.Date:
		return this.evalDate(expr)
	case ast.DurationA:
		return expr, nil
	case ast.DurationB:
		return expr, nil
	case ast.Var:
		return this.evalVar(expr)
	default:
		return nil, fmt.Errorf("unknown expression type: %d", expr.Type)
	}
}

func (this *Eval) evalOperator(op1 *ast.Expr, op2 *ast.Expr, opstr string) (*ast.Expr, error) {

	if op1.Type == ast.Int && op2.Type == ast.Int {
		var res int
		a := op1.Value.(int)
		b := op2.Value.(int)
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
			return nil, fmt.Errorf("invalid operator between Integers: %s", opstr)
		}
		return ast.NewInt(res), nil
	}

	if op1.Type == ast.Date && op2.Type == ast.DurationA {
		var ret stdtime.Time
		old := op1.Prop["time"].(stdtime.Time)
		duration := op2.Prop["time"].(stdtime.Duration)
		switch opstr {
		case "+":
			ret = old.Add(duration)
		case "-":
			ret = old.Add(-1 * duration)
		default:
			return nil, fmt.Errorf("invalid operator between Date and Duration: %s", opstr)
		}
		return ast.NewDate(ret, op1.Prop["format"].(string)), nil
	}

	if op1.Type == ast.Date && op2.Type == ast.DurationB {
		var ret stdtime.Time
		old := op1.Prop["time"].(stdtime.Time)
		year := op2.Prop["year"].(int)
		month := op2.Prop["month"].(int)
		day := op2.Prop["day"].(int)
		switch opstr {
		case "+":
			ret = old.AddDate(year, month, day)
		case "-":
			ret = old.AddDate(-1*year, -1*month, -1*day)
		default:
			return nil, fmt.Errorf("invalid operator between Date and Duration: %s", opstr)
		}
		return ast.NewDate(ret, op1.Prop["format"].(string)), nil
	}

	if op1.Type == ast.Int && op2.Type == ast.DurationA {
		if opstr != "*" {
			return nil, fmt.Errorf("invalid operator between Int and Duration: %s", opstr)
		}
		return ast.NewDurationA(stdtime.Duration(op1.Value.(int)) * op2.Prop["time"].(stdtime.Duration)), nil
	}

	if op1.Type == ast.Int && op2.Type == ast.DurationB {
		if opstr != "*" {
			return nil, fmt.Errorf("invalid operator between Int and Duration: %s", opstr)
		}
		n := op1.Value.(int)
		return ast.NewDurationB(
			n*op2.Prop["year"].(int),
			n*op2.Prop["month"].(int),
			n*op2.Prop["day"].(int)), nil
	}

	return nil, fmt.Errorf("invalid operation pair type: %s and %s",
		op1.Type, op2.Type)
}

func (this *Eval) evalDate(expr *ast.Expr) (*ast.Expr, error) {
	format := expr.Value.(string)
	t := stdtime.Now()
	expr.Value = time.Format(t, format)
	expr.Prop["time"] = t
	expr.Prop["format"] = format
	return expr, nil
}

func (this *Eval) evalDuration(expr *ast.Expr) (*ast.Expr, error) {
	format := expr.Value.(string)
	d, err := stdtime.ParseDuration(format)
	if err != nil {
		return nil, fmt.Errorf("invalid duration: %s", format)
	}
	expr.Value = d
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
