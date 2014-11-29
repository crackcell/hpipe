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
 * @file sqlite.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Fri Nov 28 18:14:08 2014
 *
 **/

package meta

import (
	_ "../../config"
	"../../log"
	_ "../../util"
	"../../yafl/ast"
	"database/sql"
	_ "database/sql/driver"
	_ "fmt"
)

//===================================================================
// Public APIs
//===================================================================

func NewSqliteDB(path string) (DB, error) {
	ret := new(SqliteDB)
	conn, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	ret.conn = conn
	log.Debugf("open sqlite3 succ: %s", path)
	return ret, nil
}

type SqliteDB struct {
	conn *sql.DB
	tx   *sql.Tx
}

func (this *SqliteDB) Close() error {
	if this.conn != nil {
		return this.conn.Close()
	}
	return nil
}

func (this *SqliteDB) SaveFlow(f *ast.Flow) error {
	tx, err := this.conn.Begin()
	this.tx = tx
	if err != nil {
		log.Fatalf("begin tx failed: %v", err)
		return err
	}
	if err := this.saveStep(f.Entry); err != nil {
		this.tx.Rollback()
		return err
	}
	if err := this.tx.Commit(); err != nil {
		log.Fatalf("tx commit failed: %v", err)
		return err
	}
	return nil
}

//===================================================================
// Private
//===================================================================

func (this *SqliteDB) saveStep(s *ast.Step) error {
	for _, dep := range s.Dep {
		if err := this.saveStep(dep); err != nil {
			log.Fatalf("saveStep failed: %v", err)
			return err
		}
	}
	for _, do := range s.Do {
		if err := this.saveJob(do); err != nil {
			log.Fatalf("saveJob failed: %v", err)
			return err
		}
	}
	return nil
}

func (this *SqliteDB) saveJob(j *ast.Job) error {
	_, err := this.tx.Exec(
		"UPDATE hpipe_task_info SET status=? WHERE instance_id=j",
		j.Status, j.InstanceID,
	)
	if err == nil {
		return nil
	}
	this.createTable()
	_, err = this.tx.Exec(
		"INSERT INTO hpipe_task_info (instance_id, status) VALUES (?,?)",
		j.InstanceID, j.Status,
	)
	log.Fatalf("saveJob failed: %v", err)
	return err
}

func (this *SqliteDB) createTable() error {
	_, err := this.conn.Exec(`CREATE TABLE "hpipe_task_info" (
"instance_id" TEXT NOT NULL,
"status" TEXT NOT NULL,
PRIMARY KEY(taskid));`)
	return err
}
