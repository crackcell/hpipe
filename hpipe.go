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
 * @file hpipe.go
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
	"github.com/crackcell/hpipe/status"
	"github.com/crackcell/hpipe/util"
	"os"
)

func newStatusFlag() (status.Saver, error) {
	switch config.StatusSaver {
	case "hdfs":
		return status.NewHDFSSaver(config.NameNode)
	case "sqlite":
		return status.NewSqliteSaver(config.SqliteFile)
	default:
		msg := fmt.Sprintf("invalid status keeper type: %s",
			config.StatusSaver)
		log.Fatal(msg)
		return nil, fmt.Errorf(msg)
	}
}

func newSched() (*sched.Sched, error) {
	saver, err := newStatusFlag()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return sched.NewSched(status.NewStatusTracker(saver))
}

func main() {
	config.InitFlags()
	config.Parse()

	loglevel := 0
	if config.Verbose {
		loglevel = log.LOG_LEVEL_ALL
	} else {
		loglevel = log.LOG_LEVEL_TRACE | log.LOG_LEVEL_INFO |
			log.LOG_LEVEL_WARN | log.LOG_LEVEL_ERROR |
			log.LOG_LEVEL_FATAL
	}

	if config.LessLog {
		log.StdLogger = log.NewDefaultCleanLogger(
			os.Stdout, "hpipe", loglevel)
	} else {
		log.StdLogger = log.NewDefault(
			os.Stdout, "hpipe", loglevel)
	}

	d, err := dag.LoadFromFile(config.EntryFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	util.LogLines(config.LogoString(), nil)
	util.LogLines(d.String(), nil)

	s, err := newSched()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if err := s.Run(d); err != nil {
		os.Exit(1)
	}
}
