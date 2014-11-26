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
 * @file cmdline.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sun Nov 23 20:59:30 2014
 *
 **/

package cmdline

import (
	"flag"
	_ "fmt"
)

//===================================================================
// Public APIs
//===================================================================

var (
	FlagHelp      bool
	FlagVerbose   bool
	FlagWorkRoot  string
	FlagEntryFile string
)

func InitFlags() {
	flag.BoolVar(&FlagHelp, "help", false, "Print help message")
	flag.BoolVar(&FlagHelp, "h", false, "Print help message")
	flag.BoolVar(&FlagVerbose, "verbose", false, "Use verbose output")
	flag.BoolVar(&FlagVerbose, "v", false, "Use verbose output")
	flag.StringVar(&FlagWorkRoot, "workdir", "./", "Root path of the flow")
	flag.StringVar(&FlagWorkRoot, "w", "./", "Work root of the flow")
	flag.StringVar(&FlagEntryFile, "flow", "", "Entry of the flow")
	flag.StringVar(&FlagEntryFile, "f", "", "Entry of the flow")
}

func Parse() {
	flag.Parse()
}

//===================================================================
// Private
//===================================================================
