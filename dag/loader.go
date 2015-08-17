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
 * @file loader.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Mon Aug 17 21:56:55 2015
 *
 **/

package dag

import (
//"fmt"
)

//===================================================================
// Public APIs
//===================================================================

type Loader interface {
	LoadFile(path string) (*DAG, error)
	LoadBytes(data []byte) (*DAG, error)
}

//===================================================================
// Private
//===================================================================
