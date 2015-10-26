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
 * @file builtins.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Mon Oct 26 17:28:27 2015
 *
 **/

package dag

import (
	"github.com/crackcell/hpipe/config"
	"github.com/crackcell/hpipe/dag/symbol/ast"
	"github.com/crackcell/hpipe/log"
	"github.com/crackcell/hpipe/util/time"
	stdtime "time"
)

//===================================================================
// Public APIs
//===================================================================

type Builtins struct {
	builtins map[string]*ast.Stmt
}

func NewBuiltins() *Builtins {
	return &Builtins{builtins: map[string]*ast.Stmt{
		// Date
		"gmtdate": ast.NewDate(stdtime.Now(), "YYYYMMDD"),
		"bizdate": ast.NewDate(stdtime.Now().AddDate(0, 0, -1), "YYYYMMDD"),
		// Duration
		"year":   ast.NewDurationExt(1, 0, 0),
		"month":  ast.NewDurationExt(0, 1, 0),
		"day":    ast.NewDurationExt(0, 0, 1),
		"hour":   ast.NewDuration(stdtime.Hour),
		"minute": ast.NewDuration(stdtime.Minute),
		"second": ast.NewDuration(stdtime.Second),
		// System
		"job_report": ast.NewString("{}"),
	}}
}

func (this *Builtins) SetBizdate(d string) error {
	bizdate, err := time.Parse(d, "YYYYMMDD")
	if err != nil {
		log.Fatalf("invalid bizdate: %s", config.Bizdate)
		return err
	}
	this.builtins["gmtdate"] = ast.NewDate(bizdate.AddDate(0, 0, 1), "YYYYMMDD")
	this.builtins["bizdate"] = ast.NewDate(bizdate, "YYYYMMDD")
	return nil
}

func (this *Builtins) SetJobReport(status string) {
	this.builtins["job_report"] = ast.NewString(status)
}

//===================================================================
// Private
//===================================================================
