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
 * @file sched_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Wed Aug 26 18:58:32 2015
 *
 **/

package sched

import (
	"github.com/crackcell/hpipe/config"
	"github.com/crackcell/hpipe/dag"
	"testing"
)

func TestDAGExecRun(t *testing.T) {
	config.WorkPath = "../examples/wordcount"
	f := dag.NewFactory()
	d, err := f.CreateDAGFromFile(config.WorkPath + "/wordcount.dot")
	if err != nil {
		t.Error(err)
		return
	}
	config.NameNode = "hpc0:8020"
	config.HadoopStreamingJar = "/usr/lib/hadoop-mapreduce/hadoop-streaming.jar"
	s, err := NewSched()
	if err != nil {
		t.Error(err)
		return
	}
	s.Run(d)
}
