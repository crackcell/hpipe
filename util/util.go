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
 * @file util.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Thu Nov 20 16:58:01 2014
 *
 **/

package util

import (
	"strings"
)

//===================================================================
// Public APIs
//===================================================================

func LogLines(s string, fn func(v ...interface{})) {
	for _, line := range strings.Split(s, "\n") {
		fn(line)
	}
}

//===================================================================
// Private
//===================================================================
