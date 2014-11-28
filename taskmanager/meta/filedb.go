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
 * @file filedb.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Fri Nov 28 18:14:08 2014
 *
 **/

package meta

import (
	_ "fmt"
)

//===================================================================
// Public APIs
//===================================================================

func NewFileDB() (DB, error) {
	ret := new(FileDB)
	return ret, nil
}

type FileDB struct {
	file string
}

func (this *FileDB) Set(key, value string) error {
	return nil
}

func (this *FileDB) Get(key string) string {
	return ""
}

//===================================================================
// Private
//===================================================================
