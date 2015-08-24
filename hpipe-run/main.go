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
 * @file main.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Tue Aug 25 00:16:25 2015
 *
 **/

package main

import (
	"fmt"
	"github.com/crackcell/hpipe/config"
	"github.com/crackcell/hpipe/dag"
	//"github.com/crackcell/hpipe/log"
	"github.com/crackcell/hpipe/util"
	"os"
)

const (
	LogoString = ` _______         __
|   |   |.-----.|__|.-----.-----.
|       ||  _  ||  ||  _  |  -__|
|___|___||   __||__||   __|_____|
         |__|       |__|
`
	HelpString = `Execute a hpipe workflow
Usage:
    hpipe [options]
Options:
    -h, --help     Print this message
    -v, --verbose  Use verbose output
    -f, --flow     Entry filename of workflow
`
)

func showHelp() {
	fmt.Print(HelpString)
	os.Exit(0)
}

func main() {
	config.InitFlags()
	config.Parse()
	if config.Help {
		showHelp()
	}
	if len(config.EntryFile) == 0 {
		showHelp()
	}

	factory := dag.NewFactory()
	flow, err := factory.CreateDAGFromFile(config.EntryFile)
	if err != nil {
		panic(err)
	}

	util.LogLines(LogoString, nil)
	util.LogLines(flow.String(), nil)
}
