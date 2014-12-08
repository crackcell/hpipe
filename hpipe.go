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
 * @file hpipe.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sat Dec  6 21:42:12 2014
 *
 **/

package hpipe

import (
	_ "fmt"
)

//===================================================================
// Public APIs
//===================================================================

const (
	TODO       = "todo"
	RUNNING    = "run"
	DONE       = "done"
	FAIL       = "fail"
	RETRY_FAIL = "retry_fail" // failed after N times retry
)

//===================================================================
// Private
//===================================================================
