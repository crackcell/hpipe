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

package config

import (
	"flag"
	"fmt"
	"os"
)

//===================================================================
// Public APIs
//===================================================================

var (
	Help               bool
	Verbose            bool
	WorkPath           string
	EntryFile          string
	MaxRetry           int
	StatusKeeper       string
	NameNode           string
	SqliteFile         string
	Hadoop             bool
	HadoopStreamingJar string
	Odps               bool
	OdpsEndpoint       string
	OdpsProject        string
	OdpsAccessID       string
	OdpsAccessKey      string
	Hive               bool
	GMTDate            string
	LessLog            bool
)

func InitFlags() {
	flag.BoolVar(&Help, "help", false, "Print help message")
	flag.BoolVar(&Help, "h", false, "Print help message")
	flag.BoolVar(&Verbose, "verbose", false, "Use verbose output")
	flag.BoolVar(&Verbose, "v", false, "Use verbose output")
	flag.StringVar(&WorkPath, "path", "./", "Working path")
	flag.StringVar(&WorkPath, "p", "./", "Working path")
	flag.StringVar(&EntryFile, "flow", "", "Entry of the flow")
	flag.StringVar(&EntryFile, "f", "", "Entry of the flow")
	flag.IntVar(&MaxRetry, "max-retry", 3, "max retry times of failed jobs, default: 3")
	flag.StringVar(&StatusKeeper, "status-keeper", "sqlite", "method to track job status, default: sqlite, available: hdfs, sqlite, file")
	flag.StringVar(&NameNode, "namenode", "", "Hadoop name node url")
	flag.StringVar(&SqliteFile, "sqlite", "", "Sqlite file")
	flag.BoolVar(&Hadoop, "hadoop", false, "enable hadoop streaming job")
	flag.StringVar(&HadoopStreamingJar, "jar", "", "Hadoop streaming jar")
	flag.BoolVar(&Odps, "odps", false, "enable ODPS job")
	flag.StringVar(&OdpsEndpoint, "odps-endpoint", "", "ODPS endpoint address")
	flag.StringVar(&OdpsProject, "odps-project", "", "ODPS project name")
	flag.StringVar(&OdpsAccessID, "odps-access-id", "", "ODPS access id")
	flag.StringVar(&OdpsAccessKey, "odps-access-key", "", "ODPS access key")
	flag.BoolVar(&Hive, "hive", false, "Enable Hive job")
	flag.StringVar(&GMTDate, "gmtdate", "", "Set variable $gmtdate in YYYYMMDD format")
	flag.BoolVar(&LessLog, "less-log", false, "Less log output")
}

func Parse() {
	flag.Parse()
	if Help {
		showHelp(0)
	}
	if len(EntryFile) == 0 {
		fmt.Println("invalid arguments: no flow")
		showHelp(1)
	}
	if Hadoop && len(NameNode) == 0 {
		fmt.Println("invalid arguments: no namenode")
		showHelp(1)
	}
	if StatusKeeper == "hdfs" && len(NameNode) == 0 {
		fmt.Println("invalid arguments: no namenode")
		showHelp(1)
	} else if StatusKeeper == "sqlite" && len(SqliteFile) == 0 {
		fmt.Println("invalid arguments: no sqlite")
		showHelp(1)
	}
}

//===================================================================
// Private
//===================================================================

const (
	logoString = ` _______         __
|   |   |.-----.|__|.-----.-----.
|       ||  _  ||  ||  _  |  -__|
|___|___||   __||__||   __|_____|
         |__|       |__|
`
	helpString = `Execute a hpipe workflow
Usage:
    hpipe [options]
Options:
    -h, --help         Print this message
    -v, --verbose      Use verbose output

    -p, --path         Working path
    -f, --flow         Entry filename of workflow
    --max-retry        Max retry times of failed jobs, default: 3

    --status-keeper    Method to track job status
                       default: sqlite
                       available: hdfs, sqlite

    --namenode         Address of Hadoop NameNode, default: 127.0.0.1:8020
    --sqlite           File path for sqlite database

    --hadoop           Enable Hadoop streaming job
    --jar              Path of Hadoop streaming jar file

    --odps             Enable ODPS job
    --odps-endpoint    Address of ODPS endpoing
    --odps-project     ODPS project name
    --odps-access-id   ODPS access id
    --odps-access-key  ODPS access key

    --hive             Enable Hive job

    --gmtdate          Set variable $gmtdate in YYYYMMDD format

    --less-log         Less log output
`
)

func showHelp(ret int) {
	fmt.Print(helpString)
	os.Exit(ret)
}

func LogoString() string {
	return logoString
}
