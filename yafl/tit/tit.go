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
 * @file tit.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Mon Nov 17 15:32:54 2014
 *
 **/

package tit

import (
	"./ast"
	"./eval"
	"./lexer"
	"./parser"
	_ "fmt"
	"strings"
)

//===================================================================
// Public APIs
//===================================================================

type Tit struct {
	Src []string
}

func NewTit() *Tit {
	return &Tit{}
}

func (this *Tit) AddStmt(k string, v *ast.Stmt) {
	this.Src = append(this.Src, k+"="+v.ValueString())
}

func (this *Tit) AddStmtMap(m map[string]*ast.Stmt) {
	for k, v := range m {
		this.AddStmt(k, v)
	}
}

func (this *Tit) AddSrc(src ...string) {
	for _, s := range src {
		this.Src = append(this.Src, strings.Trim(s, ";"))
	}
}

func (this *Tit) AddSrcMap(m map[string]string) {
	for k, v := range m {
		this.AddSrc(k, v)
	}
}

func (this *Tit) DoEval() (map[string]*ast.Stmt, error) {
	src := this.assembleStmt()
	if len(src) == 0 {
		return make(map[string]*ast.Stmt), nil
	}

	p := parser.NewParser()
	l := lexer.NewLexer([]byte(src))
	a, err := p.Parse(l)
	if err != nil {
		panic(err)
	}
	e := eval.NewEval()
	return e.DoEval(a.([]*ast.Stmt))
}

//===================================================================
// Private
//===================================================================

func (this *Tit) assembleStmt() string {
	var src string
	for _, s := range this.Src {
		src += s + ";"
	}
	src = strings.Trim(src, ";")
	//log.Debugf("src: %s", src)
	return src
}
