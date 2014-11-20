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
	"../config/flow"
	log "../levellog"
	"../util"
	"flag"
	"fmt"
	"os"
)

var (
	flVersion   bool
	flDebug     bool
	flWorkDir   string
	flFlowEntry string
)

func Init() {
	flag.BoolVar(&flVersion, "version", false, "Print version info and quit")
	flag.BoolVar(&flVersion, "v", false, "Print version info and quit")
	flag.BoolVar(&flDebug, "debug", false, "Enable debug mode")
	flag.BoolVar(&flDebug, "d", false, "Enable debug mode")
	flag.StringVar(&flWorkDir, "workdir", "./", "Work root of the flow")
	flag.StringVar(&flWorkDir, "w", "./", "Work root of the flow")
	flag.StringVar(&flFlowEntry, "flow", "", "Entry of the flow")
	flag.StringVar(&flFlowEntry, "f", "", "Entry of the flow")
}

func logo() string {
	return ` 
 _______         __                         ______
|   |   |.-----.|__|.-----.-----.   .--.--.|__    |
|       ||  _  ||  ||  _  |  -__|   |  |  ||    __|
|___|___||   __||__||   __|_____|    \___/ |______|
         |__|       |__|
 `
}

func main() {
	Init()
	flag.Parse()

	if flVersion {
		fmt.Printf("Hpipe v2\n")
		os.Exit(0)
	}

	if len(flFlowEntry) == 0 || len(flWorkDir) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	f := flow.NewFlow()
	err := f.LoadFromFile(flFlowEntry, flWorkDir)
	if err != nil {
		panic(err)
	}

	if flDebug {
		util.LogLines(logo(), log.Debug)
		util.LogLines(f.DebugString(), log.Debug)
	}
}
