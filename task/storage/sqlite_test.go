/***************************************************************
 *
 * Copyright (c) 2014, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 * Test for sqlite
 *
 * @file sqlite_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Mon Nov 10 19:50:59 2014
 *
 **/

package storage

import (
	"../../yafl/ast"
	"../../yafl/parser"
	_ "code.google.com/p/gosqlite/sqlite3"
	"fmt"
	"testing"
)

var step *ast.Step

func loadFlowFromFile(filename, workdir string) (*ast.Flow, error) {
	f := ast.NewFlow()
	p := parser.NewXMLParser()
	if step, err := p.ParseStepFromFile(filename, workdir); err != nil {
		return nil, err
	} else {
		f.Entry = step
		return f, nil
	}
}

func TestSqlite(t *testing.T) {
	f, err := loadFlowFromFile("step1.xml", "../../test")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(f.Entry.DebugString())

	db := NewSqliteDB("./test.db")

	if err := db.Open(); err != nil {
		t.Error(err)
	}
	defer db.Close()

	db.SaveFlow(f)

	db.RestoreFlow(f)
}
