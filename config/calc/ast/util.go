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
 * @file util.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Mon Nov 17 15:17:37 2014
 *
 **/

package ast

import (
	"fmt"
	"regexp"
	"time"
)

//===================================================================
// Public APIs
//===================================================================

var timeRegexp = regexp.MustCompile(`^[0-9]{4}(((0[13578]|(10|12))(0[1-9]|[1-2][0-9]|3[0-1]))|(02-(0[1-9]|[1-2][0-9]))|((0[469]|11)(0[1-9]|[1-2][0-9]|30)))$`)

func IsTime(s string) bool {
	if len(s) != 8 {
		return false
	}
	return len(timeRegexp.FindString(s)) != 0
}

const timeFmt = "20060102"

func StrToTime(s string) time.Time {
	if t, err := time.Parse(timeFmt, s); err != nil {
		panic(fmt.Sprintf("invalid time: %s", s))
	} else {
		return t
	}
}

//===================================================================
// Private
//===================================================================
