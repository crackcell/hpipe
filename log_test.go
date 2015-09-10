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
 * @file log_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Thu Nov 20 15:11:58 2014
 *
 **/

package main

import (
	"github.com/crackcell/hpipe/log"
	"os"
	"testing"
)

func TestAll(t *testing.T) {
	l := log.New(os.Stdout, "test_logger ", log.LOG_LEVEL_ALL)
	l.Debug("debug")
	l.Debugf("%s", "debug")
	l.Trace("trace")
	l.Tracef("%s", "trace")
	l.Info("info")
	l.Infof("%s", "info")
	l.Warn("warn")
	l.Warnf("%s", "warn")
	l.Error("error")
	l.Errorf("%s", "error")
	l.Fatal("fatal")
	l.Fatalf("%s", "fatal")
	log.Debug("debug")
	log.Debugf("%s", "debug")
	log.Trace("trace")
	log.Tracef("%s", "trace")
	log.Info("info")
	log.Infof("%s", "info")
	log.Warn("warn")
	log.Warnf("%s", "warn")
	log.Error("error")
	log.Errorf("%s", "error")
	log.Fatal("fatal")
	log.Fatalf("%s", "fatal")
}
