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
 * @file date.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Fri Aug 20 20:20:14 2015
 *
 **/

package time

import (
	"strings"
	stdtime "time"
)

//===================================================================
// Public APIs
//===================================================================

var dateFormats = []string{
	// year
	"YYYY", "2006",
	"YY", "06",

	// month
	"MM", "01",
	"M", "1",

	// day
	"DD", "02",
	"D", "2",

	// hour
	"hh", "03",
	"h", "3",
	"mm", "04",
	"m", "4",
	"ss", "05",
	"s", "5",
}

// Regular time format to golang format
func DateToGoFormat(format string) string {
	replacer := strings.NewReplacer(dateFormats...)
	return replacer.Replace(format)
}

// String to time.Time
func Parse(value string, format string) (stdtime.Time, error) {
	return stdtime.Parse(DateToGoFormat(format), value)
}

func Format(t stdtime.Time, format string) string {
	return t.Format(DateToGoFormat(format))
}

//===================================================================
// Private
//===================================================================
