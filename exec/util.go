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
 * @file util.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Thu Aug 27 19:53:11 2015
 *
 **/

package exec

import (
	"bufio"
	"fmt"
	"github.com/crackcell/hpipe/dag"
	"github.com/crackcell/hpipe/log"
	"os/exec"
	"syscall"
)

//===================================================================
// Public APIs
//===================================================================

//===================================================================
// Private
//===================================================================

func isInMap(key string, m dag.Attrs) bool {
	_, ok := m[key]
	return ok
}

func checkJobAttr(job *dag.Job, keys []string) bool {
	for _, key := range keys {
		if !isInMap(key, job.Attrs) {
			return false
		}
	}
	return true
}

func cmdExec(jobname, name string, arg ...string) (int, error) {
	cmd := exec.Command(name, arg...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
		return 0, err
	}
	errscanner := bufio.NewScanner(stderr)
	for errscanner.Scan() {
		if len(errscanner.Text()) != 0 {
			log.Infof("<%s> %s", jobname, errscanner.Text())
		}
	}
	if err := errscanner.Err(); err != nil {
		log.Fatalf("reading standard input: %v", err)
	}
	outscanner := bufio.NewScanner(stdout)
	for outscanner.Scan() {
		if len(outscanner.Text()) != 0 {
			log.Infof("<%s> %s", jobname, outscanner.Text())
		}
	}
	if err := outscanner.Err(); err != nil {
		log.Fatalf("reading standard input: %v", err)
	}

	if err := cmd.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			// The program has exited with an exit code != 0

			// This works on both Unix and Windows. Although package
			// syscall is generally platform dependent, WaitStatus is
			// defined for both Unix and Windows and in both cases has
			// an ExitStatus() method with the same signature.
			status, ok := exiterr.Sys().(syscall.WaitStatus)

			if !ok {
				return 0, fmt.Errorf("get exit status fail")
			}

			return status.ExitStatus(), nil
		} else {
			log.Fatalf("cmd.Wait: %v", err)
			return 0, err
		}
	}
	return 0, err
}
