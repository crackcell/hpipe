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
	Duration    // natively supported by time.Duration: hour, minute, second
	DurationExt // extended duration: year, month, day
	String
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
	case Duration:
		return "Duration"
	case DurationExt:
		return "DurationExt"
	case String:
		return "String"
	case Operator:
		return "Operator"
	}
	return "unknown"
}

func (this *Expr) Equals(other *Expr) bool {
	if this.Type != other.Type {
		return false
	}
	switch this.Type {
	case Var:
		if this.Value.(string) != other.Value.(string) {
			return false
		}
	case Int:
		if this.Prop["value"].(int) != other.Prop["value"].(int) {
			return false
		}
	case Date:
		if this.Value.(string) != other.Value.(string) {
			return false
		}
	case Duration:
		if this.Value.(string) != other.Value.(string) {
			return false
		}
	case DurationExt:
		if this.Value.(string) != other.Value.(string) {
			return false
		}
	case String:
		if this.Value.(string) != other.Value.(string) {
			return false
		}
	case Operator:
		if this.Value.(string) != other.Value.(string) {
			return false
		}
	default:
		panic(fmt.Errorf("invalid expression type: %v", this.Type))
	}
	return true
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
		Value: strconv.Itoa(n),
		Prop: map[string]interface{}{
			"value": n,
		},
	}
}

func NewIntFromParser(num string) (*Expr, error) {
	if n, err := strconv.Atoi(num); err != nil {
		return nil, err
	} else {
		return &Expr{
			Type:  Int,
			Value: num,
			Prop: map[string]interface{}{
				"value": n,
			},
		}, nil
	}
}

func NewVarFromParser(lit string) (*Expr, error) {
	return &Expr{
		Type:  Var,
		Value: strings.TrimLeft(lit, "$"),
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
		Value: strings.Trim(lit, "${}"),
		Prop:  make(map[string]interface{}),
	}, nil
}

func NewDuration(d stdtime.Duration) *Expr {
	return &Expr{
		Type:  Duration,
		Value: d.String(),
		Prop: map[string]interface{}{
			"time": d,
		},
	}
}

func NewDurationExt(year, month, day int) *Expr {
	return &Expr{
		Type:  DurationExt,
		Value: fmt.Sprintf("%dY%dM%dD", year, month, day),
		Prop: map[string]interface{}{
			"year":  year,
			"month": month,
			"day":   day,
		},
	}
}

func NewStringFromParser(lit string) (*Expr, error) {
	return &Expr{
		Type:  String,
		Value: lit,
		Prop:  make(map[string]interface{}),
	}, nil
}

//===================================================================
// Private
//===================================================================
