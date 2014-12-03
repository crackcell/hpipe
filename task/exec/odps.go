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
	"../../yafl/ast"
	_ "fmt"
)

//===================================================================
// Public APIs
//===================================================================

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

func (this *ODPSExecutor) Run(job ast.Job) (string, error) {
	return nil
}
