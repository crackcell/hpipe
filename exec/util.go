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
 * @file util.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Thu Aug 27 19:53:11 2015
 *
 **/

package exec

import (
	"github.com/crackcell/hpipe/dag"
)

//===================================================================
// Public APIs
//===================================================================

//===================================================================
// Private
//===================================================================

func isInMap(key string, m dag.Attrs) bool {
	_, ok := m[key]
	return ok
}

func checkJobAttr(job *dag.Job, keys []string) bool {
	for _, key := range keys {
		if !isInMap(key, job.Attrs) {
			return false
		}
	}
	return true
}
