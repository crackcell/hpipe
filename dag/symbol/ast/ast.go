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
	"github.com/crackcell/hpipe/util/time"
	"strconv"
	"strings"
	stdtime "time"
)

//===================================================================
// Public APIs
//===================================================================

type NodeType int

const (
	Var NodeType = iota
	Int
	Date
	DurationA // natively supported by time.Duration: hour, minute, second
	DurationB // extended duration: year, month, day
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
	case Int:
		return "Int"
	case Date:
		return "Date"
	case DurationA:
		return "DurationA"
	case DurationB:
		return "DurationB"
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

func NewOperatorFromParser(op1 *Expr, op string, op2 *Expr) (*Expr, error) {
	return &Expr{
		Type:  Operator,
		Value: op,
		Prop:  make(map[string]interface{}),
		Children: []*Expr{
			op1,
			op2,
		},
	}, nil
}

func NewInt(n int) *Expr {
	return &Expr{
		Type:  Int,
		Value: int(n),
		Prop:  make(map[string]interface{}),
	}
}

func NewIntFromParser(num string) (*Expr, error) {
	if n, err := strconv.Atoi(num); err != nil {
		return nil, err
	} else {
		return &Expr{
			Type:  Int,
			Value: int(n),
			Prop:  make(map[string]interface{}),
		}, nil
	}
}

func NewVarFromParser(lit string) (*Expr, error) {
	return &Expr{
		Type:  Var,
		Value: lit,
		Prop:  make(map[string]interface{}),
	}, nil
}

func NewDate(t stdtime.Time, format string) *Expr {
	return &Expr{
		Type:  Date,
		Value: time.Format(t, format),
		Prop: map[string]interface{}{
			"time":   t,
			"format": format,
		},
	}
}

func NewDateFromParser(lit string) (*Expr, error) {
	return &Expr{
		Type:  Date,
		Value: lit,
		Prop:  make(map[string]interface{}),
	}, nil
}

func NewDurationA(d stdtime.Duration) *Expr {
	return &Expr{
		Type:  DurationA,
		Value: d.String(),
		Prop: map[string]interface{}{
			"time": d,
		},
	}
}

func NewDurationB(year, month, day int) *Expr {
	return &Expr{
		Type:  DurationB,
		Value: fmt.Sprintf("%dY%dM%dD", year, month, day),
		Prop: map[string]interface{}{
			"year":  year,
			"month": month,
			"day":   day,
		},
	}
}

//===================================================================
// Private
//===================================================================
