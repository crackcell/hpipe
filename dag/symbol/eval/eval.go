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
	"os"
	"strconv"
	stdtime "time"
)

//===================================================================
// Public APIs
//===================================================================

type Eval struct {
	Root     *ast.Stmt
	Builtins map[string]*ast.Stmt
}

func NewEval() *Eval {
	return &Eval{}
}

func (this *Eval) Evaluate(stmtlist []*ast.Stmt, builtins map[string]*ast.Stmt) ([]*ast.Stmt, error) {
	if builtins != nil {
		this.Builtins = builtins
	} else {
		this.Builtins = make(map[string]*ast.Stmt)
	}
	list := []*ast.Stmt{}
	for _, stmt := range stmtlist {
		if s, err := this.evalStmt(stmt); err != nil {
			return nil, err
		} else {
			list = append(list, s)
		}
	}
	for k, v := range this.Builtins {
		list = append(list, ast.NewLeftID(k, v))
	}
	return list, nil
}

//===================================================================
// Private
//===================================================================

func (this *Eval) evalStmt(stmt *ast.Stmt) (*ast.Stmt, error) {
	switch stmt.Type {
	case ast.Operator:
		op1, err := this.evalStmt(stmt.Children[0])
		if err != nil {
			return nil, err
		}
		op2, err := this.evalStmt(stmt.Children[1])
		if err != nil {
			return nil, err
		}
		return this.evalOperator(op1, op2, stmt.Value.(string))
	case ast.Int:
		return stmt, nil
	case ast.Float:
		return stmt, nil
	case ast.Date:
		return this.evalDate(stmt)
	case ast.Duration:
		return stmt, nil
	case ast.DurationExt:
		return stmt, nil
	case ast.Env:
		return this.evalEnv(stmt)
	case ast.LeftID:
		return stmt, nil
	case ast.RightID:
		return this.evalRightID(stmt)
	case ast.String:
		return stmt, nil
	default:
		return nil, fmt.Errorf("unknown stmt type: %d", stmt.Type)
	}
}

func (this *Eval) evalOperator(op1 *ast.Stmt, op2 *ast.Stmt, opstr string) (*ast.Stmt, error) {

	if op1.Type == ast.Int && op2.Type == ast.Int {
		var res int
		a := op1.Prop["value"].(int)
		b := op2.Prop["value"].(int)
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

	if (op1.Type == ast.Float || op2.Type == ast.Int) && (op1.Type == ast.Int || op2.Type == ast.Float) {
		var res float64
		var a float64
		var b float64
		var err error
		a, err = strconv.ParseFloat(op1.Value.(string), 64)
		if err != nil {
			return nil, fmt.Errorf("invalid float: %s", op1.Value)
		}
		b, err = strconv.ParseFloat(op2.Value.(string), 64)
		if err != nil {
			return nil, fmt.Errorf("invalid float: %s", op2.Value)
		}
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
		return ast.NewFloat(res), nil
	}

	if op1.Type == ast.LeftID {
		if opstr != "=" {
			return nil, fmt.Errorf("invalid operator for LeftID: %s", opstr)
		}
		return ast.NewLeftID(op1.Value.(string), op2), nil
	}

	if op1.Type == ast.Date && op2.Type == ast.Duration {
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

	if op1.Type == ast.Date && op2.Type == ast.DurationExt {
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

	if op1.Type == ast.Int && op2.Type == ast.Duration {
		if opstr != "*" {
			return nil, fmt.Errorf("invalid operator between Int and Duration: %s", opstr)
		}
		n := op1.Prop["value"].(int)
		return ast.NewDuration(stdtime.Duration(n) * op2.Prop["time"].(stdtime.Duration)), nil
	}

	if op1.Type == ast.Int && op2.Type == ast.DurationExt {
		if opstr != "*" {
			return nil, fmt.Errorf("invalid operator between Int and Duration: %s", opstr)
		}
		n := op1.Prop["value"].(int)
		return ast.NewDurationExt(
			n*op2.Prop["year"].(int),
			n*op2.Prop["month"].(int),
			n*op2.Prop["day"].(int)), nil
	}

	return nil, fmt.Errorf("invalid operation pair type: %s and %s",
		op1.Type, op2.Type)
}

func (this *Eval) evalDate(stmt *ast.Stmt) (*ast.Stmt, error) {
	format := stmt.Value.(string)
	t := stdtime.Now()
	stmt.Value = time.Format(t, format)
	stmt.Prop["time"] = t
	stmt.Prop["format"] = format
	return stmt, nil
}

func (this *Eval) evalDuration(stmt *ast.Stmt) (*ast.Stmt, error) {
	format := stmt.Value.(string)
	d, err := stdtime.ParseDuration(format)
	if err != nil {
		return nil, fmt.Errorf("invalid duration: %s", format)
	}
	stmt.Value = d
	stmt.Prop["format"] = format
	return stmt, nil
}

func (this *Eval) evalRightID(stmt *ast.Stmt) (*ast.Stmt, error) {
	name := stmt.Value.(string)
	stmt.Prop["name"] = name
	if v, ok := this.Builtins[name]; !ok {
		return nil, fmt.Errorf("invalid var: %s", name)
	} else {
		return v, nil
	}
}

func (this *Eval) evalEnv(stmt *ast.Stmt) (*ast.Stmt, error) {
	name := stmt.Value.(string)
	stmt.Prop["name"] = name
	env := os.Getenv(name)
	if len(env) == 0 {
		return nil, fmt.Errorf("invalid env: %s", name)
	} else {
		return ast.NewString(env), nil
	}
}
