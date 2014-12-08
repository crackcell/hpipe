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
 * @file errors.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Mon Nov 17 17:00:12 2014
 *
 **/

package eval

import (
	"errors"
)

//===================================================================
// Public APIs
//===================================================================

var (
	ErrFmtUndefinedVar    = "undefined var: %s"
	ErrFmtUnknownOperator = "unknown operator: %s"
	ErrInvalidOperation   = errors.New("invalid operation")
)

//===================================================================
// Private
//===================================================================
