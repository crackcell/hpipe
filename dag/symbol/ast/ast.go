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
 * @file ast.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Thu Aug 20 21:38:08 2015
 *
 **/

package ast

import (
	"fmt"
	"strconv"
	"strings"
)

//===================================================================
// Public APIs
//===================================================================

type NodeType int

const (
	Var NodeType = iota
	Int64
	Date
	Operator
)

type Expr struct {
	Type     NodeType
	Value    interface{}
	Prop     map[string]interface{}
	Children []*Expr
}

func (this NodeType) String() string {
	switch this {
	case Var:
		return "Var"
	case Int64:
		return "Int64"
	case Date:
		return "Date"
	case Operator:
		return "Operator"
	}
	return "unknown"
}

func (this *Expr) String() string {
	return debugExpr(this, 0)
}

func debugExpr(expr *Expr, depth int) string {
	indent := strings.Repeat("\t", depth)
	var str string
	str += fmt.Sprintf("%sexpr:{\n", indent)
	str += fmt.Sprintf("%s\ttype:%s\n", indent, expr.Type)
	str += fmt.Sprintf("%s\tvalue:%v\n", indent, expr.Value)
	str += fmt.Sprintf("%s\tprop:%v\n", indent, expr.Prop)
	for _, child := range expr.Children {
		str += debugExpr(child, depth+1)
	}
	str += fmt.Sprintf("%s}\n", indent)
	return str
}

func NewOperator(op1, op, op2 interface{}) (*Expr, error) {
	return &Expr{
		Type:  Operator,
		Value: op.(string),
		Prop:  make(map[string]interface{}),
		Children: []*Expr{
			op1.(*Expr),
			op2.(*Expr),
		},
	}, nil
}

func NewInt64FromParser(num string) (*Expr, error) {
	if n, err := strconv.Atoi(num); err != nil {
		return nil, err
	} else {
		return &Expr{
			Type:  Int64,
			Value: int64(n),
			Prop:  make(map[string]interface{}),
		}, nil
	}
}

func NewVarFromParser(id interface{}) (*Expr, error) {
	return &Expr{
		Type:  Var,
		Value: id,
		Prop:  make(map[string]interface{}),
	}, nil
}

func NewBuiltinVarFromParser(lit interface{}) (*Expr, error) {
	return &Expr{
		Type:  Date,
		Value: lit,
		Prop:  make(map[string]interface{}),
	}, nil
}

//===================================================================
// Private
//===================================================================
