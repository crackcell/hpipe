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
 * @date Tue Nov 18 18:10:11 2014
 *
 **/

package ast

import (
	"errors"
)

var (
	ErrInvalidJob = errors.New("invalid job definition")
)
