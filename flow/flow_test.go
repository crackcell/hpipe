/***************************************************************
 *
 * Copyright (c) 2014, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 * Test for Flow
 *
 * @file flow_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Mon Nov 10 19:50:59 2014
 *
 **/

package flow

import (
	"fmt"
	"testing"
)

var step *Step

func TestParseXML(t *testing.T) {
	p := NewXMLParser()
	step = p.ParseStepFromFile("step1.xml", ".")
	fmt.Println(step.DebugString())
}
