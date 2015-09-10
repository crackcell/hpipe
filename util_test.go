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
 * @file time_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Mon Aug 23 14:00:33 2015
 *
 **/

package main

import (
	"fmt"
	"github.com/crackcell/hpipe/util/time"
	"testing"
	stdtime "time"
)

func TestDateToGoFormat(t *testing.T) {
	fmt.Println(time.DateToGoFormat("YYYYMMDD hh:mm:ss"))
	fmt.Println(time.DateToGoFormat("YYYY-MM-DD hh:mm:ss"))
	fmt.Println(time.DateToGoFormat("YY/M/D hh:mm:ss"))
}

func TestParse(t *testing.T) {
	fmt.Println(time.Parse("20060102 03:04:05", "YYYYMMDD hh:mm:ss"))
}

func TestFormat(t *testing.T) {
	fmt.Println(time.Format(stdtime.Now(), "YYYYMMDD hh:mm:ss"))
}
