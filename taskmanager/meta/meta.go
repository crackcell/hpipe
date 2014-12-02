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
 * @file meta.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Fri Nov 28 18:05:38 2014
 *
 **/

package meta

import (
	"../../yafl/ast"
	_ "fmt"
)

//===================================================================
// Public APIs
//===================================================================

type Meta struct {
	db DB
}

type DB interface {
	Close() error
	SaveFlow(f *ast.Flow) error
	FetchFlow(f *ast.Flow) error
}

func NewMeta() *Meta {
	return new(Meta)
}

func (this *Meta) SetFlowInfo(info string) error {
	return nil
}

func (this *Meta) GetFlowInfo() (string, error) {
	return "", nil
}

func (this *Meta) SetJobInfo(info string) error {
	return nil
}

func (this *Meta) GetJobInfo() (string, error) {
	return "", nil
}

//===================================================================
// Private
//===================================================================
