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
	"fmt"
)

//===================================================================
// Public APIs
//===================================================================

type Meta struct {
	db DB
}

type DB interface {
	Set(key, value string) error
	Get(key string) string
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
