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
 * @file attr.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Tue Aug 18 14:04:41 2015
 *
 **/

package dag

import (
//"fmt"
)

//===================================================================
// Public APIs
//===================================================================

//===================================================================
// Private
//===================================================================

type Attrs map[string]string

func NewAttrs() Attrs {
	return make(Attrs)
}

func (this Attrs) Set(key string, value string) {
	this[key] = value
}

func (this Attrs) Extend(more Attrs) {
	for key, value := range more {
		this.Set(key, value)
	}
}

// Only add missing values
func (this Attrs) Ammend(more Attrs) {
	for key, value := range more {
		if _, ok := this[key]; !ok {
			this.Set(key, value)
		}
	}
}
