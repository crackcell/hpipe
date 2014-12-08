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
	"../task"
	"../util"
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
    hpipe [options]

Options:
    -h, --help     Print this message
    -v, --verbose  Use verbose output

    -w, --work     Root path of workflow
    -m, --meta     Path of meta data
    -f, --flow     Entry filename of workflow

    --rerun        Rerun jobs with status DONE
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

	p := parser.NewXMLParser()
	f, err := p.ParseFile(config.EntryFile, config.WorkPath)
	if err != nil {
		panic(err)
	}

	if config.Verbose {
		util.LogLines(LogoString, log.Debug)
		util.LogLines(f.DebugString(), log.Debug)
	}

	taskmgr, err := task.NewTaskManager()
	if err != nil {
		panic(err)
	}
	taskmgr.Run(f)
}
