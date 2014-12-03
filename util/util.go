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
	"../log"
	"os"
	"strings"
)

//===================================================================
// Public APIs
//===================================================================

func DebugLines(s string, fn func(v ...interface{})) {
	for _, line := range strings.Split(s, "\n") {
		if len(line) == 0 {
			line = " "
		}
		fn(line)
	}
}

func IsInMap(keys []string, m map[string]string) bool {
	for _, k := range keys {
		if _, ok := m[k]; !ok {
			return false
		}
	}
	return true
}

func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	log.Fatal("check file exist failed: %v", err)
	return false
}

//===================================================================
// Private
//===================================================================
