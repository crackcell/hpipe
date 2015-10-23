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
	"fmt"
	"github.com/crackcell/gotabulate"
	"github.com/crackcell/hpipe/dag"
)

//===================================================================
// Public APIs
//===================================================================

type StatusTracker struct {
	saver  Saver
	status map[string]dag.JobStatus
	order  []string
}

func NewStatusTracker(saver Saver) *StatusTracker {
	return &StatusTracker{
		saver:  saver,
		status: make(map[string]dag.JobStatus),
	}
}

func (this *StatusTracker) String() string {
	table := [][]string{[]string{"Job", "Status"}}
	for _, name := range this.order {
		if v, ok := this.status[name]; ok {
			table = append(table, []string{name, v.String()})
		} else {
			panic(fmt.Errorf("no job in map: %s", name))
		}
	}
	tabulator := gotabulate.NewTabulator()
	tabulator.SetFirstRowHeader(true)
	tabulator.SetFormat("psql")
	return tabulator.Tabulate(table)
}

func (this *StatusTracker) GetStatus(job *dag.Job) (dag.JobStatus, error) {
	if s, err := this.saver.GetFlag(job); err != nil {
		return dag.UnknownStatus, err
	} else {
		if _, ok := this.status[job.Name]; !ok {
			this.order = append(this.order, job.Name)
		}
		this.status[job.Name] = s
		return s, nil
	}
}

func (this *StatusTracker) SetStatus(job *dag.Job, status dag.JobStatus) error {
	if err := this.saver.SetFlag(job, status); err != nil {
		return err
	} else {
		if _, ok := this.status[job.Name]; !ok {
			this.order = append(this.order, job.Name)
		}
		this.status[job.Name] = status
		return nil
	}
}
