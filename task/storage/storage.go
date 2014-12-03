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
 * @file meta.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Fri Nov 28 18:05:38 2014
 *
 **/

package storage

import (
	"../../yafl/ast"
)

//===================================================================
// Public APIs
//===================================================================

type DB interface {
	Open() error
	Close() error
	SaveFlow(f *ast.Flow) error
	RestoreFlow(f *ast.Flow) (*ast.Flow, error)
}

//===================================================================
// Private
//===================================================================
