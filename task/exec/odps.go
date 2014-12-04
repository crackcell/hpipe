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
 * @file odps.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Fri Nov 28 16:57:33 2014
 *
 **/

package exec

import (
	"../../log"
	"../../yafl/ast"
	//"bufio"
	"fmt"
	//"os/exec"
	"time"
)

//===================================================================
// Public APIs
//===================================================================

func NewODPSExec() Exec {
	return new(ODPSExec)
}

type ODPSExec struct {
	prop map[string]string
	cmd  string
}

func (this *ODPSExec) Setup(prop map[string]string) error {
	for _, p := range propNames {
		if _, ok := prop[p]; !ok {
			log.Fatalf("%s not found", p)
			return fmt.Errorf("%s not found", p)
		}
	}
	this.prop = prop
	return nil
}

func (this *ODPSExec) Run(job *ast.Job) (string, error) {
	log.Debug(job.Name + " - odpscmd" +
		" -u " + this.getEnv("odps_access_id") +
		" -p " + this.getEnv("odps_access_key") +
		" --project=" + this.getEnv("odps_project") +
		" --endpoint=" + this.getEnv("odps_endpoint") +
		" -e \"" + this.cmd + "\"")
	time.Sleep(10000)
	return ast.DONE, nil
	/*
		cmd := exec.Command("odpscmd",
			"-u "+this.getEnv("odps_access_id"),
			"-p "+this.getEnv("odps_access_key"),
			"--project="+this.getEnv("odps_project"),
			"--endpoint="+this.getEnv("odps_endpoint"),
			"-e \""+this.cmd+"\"")
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
			return ast.FAIL, err
		}
		stderr, err := cmd.StderrPipe()
		if err != nil {
			log.Fatal(err)
			return ast.FAIL, err
		}
		if err := cmd.Start(); err != nil {
			log.Fatal(err)
			return ast.FAIL, err
		}
		errscanner := bufio.NewScanner(stderr)
		for errscanner.Scan() {
			log.Info(errscanner.Text())
		}
		if err := errscanner.Err(); err != nil {
			log.Fatalf("reading standard input: %v", err)
		}
		outscanner := bufio.NewScanner(stdout)
		for outscanner.Scan() {
			log.Info(outscanner.Text())
		}
		if err := outscanner.Err(); err != nil {
			log.Fatalf("reading standard input: %v", err)
		}
		return ast.DONE, nil
	*/
}

//===================================================================
// Private
//===================================================================

var propNames []string = []string{
	"odps_access_id",
	"odps_access_key",
	"odps_project",
	"odps_endpoint",
}

func (this *ODPSExec) getEnv(key string) string {
	v, ok := this.prop[key]
	if !ok {
		panic(fmt.Errorf("no prop for %s", key))
	}
	return v
}
