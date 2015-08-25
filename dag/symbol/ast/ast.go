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
	LeftID NodeType = iota
	RightID
	Int
	Date
	Duration    // natively supported by time.Duration: hour, minute, second
	DurationExt // extended duration: year, month, day
	String
	Operator
)

type Stmt struct {
	Type     NodeType
	Value    interface{}
	Prop     map[string]interface{}
	Children []*Stmt
}

func (this NodeType) String() string {
	switch this {
	case LeftID:
		return "LeftID"
	case RightID:
		return "RightID"
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

func (this *Stmt) Equals(other *Stmt) bool {
	if this.Type != other.Type {
		return false
	}
	switch this.Type {
	case LeftID:
		if this.Value.(string) != other.Value.(string) ||
			len(this.Children) != len(other.Children) ||
			len(this.Children) == 0 ||
			!this.Children[0].Equals(other.Children[0]) {
			return false
		}
	case RightID:
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

func (this *Stmt) String() string {
	return debugStmt(this, 0)
}

func debugStmt(expr *Stmt, depth int) string {
	indent := strings.Repeat("\t", depth)
	var str string
	str += fmt.Sprintf("%sexpr:{\n", indent)
	str += fmt.Sprintf("%s\ttype:%s\n", indent, expr.Type)
	str += fmt.Sprintf("%s\tvalue:%v\n", indent, expr.Value)
	str += fmt.Sprintf("%s\tprop:%v\n", indent, expr.Prop)
	for _, child := range expr.Children {
		str += debugStmt(child, depth+1)
	}
	str += fmt.Sprintf("%s}\n", indent)
	return str
}

func NewStmtList(stmt *Stmt) []*Stmt {
	return []*Stmt{stmt}
}

func AppendStmtList(stmtlist []*Stmt, stmt *Stmt) []*Stmt {
	return append(stmtlist, stmt)
}

func NewOperatorFromParser(op1 *Stmt, op string, op2 *Stmt) (*Stmt, error) {
	if op != "+" && op != "-" && op != "*" && op != "/" && op != "=" {
		return nil, fmt.Errorf("invalid operator: %s", op)
	}
	return &Stmt{
		Type:  Operator,
		Value: op,
		Prop:  make(map[string]interface{}),
		Children: []*Stmt{
			op1,
			op2,
		},
	}, nil
}

func NewInt(n int) *Stmt {
	return &Stmt{
		Type:  Int,
		Value: strconv.Itoa(n),
		Prop: map[string]interface{}{
			"value": n,
		},
	}
}

func NewIntFromParser(num string) (*Stmt, error) {
	if n, err := strconv.Atoi(num); err != nil {
		return nil, err
	} else {
		return &Stmt{
			Type:  Int,
			Value: num,
			Prop: map[string]interface{}{
				"value": n,
			},
		}, nil
	}
}

func NewLeftID(id string, children ...*Stmt) *Stmt {
	return &Stmt{
		Type:     LeftID,
		Value:    strings.TrimLeft(id, "$"),
		Prop:     make(map[string]interface{}),
		Children: children,
	}
}

func NewLeftIDFromParser(lit string) (*Stmt, error) {
	return &Stmt{
		Type:  LeftID,
		Value: strings.TrimLeft(lit, "$"),
		Prop:  make(map[string]interface{}),
	}, nil
}

func NewRightID(id string) *Stmt {
	return &Stmt{
		Type:  RightID,
		Value: strings.TrimLeft(id, "$"),
		Prop:  make(map[string]interface{}),
	}
}

func NewRightIDFromParser(lit string) (*Stmt, error) {
	return &Stmt{
		Type:  RightID,
		Value: strings.TrimLeft(lit, "$"),
		Prop:  make(map[string]interface{}),
	}, nil
}

func NewDate(t stdtime.Time, format string) *Stmt {
	return &Stmt{
		Type:  Date,
		Value: time.Format(t, format),
		Prop: map[string]interface{}{
			"time":   t,
			"format": format,
		},
	}
}

func NewDateFromParser(lit string) (*Stmt, error) {
	return &Stmt{
		Type:  Date,
		Value: strings.Trim(lit, "${}"),
		Prop:  make(map[string]interface{}),
	}, nil
}

func NewDuration(d stdtime.Duration) *Stmt {
	return &Stmt{
		Type:  Duration,
		Value: d.String(),
		Prop: map[string]interface{}{
			"time": d,
		},
	}
}

func NewDurationExt(year, month, day int) *Stmt {
	return &Stmt{
		Type:  DurationExt,
		Value: fmt.Sprintf("%dY%dM%dD", year, month, day),
		Prop: map[string]interface{}{
			"year":  year,
			"month": month,
			"day":   day,
		},
	}
}

func NewStringFromParser(lit string) (*Stmt, error) {
	return &Stmt{
		Type:  String,
		Value: lit,
		Prop:  make(map[string]interface{}),
	}, nil
}

//===================================================================
// Private
//===================================================================
