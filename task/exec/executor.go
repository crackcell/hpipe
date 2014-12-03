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
 * @file executor.go
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

type Executor interface {
	Setup(prop map[string]string) error
	Run(job ast.Job) (string, error) // ret_status, error
}

//===================================================================
// Private
//===================================================================
