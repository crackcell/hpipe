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
	"../config"
	"../log"
	"../taskmanager"
	"../util"
	"../yafl/ast"
	"../yafl/parser"
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
	config.InitFlags()
	config.Parse()
	if config.FlagHelp {
		showHelp()
	}
	if len(config.FlagEntryFile) == 0 || len(config.FlagWorkRoot) == 0 {
		showHelp()
	}
	f, err := loadFlowFromFile(config.FlagEntryFile, config.FlagWorkRoot)
	if err != nil {
		panic(err)
	}
	if config.FlagVerbose {
		util.LogLines(LogoString, log.Debug)
		util.LogLines(f.DebugString(), log.Debug)
	}
	taskmgr := taskmanager.NewTaskManager()
	taskmgr.Setup(f)
	taskmgr.Run()
}
