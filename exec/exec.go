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
 * @file exec.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Tue Aug 25 18:09:11 2015
 *
 **/

package exec

import (
	//"fmt"
	"github.com/crackcell/hpipe/dag"
)

//===================================================================
// Public APIs
//===================================================================

type Exec interface {
	Run(d *dag.DAG) error
}

//===================================================================
// Private
//===================================================================
