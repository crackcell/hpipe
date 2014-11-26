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
 * @file varbin.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sun Nov 16 23:01:40 2014
 *
 **/

package ast

import (
	"fmt"
	_ "regexp"
	"strconv"
	"strings"
	"time"
)

//===================================================================
// Public APIs
//===================================================================

const (
	LeftID   = "leftid"
	RightID  = "rightid"
	String   = "string"
	Int64    = "int64"
	Date     = "date"
	Operator = "operator"
)

type (
	StmtList []*Stmt
	Stmt     struct {
		Type     string
		Children []*Stmt
		Value    interface{}
	}
)

func (this *Stmt) DebugString() string {
	return fmt.Sprintf("stmt:{type:%s,value:%s}", this.Type, this.Value)
}

func (this *Stmt) ValueString() string {
	switch this.Type {
	case "int64":
		return strconv.Itoa(int(this.Value.(int64)))
	case "date":
		return (this.Value.(time.Time)).Format("20060102")
	default:
		return this.Value.(string)
	}
}

func GetStmtListDebugString(stmtList []*Stmt) string {
	var str string
	str += fmt.Sprintln("stmtlist{\n")
	for _, stmt := range stmtList {
		str += debugStringStmt(stmt, 1)
	}
	str += fmt.Sprintln("}\n")
	return str
}

func debugStringStmt(stmt *Stmt, depth int) string {
	indent := strings.Repeat("\t", depth)
	var str string
	str += fmt.Sprintf("%sstmt:{\n", indent)
	str += fmt.Sprintf("%s\ttype:%s\n", indent, stmt.Type)
	str += fmt.Sprintf("%s\tvalue:%v\n", indent, stmt.Value)
	for _, child := range stmt.Children {
		str += debugStringStmt(child, depth+1)
	}
	str += fmt.Sprintf("%s}\n", indent)
	return str
}

func NewStmtList(stmt interface{}) ([]*Stmt, error) {
	return []*Stmt{stmt.(*Stmt)}, nil
}

func AppendStmtList(stmtList, stmt interface{}) ([]*Stmt, error) {
	return append(stmtList.([]*Stmt), stmt.(*Stmt)), nil
}

func NewOperator(node1, op, node2 interface{}) (*Stmt, error) {
	var typ string
	opstr := op.(string)
	if opstr == "=" || opstr == "+" || opstr == "-" || opstr == "*" ||
		opstr == "/" {
		typ = Operator
	} else {
		return nil, ErrUnknownOperator
	}
	return &Stmt{Type: typ, Value: opstr, Children: []*Stmt{node1.(*Stmt),
		node2.(*Stmt)}}, nil
}

func NewDate(d interface{}) *Stmt {
	return &Stmt{Type: Date, Value: d.(time.Time)}
}

func NewString(val interface{}) *Stmt {
	return &Stmt{Type: String, Value: val.(string)}
}

func NewValueFromParser(num interface{}) (*Stmt, error) {
	val := strings.TrimSuffix(strings.TrimPrefix(num.(string), "\""), "\"")
	if IsTime(val) {
		return &Stmt{Type: Date, Value: StrToTime(val)}, nil
	} else {
		return &Stmt{Type: String, Value: val}, nil
	}
}

func NewInt64(num interface{}) *Stmt {
	return &Stmt{Type: Int64, Value: num.(int64)}
}

func NewInt64FromParser(num interface{}) (*Stmt, error) {
	if n, err := strconv.Atoi(num.(string)); err != nil {
		return nil, err
	} else {
		return &Stmt{Type: Int64, Value: int64(n)}, nil
	}
}

func NewLeftID(id interface{}) *Stmt {
	return &Stmt{Type: LeftID, Value: id.(string)}
}

func NewRightID(id interface{}) *Stmt {
	return &Stmt{Type: RightID, Value: id.(string)}
}

//===================================================================
// Private
//===================================================================
