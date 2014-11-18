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
 * @file calc.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Mon Nov 17 15:32:54 2014
 *
 **/

package calc

import (
	"./ast"
	"./eval"
	"./lexer"
	"./parser"
	"strings"
)

//===================================================================
// Public APIs
//===================================================================

type Calc struct {
	SrcPieces []string
}

func NewCalc() *Calc {
	return &Calc{}
}

func (this *Calc) AddVar(k, v string) {
	this.SrcPieces = append(this.SrcPieces, k+"="+v)
}

func (this *Calc) AddVarMap(m map[string]string) {
	for k, v := range m {
		this.AddVar(k, v)
	}
}

func (this *Calc) AddPiece(src string) {
	this.SrcPieces = append(this.SrcPieces, src)
}

func (this *Calc) DoEval() (map[string]string, error) {
	src := this.assembleSrc()

	p := parser.NewParser()
	l := lexer.NewLexer([]byte(src))
	a, err := p.Parse(l)
	if err != nil {
		return nil, err
	}
	e := eval.NewEval()
	output, err := e.DoEval(a.([]*ast.Stmt))
	if err != nil {
		return nil, err
	}
	ret := make(map[string]string)
	for name, stmt := range output {
		ret[name] = stmt.ValueString()
	}
	return ret, nil
}

//===================================================================
// Private
//===================================================================

func (this *Calc) assembleSrc() string {
	var src string
	for _, piece := range this.SrcPieces {
		src += strings.Trim(piece, ";") + ";"
	}
	src = strings.TrimSuffix(src, ";")
	//fmt.Println("src:", src)
	return src
}
