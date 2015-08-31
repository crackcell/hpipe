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
 * @date Mon Aug 17 21:40:18 2015
 *
 **/

package dag

import (
	"errors"
)

//===================================================================
// Public APIs
//===================================================================

var (
	InvalidNodeType = errors.New("invalid node type")
)

//===================================================================
// Private
//===================================================================
