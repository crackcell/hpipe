/***************************************************************
 *
 * Copyright (c) 2014, Menglong TAN <tanmenglong@gmail.com>
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
 * @date Mon Nov 17 15:32:54 2014
 *
 **/

package eval

import (
	"../../../log"
	"../ast"
	"fmt"
	"time"
)

//===================================================================
// Public APIs
//===================================================================

type Eval struct {
	context map[string]*ast.Stmt
}

func NewEval() *Eval {
	return &Eval{context: make(map[string]*ast.Stmt)}
}

func (this *Eval) DoEval(stmtList []*ast.Stmt) (map[string]*ast.Stmt, error) {
	return this.evalStmtList(stmtList)
}

//===================================================================
// Private
//===================================================================

func (this *Eval) evalStmtList(stmtList []*ast.Stmt) (map[string]*ast.Stmt, error) {
	for _, stmt := range stmtList {
		_, err := this.evalStmt(stmt)
		if err != nil {
			return nil, err
		}
	}
	return this.context, nil
}

func (this *Eval) evalStmt(stmt *ast.Stmt) (*ast.Stmt, error) {
	switch stmt.Type {
	case "operator":
		s1, err := this.evalStmt(stmt.Children[0])
		if err != nil {
			return nil, err
		}
		s2, err := this.evalStmt(stmt.Children[1])
		if err != nil {
			return nil, err
		}
		switch stmt.Value.(string) {
		case "=":
			this.equal(s1, s2)
			return nil, nil
		case "+":
			return this.evalOperator(s1, s2, "+")
		case "-":
			return this.evalOperator(s1, s2, "-")
		case "*":
			return this.evalOperator(s1, s2, "*")
		case "/":
			return this.evalOperator(s1, s2, "/")
		default:
			err := fmt.Errorf(ErrFmtUnknownOperator, stmt.Value.(string))
			log.Fatal(err)
			return nil, err
		}
	case "leftid":
		return stmt, nil
	case "rightid":
		val, ok := this.context[stmt.Value.(string)]
		if !ok {
			err := fmt.Errorf(ErrFmtUndefinedVar, stmt.Value.(string))
			log.Fatal(err)
			return nil, err
		}
		return val, nil
	case "string":
		return stmt, nil
	case "int64":
		return stmt, nil
	case "date":
		return stmt, nil
	default:
		log.Fatal(ast.ErrUnknownOperator)
		return nil, ast.ErrUnknownOperator
	}
}

func (this *Eval) equal(s1, s2 *ast.Stmt) {
	this.context[s1.Value.(string)] = s2
}

func (this *Eval) evalOperator(s1, s2 *ast.Stmt, op string) (*ast.Stmt, error) {
	// can't conduct operation between values
	if s1.Type == "value" || s2.Type == "value" {
		return nil, ErrInvalidOperation
	}

	// can't conduct times and divide between time
	if (s1.Type == "date" || s2.Type == "date") && (op == "*" || op == "/") {
		return nil, ErrInvalidOperation
	}

	hasTime := false
	var date time.Time
	var num int64
	if s1.Type == "date" {
		if s2.Type == "date" {
			return nil, ErrInvalidOperation
		}
		hasTime = true
		date = s1.Value.(time.Time)
		num = s2.Value.(int64)
	} else {
		if s2.Type == "date" {
			if s1.Type == "date" {
				return nil, ErrInvalidOperation
			}
			hasTime = true
			date = s2.Value.(time.Time)
			num = s1.Value.(int64)
		}
	}

	if !hasTime {
		n1 := s1.Value.(int64)
		n2 := s2.Value.(int64)

		var ret int64
		switch op {
		case "+":
			ret = n1 + n2
		case "-":
			ret = n1 - n2
		case "*":
			ret = int64(n1 * n2)
		case "/":
			ret = int64(n1 / n2)
		default:
			return nil, ErrInvalidOperation
		}
		return ast.NewInt64(ret), nil
	}

	for i := int64(0); i < num; i++ {
		date = date.Add(time.Hour * 24)
	}
	return ast.NewDate(date), nil
}
