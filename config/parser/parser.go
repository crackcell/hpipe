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
 * @file parser.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Mon Nov 10 15:41:35 2014
 *
 **/

package parser

import (
	"../../ast"
)

//===================================================================
// Public APIs
//===================================================================

type Parser interface {
	ParseStepFromFile(entry string, workdir string) (*ast.Step, error)
	ParseJobFromFile(entry string, workdir string) (ast.Job, error)
}

//===================================================================
// Private
//===================================================================
