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

package storage

import (
	"../../log"
	"../../yafl/ast"
	_ "code.google.com/p/gosqlite/sqlite3"
	"database/sql"
	"os"
)

//===================================================================
// Public APIs
//===================================================================

func NewSqliteDB(path string) DB {
	ret := new(SqliteDB)
	ret.path = path
	return ret
}

func (this *SqliteDB) Open() error {
	isInit := true
	if _, err := os.Stat(this.path); os.IsNotExist(err) {
		isInit = false
	}
	conn, err := sql.Open("sqlite3", this.path)
	if err != nil {
		log.Fatal(err)
		return err
	}
	this.conn = conn
	if !isInit {
		if err := this.init(); err != nil {
			return err
		}
	}
	log.Debugf("open: %s", this.path)
	return nil
}

type SqliteDB struct {
	path string
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
	return this.walkFlow(f, this.saveJob)
}

func (this *SqliteDB) RestoreFlow(f *ast.Flow) (*ast.Flow, error) {
	err := this.walkFlow(f, this.restoreJob)
	return f, err
}

func (this *SqliteDB) walkFlow(f *ast.Flow, fn func(j *ast.Job) error) error {
	tx, err := this.conn.Begin()
	this.tx = tx
	if err != nil {
		log.Fatalf("begin tx failed: %v", err)
		return err
	}
	if err := this.walkStep(f.Entry, fn); err != nil {
		this.tx.Rollback()
		return err
	}
	if err := this.tx.Commit(); err != nil {
		log.Fatalf("commit tx failed: %v", err)
		return err
	}
	return nil
}

//===================================================================
// Private
//===================================================================

func (this *SqliteDB) init() error {
	if err := this.createTable(); err != nil {
		return err
	}
	return nil
}

func (this *SqliteDB) createTable() error {
	_, err := this.conn.Exec(`CREATE TABLE "hpipe_task_info" (
"instance_id" TEXT NOT NULL,
"status" TEXT NOT NULL,
PRIMARY KEY(instance_id));`)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (this *SqliteDB) walkStep(s *ast.Step, f func(j *ast.Job) error) error {
	for _, dep := range s.Dep {
		if err := this.walkStep(dep, f); err != nil {
			return err
		}
	}
	for _, do := range s.Do {
		if err := f(do); err != nil {
			return err
		}
	}
	return nil
}

func (this *SqliteDB) saveJob(j *ast.Job) error {
	_, err := this.tx.Exec(
		"DELETE FROM hpipe_task_info WHERE instance_id=?",
		j.InstanceID,
	)
	if err != nil {
		log.Fatal(err)
		return err
	}
	_, err = this.tx.Exec(
		"INSERT INTO hpipe_task_info (instance_id, status) VALUES (?,?)",
		j.InstanceID, j.Status,
	)
	log.Debugf("%s=%s", j.InstanceID, j.Status)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (this *SqliteDB) restoreJob(j *ast.Job) error {
	var status string
	err := this.tx.QueryRow(
		"SELECT status FROM hpipe_task_info WHERE instance_id=?",
		j.InstanceID,
	).Scan(&status)
	switch {
	case err == sql.ErrNoRows:
		return nil
	case err != nil:
		log.Fatal(err)
		return err
	default:
		log.Debugf("%s=%s", j.InstanceID, status)
		j.Status = status
		return nil
	}
}
