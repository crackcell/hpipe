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
 * @file sqlitesaver.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sun Sep  6 23:26:06 2015
 *
 **/

package status

import (
	"database/sql"
	"fmt"
	"github.com/crackcell/hpipe/dag"
	"github.com/crackcell/hpipe/log"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"sync"
)

//===================================================================
// Public APIs
//===================================================================

type StatusTable struct {
	Output string
	Status string
}

type SqliteSaver struct {
	db   *sql.DB
	lock sync.Mutex
}

func NewSqliteSaver(path string) (*SqliteSaver, error) {
	exist := isExist(path)

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if !exist {
		sql := "drop table if exists status;create table status(output string, status string);"
		log.Debugf("sql: %s", sql)
		if _, err := db.Exec(sql); err != nil {
			log.Fatal(err)
			return nil, err
		}
	}

	return &SqliteSaver{
		db: db,
	}, nil
}

func (this *SqliteSaver) GetFlag(job *dag.Job) (dag.JobStatus, error) {
	this.lock.Lock()
	defer this.lock.Unlock()

	sql := fmt.Sprintf(
		"select * from status where output='%s'",
		job.Attrs["output"],
	)
	log.Debugf("sql: %s", sql)
	rows, err := this.db.Query(sql)
	if err != nil {
		log.Fatal(err)
		return dag.UnknownStatus, err
	}
	defer rows.Close()
	for rows.Next() {
		var s StatusTable
		rows.Scan(&s.Output, &s.Status)
		if s := dag.ParseJobStatus(s.Status); s == dag.UnknownStatus {
			return dag.NotStarted, nil
		} else {
			return s, nil
		}
	}
	return dag.NotStarted, nil
}

func (this *SqliteSaver) SetFlag(job *dag.Job) error {
	if err := this.ClearFlag(job); err != nil {
		return err
	}
	this.lock.Lock()
	defer this.lock.Unlock()
	sql := fmt.Sprintf(
		"insert into status(output, status) values('%s', '%s')",
		job.Attrs["output"], job.Status,
	)
	log.Debugf("sql: %s", sql)
	if _, err := this.db.Exec(sql); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (this *SqliteSaver) DeleteFlag(job *dag.Job, status dag.JobStatus) error {
	return this.ClearFlag(job)
}

func (this *SqliteSaver) ClearFlag(job *dag.Job) error {
	this.lock.Lock()
	defer this.lock.Unlock()

	sql := fmt.Sprintf(
		"delete from status where output='%s'",
		job.Attrs["output"],
	)
	log.Debugf("sql: %s", sql)
	if _, err := this.db.Exec(sql); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

//===================================================================
// Private
//===================================================================

func isExist(path string) bool {
	f, err := os.Open(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	f.Close()
	return true
}
