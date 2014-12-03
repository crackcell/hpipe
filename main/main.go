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
    hpipe-run [options]

Options:
    -h, --help     Print this message
    -w, --work     Root path of workflow
    -m, --meta     Path for meta data
    -f, --flow     Entry filename of workflow
    -v, --verbose  Use verbose output
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
	if len(config.EntryFile) == 0 || len(config.WorkPath) == 0 ||
		len(config.MetaPath) == 0 || len(config.NodeName) == 0 {
		showHelp()
	}
	fmt.Println(len(config.MetaPath), config.MetaPath)

	p := parser.NewXMLParser()
	f, err := p.ParseFile(config.EntryFile, config.WorkPath)
	if err != nil {
		panic(err)
	}

	if config.Verbose {
		util.DebugLines(LogoString, log.Debug)
		util.DebugLines(f.DebugString(), log.Debug)
	}

	taskmgr, err := task.NewTaskManager()
	if err != nil {
		panic(err)
	}
	taskmgr.Setup(f)
	taskmgr.Run()
}
