/***************************************************************
 *
 * Copyright (c) 2014, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 * Main
 *
 * @file main.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Tue Nov 18 17:18:03 2014
 *
 **/

package main

import (
	"../ast"
	"../config/cmdline"
	"../config/parser"
	"../log"
	"../util"
	"fmt"
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
    hpipe-run [options]

Options:
    -h, --help     Print this message
    -w, --worddir  Root path of workflow
    -f, --flow     Entry filename of workflow
    -v, --verbose  Use verbose output
`
)

func showHelp() {
	fmt.Print(HelpString)
	os.Exit(0)
}

func loadFlowFromFile(filename, workdir string) (*ast.Flow, error) {
	f := ast.NewFlow()
	p := parser.NewXMLParser()
	if step, err := p.ParseStepFromFile(filename, workdir); err != nil {
		return nil, err
	} else {
		f.Entry = step
		return f, nil
	}
}

func main() {
	cmdline.InitFlags()
	cmdline.Parse()
	if cmdline.FlagHelp {
		showHelp()
	}
	if len(cmdline.FlagEntryFile) == 0 || len(cmdline.FlagWorkRoot) == 0 {
		showHelp()
	}
	f, err := loadFlowFromFile(cmdline.FlagEntryFile, cmdline.FlagWorkRoot)
	if err != nil {
		panic(err)
	}
	if cmdline.FlagVerbose {
		util.LogLines(LogoString, log.Debug)
		util.LogLines(f.DebugString(), log.Debug)
	}
}
