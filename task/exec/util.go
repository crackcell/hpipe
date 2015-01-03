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
 * @file util.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Fri Dec  5 12:42:28 2014
 *
 **/

package exec

import (
	"../../etc"
	"../../log"
	"bufio"
	"fmt"
	"os/exec"
	"strings"
	"syscall"
)

//===================================================================
// Public APIs
//===================================================================

//===================================================================
// Private
//===================================================================

func CmdExec(jobname, name string, arg ...string) (int, error) {
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
			//log.Fatalf("<%s> %s", jobname, errscanner.Text())
			log.Infof("<%s> %s", jobname, errscanner.Text())
			//fmt.Printf("<%s> %s", jobname, errscanner.Text())
		}
	}
	if err := errscanner.Err(); err != nil {
		log.Fatalf("reading standard input: %v", err)
	}
	outscanner := bufio.NewScanner(stdout)
	for outscanner.Scan() {
		if len(outscanner.Text()) != 0 {
			log.Infof("<%s> %s", jobname, outscanner.Text())
			//fmt.Printf("<%s> %s", jobname, outscanner.Text())
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

func GetProp(m map[string]string, key string) string {
	v, ok := m[key]
	if ok {
		return v
	}
	v, ok = etc.Env[key]
	if !ok {
		panic(fmt.Errorf("no prop for %s", key))
	}
	return v
}

func ExistProp(m map[string]string, key string) bool {
	_, ok1 := m[key]
	_, ok2 := etc.Env[key]
	return ok1 || ok2
}

func LogArgList(cmd string, args ...string) {
	var argstr string
	for _, arg := range args {
		argstr += " " + arg
	}
	log.Debugf("command: %s %s", cmd, argstr)
}

func PrepareArgList(prop map[string]string, args [][]string) []string {
	var str []string
	for _, a := range args {
		if len(a) != 3 {
			panic(fmt.Errorf("not valid args"))
		}
		str = append(str, PrepareArg(prop, a[0], a[1], a[2] == "s")...)
	}
	return str
}

func PrepareArg(prop map[string]string, propname, argname string,
	split bool) []string {

	p := []string{}
	var t []string
	if split {
		t = strings.Split(GetProp(prop, propname), ",")
	} else {
		t = append(t, GetProp(prop, propname))
	}

	for _, v := range t {
		tmp := strings.Trim(v, " ")
		if len(tmp) == 0 {
			continue
		}
		if strings.Contains(tmp, " ") {
			tmp = `"` + tmp + `"`
		}
		p = append(p, tmp)
	}
	a := []string{}
	for _, v := range p {
		a = append(a, argname, v)
	}
	return a
}

func ValidProp(prop map[string]string, names []string) bool {
	for _, p := range names {
		if !ExistProp(prop, p) {
			log.Fatalf("%s not found or empty", p)
			return false
		}
	}
	return true
}
