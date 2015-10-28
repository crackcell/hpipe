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
 * @file convert.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Wed Oct 28 10:59:35 2015
 *
 **/

package util

import (
	"fmt"
	"strings"
)

//===================================================================
// Public APIs
//===================================================================

func StringToBool(str string) (bool, error) {
	switch strings.ToLower(str) {
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		return false, fmt.Errorf("invalid bool value: %s", str)
	}
}

//===================================================================
// Private
//===================================================================
