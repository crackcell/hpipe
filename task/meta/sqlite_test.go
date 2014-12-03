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

package meta

import (
	"../../yafl/parser"
	_ "code.google.com/p/gosqlite/sqlite3"
	"fmt"
	"testing"
)

func TestSqlite(t *testing.T) {
	p := parser.NewXMLParser()
	f, err := p.ParseFile("flow1.xml", "../../test")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(f.DebugString())

	db, err := NewSqliteDB("./test.db")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	db.SaveFlow(f)
	db.FetchFlow(f)
}
