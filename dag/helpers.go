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
 * @file helpers.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Wed Oct  7 11:13:43 2015
 *
 **/

package dag

import (
//"fmt"
)

//===================================================================
// Public APIs
//===================================================================

func NewHiveJob(name, output string) *Job {
	j := NewJob()
	j.Name = name
	j.Type = HiveJob
	j.Status = NotStarted
	return j
}

func (this *Job) SetHiveHql(hql string) {
	this.Vars["hql"] = hql
}

func (this *Job) SetHiveScript(script string) {
	this.Vars["script"] = script
}

//===================================================================
// Private
//===================================================================
