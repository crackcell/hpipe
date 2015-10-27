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
 * @file tracker.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sun Sep  6 23:26:06 2015
 *
 **/

package status

import (
	"encoding/json"
	"github.com/crackcell/gotabulate"
	"github.com/crackcell/hpipe/dag"
)

//===================================================================
// Public APIs
//===================================================================

type StatusTracker struct {
	saver  Saver
	status map[string]dag.JobStatus
	Fails  map[string]int
}

func NewStatusTracker(saver Saver) *StatusTracker {
	return &StatusTracker{
		saver:  saver,
		status: make(map[string]dag.JobStatus),
		Fails:  make(map[string]int),
	}
}

func (this *StatusTracker) String() string {
	table := [][]string{[]string{"Job", "Status"}}
	for name, stat := range this.status {
		table = append(table, []string{name, stat.String()})
	}
	tabulator := gotabulate.NewTabulator()
	tabulator.SetFirstRowHeader(true)
	tabulator.SetFormat("psql")
	return tabulator.Tabulate(table)
}

func (this *StatusTracker) ToJson() string {
	table := map[string]string{}
	for name, stat := range this.status {
		table[name] = stat.String()
	}
	if b, err := json.Marshal(table); err != nil {
		return ""
	} else {
		return string(b)
	}
}

func (this *StatusTracker) GetStatus(job *dag.Job) (dag.JobStatus, error) {
	if s, err := this.saver.GetFlag(job); err != nil {
		return dag.UnknownStatus, err
	} else {
		this.status[job.Name] = s
		return s, nil
	}
}

func (this *StatusTracker) SetStatus(job *dag.Job) error {
	if err := this.saver.SetFlag(job); err != nil {
		return err
	} else {
		this.status[job.Name] = job.Status
		return nil
	}
}
