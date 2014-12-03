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
 * @file executor.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Fri Nov 28 16:57:33 2014
 *
 **/

package exec

import (
	"../../yafl/ast"
	_ "fmt"
)

//===================================================================
// Public APIs
//===================================================================

type Executor interface {
	Run(job *ast.Job) error
}

func NewODPSExecutor(id, key, project, endpoint, cmd string) Executor {
	ret := new(ODPSExecutor)
	ret.accessId = id
	ret.accessKey = key
	ret.project = project
	ret.endpoint = endpoint
	ret.cmd = cmd
	return ret
}

//===================================================================
// Private
//===================================================================

type ODPSExecutor struct {
	accessId  string
	accessKey string
	project   string
	endpoint  string
	cmd       string
}

func (this *ODPSExecutor) Run(job *ast.Job) error {
	cmd := exec.Command("odpscmd",
		"-u "+this.accessId, "-p "+this.accessKey,
		"--project="+this.project, "--endpoint="+this.endpoint,
		"-e "+this.cmd)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
		return err
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
		return err
	}
	errscanner := bufio.NewScanner(stderr)
	for errscanner.Scan() {
		log.Info(errscanner.Text()) // Println will add back the final '\n'
	}
	if err := errscanner.Err(); err != nil {
		log.Errorf("read from pipe failed: %v", err)
	}
	outscanner := bufio.NewScanner(stdout)
	for outscanner.Scan() {
		log.Info(outscanner.Text()) // Println will add back the final '\n'
	}
	if err := outscanner.Err(); err != nil {
		log.Errorf("read from pipe failed: %v", err)
	}
	return nil
}
