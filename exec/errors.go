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
 * @file errors.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Wed Aug 26 18:12:08 2015
 *
 **/

package exec

import (
	"errors"
)

//===================================================================
// Public APIs
//===================================================================

var (
	JobFailedError = errors.New("job failed")
)

//===================================================================
// Private
//===================================================================
