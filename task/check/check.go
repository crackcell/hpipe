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
 * @file check.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sun Dec  7 16:33:05 2014
 *
 **/

package check

import (
	"../../yafl/ast"
)

//===================================================================
// Public APIs
//===================================================================

type Check interface {
	CheckStatus(job *ast.Job) (string, error)
}

//===================================================================
// Private
//===================================================================
