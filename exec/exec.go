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
 * @file exec.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Tue Aug 25 18:09:11 2015
 *
 **/

package exec

import (
	"fmt"
	"github.com/crackcell/hpipe/config"
	"github.com/crackcell/hpipe/dag"
)

//===================================================================
// DAGExec
//===================================================================

type DAGExec struct {
}

func NewDAGExec() *DAGExec {
	return &DAGExec{}
}

func (this *DAGExec) Run(d *dag.DAG) error {
	for len(d.InDegrees) != 0 {
		queue := []*dag.Job{}
		for name, indegree := range d.InDegrees {
			if indegree == 0 {
				queue = append(queue, d.Jobs[name])
			}
		}

		fmt.Printf("run: ")
		for _, j := range queue {
			fmt.Printf("%s ", j.Name)
		}
		fmt.Println()

		if err := this.RunQueue(queue); err != nil {
			return err
		}

		for _, j := range queue {
			if j.Finished == true {
				delete(d.InDegrees, j.Name)
				for _, post := range j.Post {
					d.InDegrees[post] = d.InDegrees[post] - 1
				}
			}
		}
	}
	return nil
}

func (this *DAGExec) RunQueue(queue []*dag.Job) error {
	for _, j := range queue {
		j.Finished = true
	}
	return nil
}

//===================================================================
// JobExec
//===================================================================

type JobExec interface {
	Submit(d *dag.Job) error
}

func NewJobExec(job *dag.Job) (JobExec, error) {
	switch job.Type {
	case dag.DummyJob:
		// TODO
		return nil, nil
	case dag.HadoopStreamingJob:
		// TODO
		return nil, nil
	case dag.ShellJob:
		fmt.Println("workpath:", config.WorkPath)
		return NewShellExec(config.WorkPath), nil
	default:
		return nil, fmt.Errorf("unknown job type: %s", job.Type)
	}
}

//===================================================================
// Private
//===================================================================
