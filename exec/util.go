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
	"fmt"
	"github.com/crackcell/hpipe/dag"
)

//===================================================================
// Public APIs
//===================================================================

//===================================================================
// Private
//===================================================================

func getProp(key string, m dag.Attrs) string {
	v, ok := m[key]
	if !ok {
		panic(fmt.Errorf("key not found: %s", key))
	}
	return v
}
