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
 * @file main.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Tue Aug 25 00:16:25 2015
 *
 **/

package main

import (
	"fmt"
	"github.com/crackcell/hpipe/config"
	"github.com/crackcell/hpipe/dag"
	"github.com/crackcell/hpipe/log"
	"github.com/crackcell/hpipe/sched"
	"github.com/crackcell/hpipe/util"
	"os"
)

const (
	LogoString = ` _______         __
|   |   |.-----.|__|.-----.-----.
|       ||  _  ||  ||  _  |  -__|
|___|___||   __||__||   __|_____|
         |__|       |__|
`
	HelpString = `Execute a hpipe workflow
Usage:
    hpipe [options]
Options:
    -h, --help     Print this message
    -v, --verbose  Use verbose output

    -p, --path     Working path
    -f, --flow     Entry filename of workflow
    --max-retry    Max retry times of failed jobs, default: 3

    --namenode     Address of Hadoop NameNode, default: 127.0.0.1:8020
    --jar          Path of Hadoop streaming jar file
`
)

func showHelp() {
	fmt.Print(HelpString)
	os.Exit(0)
}

func main() {
	config.InitFlags()
	config.Parse()
	if config.Help {
		showHelp()
	}
	if len(config.EntryFile) == 0 {
		showHelp()
	}
	if config.Verbose {
		log.StdLogger = log.NewDefault(os.Stdout, "hpipe", log.LOG_LEVEL_ALL)
	} else {
		log.StdLogger = log.NewDefault(os.Stdout, "hpipe",
			log.LOG_LEVEL_TRACE|log.LOG_LEVEL_INFO|log.LOG_LEVEL_WARN|log.LOG_LEVEL_ERROR|log.LOG_LEVEL_FATAL)
	}

	factory := dag.NewFactory()
	d, err := factory.CreateDAGFromFile(config.EntryFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	util.LogLines(LogoString, nil)
	util.LogLines(d.String(), nil)

	s, err := sched.NewSched()
	if err != nil {
		os.Exit(1)
	}
	if err := s.Run(d); err != nil {
		os.Exit(1)
	}
}
