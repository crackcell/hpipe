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
 * @file sched.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Tue Aug 25 18:09:11 2015
 *
 **/

package sched

import (
	"fmt"
	"github.com/crackcell/hpipe/config"
	"github.com/crackcell/hpipe/dag"
	"github.com/crackcell/hpipe/exec"
	"github.com/crackcell/hpipe/log"
	"github.com/crackcell/hpipe/status"
	"github.com/crackcell/hpipe/util"
	"strings"
	"sync"
)

//===================================================================
// Sched
//===================================================================

type Sched struct {
	exec    map[dag.JobType]exec.Exec
	tracker *status.StatusTracker
	failing map[string]int
	failed  map[string]int
}

func NewSched(tracker *status.StatusTracker) (*Sched, error) {
	e := map[dag.JobType]exec.Exec{
		dag.DummyJob:  exec.NewDummyExec(),
		dag.ScriptJob: exec.NewScriptExec(),
	}

	if config.Hadoop {
		e[dag.HadoopJob] = exec.NewHadoopExec()
	}
	if config.Hive {
		e[dag.HiveJob] = exec.NewHiveExec()
	}

	if config.Odps {
		e[dag.OdpsJob] = exec.NewOdpsExec()
	}

	for _, jexec := range e {
		if err := jexec.Setup(); err != nil {
			return nil, err
		}
	}

	return &Sched{
		exec:    e,
		tracker: tracker,
		failing: make(map[string]int),
		failed:  make(map[string]int),
	}, nil
}

func (this *Sched) Run(d *dag.DAG) error {
	if err := this.checkDAG(d); err != nil {
		log.Fatal(err)
		return err
	}

	queue := this.genRunQueue(d)
	for len(queue) != 0 {

		if err := this.runQueue(queue, d); err != nil {
			return err
		}

		for _, job := range queue {
			switch job.Status {
			case dag.Finished:
				this.markJobFinished(job, d)
			case dag.Failed:
				log.Errorf("job %s failed", job.Name)
				if n, ok := this.failing[job.Name]; !ok {
					this.failing[job.Name] = 1
				} else {
					this.failing[job.Name] = n + 1
				}
			}
		}

		queue = this.genRunQueue(d)
	}

	util.LogLines(strings.Trim(this.tracker.String(), "\n"), log.Info)

	if len(this.failed) == 0 {
		log.Info("All jobs done")
		return nil
	} else {
		log.Errorf("some job failed")
		return fmt.Errorf("some job failed")
	}
}

//===================================================================
// Private
//===================================================================

func (this *Sched) genRunQueue(d *dag.DAG) []*dag.Job {
	queue := []*dag.Job{}
	for name, in := range d.InDegrees {
		job, ok := d.Jobs[name]
		if !ok {
			panic(fmt.Errorf("panic: no corresponding job"))
		}
		if in == 0 && job.Status != dag.Finished &&
			job.Status != dag.Started &&
			this.failing[job.Name] < config.MaxRetry {
			queue = append(queue, job)
		}
		if this.failing[job.Name] >= config.MaxRetry {
			log.Errorf("job %s reaches max retry times: %d",
				job.Name, config.MaxRetry)
			this.failed[job.Name] = config.MaxRetry
		}
	}
	return queue
}

func (this *Sched) runQueue(queue []*dag.Job, d *dag.DAG) error {
	var wg sync.WaitGroup
	for _, job := range queue {
		wg.Add(1)
		go func(job *dag.Job) {
			// !!! All shared objects need to be thread-safe !!!
			defer wg.Done()

			log.Infof("run job: %s", job.Name)
			if job.Type == dag.DummyJob {
				job.Status = dag.Finished
			} else {
				jexec, err := this.getExec(job)
				if err != nil {
					panic(err)
				}
				if s, err := this.tracker.GetStatus(job); err != nil {
					panic(err)
				} else {
					job.Status = s
				}
				log.Debugf("check job status: %s -> %s", job.Name, job.Status)

				switch job.Status {
				case dag.Finished:
					log.Infof("job is already finished, skip: %s", job.Name)
					return
				case dag.Started:
					log.Warnf("job is already started: %s", job.Name)
					return
				}

				this.tracker.SetStatus(job, dag.Started)
				d.Builtins.SetJobReport(this.tracker.ToJson())
				if err := d.ResolveJob(job); err != nil {
					log.Error(err)
					job.Status = dag.Failed
				}
				if err = jexec.Run(job); err != nil {
					log.Error(err)
					job.Status = dag.Failed
				}
				this.tracker.SetStatus(job, job.Status)

				log.Debugf("check job status: %s -> %s", job.Name, job.Status)
			}
		}(job)
	}
	wg.Wait()
	return nil
}

func (this *Sched) getExec(job *dag.Job) (exec.Exec, error) {
	if e, ok := this.exec[job.Type]; !ok {
		return nil, fmt.Errorf("no vailid executor for job type: %v", job.Type)
	} else {
		return e, nil
	}
}

func (this *Sched) markJobFinished(job *dag.Job, d *dag.DAG) {
	job.Status = dag.Finished
	for _, post := range job.Post {
		in := d.InDegrees[post]
		if in != 0 {
			d.InDegrees[post] = in - 1
		}
	}
}

func (this *Sched) checkDAG(d *dag.DAG) error {
	for _, job := range d.Jobs {
		if _, ok := this.exec[job.Type]; !ok {
			return fmt.Errorf("no vailid executor for job type: %v", job.Type)
		}
	}
	return nil
}
