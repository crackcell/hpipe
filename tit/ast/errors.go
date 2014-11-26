/***************************************************************
 *
 * Copyright (c) 2014, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 * Error
 *
 * @file errors.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Mon Nov 17 13:14:49 2014
 *
 **/

package ast

import (
	"errors"
)

var (
	ErrUnknownOperator = errors.New("unknown operator")
)
