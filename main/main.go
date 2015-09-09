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
	//"fmt"
	"github.com/crackcell/hpipe/config"
	"github.com/crackcell/hpipe/dag"
	"github.com/crackcell/hpipe/log"
	"github.com/crackcell/hpipe/sched"
	"github.com/crackcell/hpipe/util"
	"os"
)

func main() {
	config.InitFlags()
	config.Parse()
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

	util.LogLines(config.LogoString(), nil)
	util.LogLines(d.String(), nil)

	s, err := sched.NewSched()
	if err != nil {
		os.Exit(1)
	}
	if err := s.Run(d); err != nil {
		os.Exit(1)
	}
}
