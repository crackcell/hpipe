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

package log

import (
	"os"
	"testing"
)

func TestAll(t *testing.T) {
	l := New(os.Stdout, "test_logger ", LOG_LEVEL_ALL)
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
	Debug("debug")
	Debugf("%s", "debug")
	Trace("trace")
	Tracef("%s", "trace")
	Info("info")
	Infof("%s", "info")
	Warn("warn")
	Warnf("%s", "warn")
	Error("error")
	Errorf("%s", "error")
	Fatal("fatal")
	Fatalf("%s", "fatal")
}
