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
	//"github.com/crackcell/hpipe/config"
	"github.com/crackcell/hpipe/dag"
	"github.com/crackcell/hpipe/log"
)

//===================================================================
// JobExec
//===================================================================

type JobExec interface {
	Run(job *dag.Job) error
	GetJobStatus(job *dag.Job) dag.JobStatus
	CheckJobAttrs(job *dag.Job) bool
}

var JobExecs = map[dag.JobType]JobExec{
	dag.DummyJob:  NewDummyExec(),
	dag.HadoopJob: NewHadoopExec(),
	dag.ShellJob:  NewShellExec(),
}

func GetJobExec(job *dag.Job) (JobExec, error) {
	if e, ok := JobExecs[job.Type]; !ok {
		return nil, fmt.Errorf("unknown job type: %v", job.Type)
	} else {
		return e, nil
	}
}

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

		if err := this.RunQueue(queue); err != nil {
			return err
		}

		for _, job := range queue {
			if job.Status == dag.Finished {
				delete(d.InDegrees, job.Name)
				for _, post := range job.Post {
					d.InDegrees[post] = d.InDegrees[post] - 1
				}
			}
		}
	}

	log.Info("All jobs done")
	return nil
}

func (this *DAGExec) RunQueue(queue []*dag.Job) error {
	for _, job := range queue {
		log.Debugf("run job: %s", job.Name)
		jexec, err := GetJobExec(job)
		if err != nil {
			return err
		}
		if job.Type == dag.DummyJob {
			job.Status = dag.Finished
		} else {
			err := jexec.Run(job)
			if err != nil {
				job.Status = dag.Failed
				return err
			}
			job.Status = jexec.GetJobStatus(job)
		}
	}
	return nil
}

//===================================================================
// Private
//===================================================================
