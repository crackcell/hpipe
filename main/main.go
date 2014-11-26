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
	"../config/ast"
	"../config/cmdline"
	"../config/parser"
	log "../levellog"
	"../util"
	"flag"
	"fmt"
	"os"
)

func Init() {
	flag.BoolVar(&cmdline.FlagHelp, "help", false, "Print help message")
	flag.BoolVar(&cmdline.FlagHelp, "h", false, "Print help message")
	flag.BoolVar(&cmdline.FlagVerbose, "verbose", false, "Use verbose output")
	flag.BoolVar(&cmdline.FlagVerbose, "v", false, "Use verbose output")
	flag.StringVar(&cmdline.FlagWorkRoot, "workdir", "./", "Root path of the flow")
	flag.StringVar(&cmdline.FlagWorkRoot, "w", "./", "Work root of the flow")
	flag.StringVar(&cmdline.FlagEntryFile, "flow", "", "Entry of the flow")
	flag.StringVar(&cmdline.FlagEntryFile, "f", "", "Entry of the flow")
}

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
	Init()
	flag.Parse()
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
