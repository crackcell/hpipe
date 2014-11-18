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

package flow

import (
	_ "fmt"
)

//===================================================================
// Public APIs
//===================================================================

type Parser interface {
	ParseStepFromFile(entry string, workdir string) *Step
	ParseJobFromFile(entry string, workdir string) Job
}

//===================================================================
// Private
//===================================================================
