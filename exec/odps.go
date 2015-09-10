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
 * @file odps.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Tue Aug 25 18:28:05 2015
 *
 **/

package exec

import (
	"fmt"
	"github.com/crackcell/hpipe/config"
	"github.com/crackcell/hpipe/dag"
	"github.com/crackcell/hpipe/log"
	"strings"
)

//===================================================================
// Public APIs
//===================================================================

type OdpsExec struct {
	endpoint  string
	project   string
	accessId  string
	accessKey string
}

func NewOdpsExec() *OdpsExec {
	return &OdpsExec{}
}

func (this *OdpsExec) Setup() error {
	if len(config.OdpsEndpoint) != 0 {
		this.endpoint = config.OdpsEndpoint
	} else {
		msg := fmt.Errorf("odps endpoint not set")
		log.Fatal(msg)
		return msg
	}
	if len(config.OdpsProject) != 0 {
		this.project = config.OdpsProject
	} else {
		msg := fmt.Errorf("odps project not set")
		log.Fatal(msg)
		return msg
	}
	if len(config.OdpsAccessID) != 0 {
		this.accessId = config.OdpsAccessID
	} else {
		msg := fmt.Errorf("odps accessid not set")
		log.Fatal(msg)
		return msg
	}
	if len(config.OdpsAccessKey) != 0 {
		this.accessKey = config.OdpsAccessKey
	} else {
		msg := fmt.Errorf("odps accesskey not set")
		log.Fatal(msg)
		return msg
	}
	return nil
}

func (this *OdpsExec) Run(job *dag.Job) error {
	if !checkJobAttr(job, []string{"command", "output"}) {
		msg := "invalid job"
		log.Error(msg)
		return fmt.Errorf(msg)
	}

	// !!!VERY IMPORTANT!!!
	// Many other operations relay on this TrimRight.
	job.Attrs["output"] = strings.TrimRight(job.Attrs["output"], "/")

	args := this.genCmdArgs(job)
	log.Debugf("CMD: odpscmd %s", strings.Join(args, " "))
	retcode, err := cmdExec(job.Name, "odpscmd", args...)
	if err != nil {
		job.Status = dag.Failed
		return err
	}
	if retcode != 0 {
		job.Status = dag.Failed
		return fmt.Errorf("script failed: %d", retcode)
	}
	job.Status = dag.Finished
	return nil
}

//===================================================================
// Private
//===================================================================

func (this *OdpsExec) genCmdArgs(job *dag.Job) []string {
	args := []string{}

	cmd := ""

	args = append(args, "--endpoint")
	args = append(args, this.endpoint)

	args = append(args, "--project")
	args = append(args, this.project)

	args = append(args, "-u")
	args = append(args, this.accessId)

	args = append(args, "-p")
	args = append(args, this.accessKey)

	for k, v := range job.Attrs {
		if _, ok := dag.JobReservedAttrs[k]; ok {
			continue
		}
		cmd += fmt.Sprintf("set %s=%s;", k, v)
	}
	cmd += strings.Trim(job.Attrs["command"], ";") + ";"

	args = append(args, "-e")
	args = append(args, cmd)

	return args
}
