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
 * @file hdfs.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sun Sep  6 23:25:23 2015
 *
 **/

package storage

import (
	"fmt"
	"github.com/colinmarc/hdfs"
	"github.com/crackcell/hpipe/config"
	"github.com/crackcell/hpipe/log"
	"os"
)

//===================================================================
// Public APIs
//===================================================================

type HDFS struct {
	client *hdfs.Client
}

func NewHDFS(namenode string) (Storage, error) {
	if client, err := hdfs.New(config.NameNode); err != nil {
		msg := fmt.Sprintf("connect to hdfs namenode failed: %s", config.NameNode)
		log.Fatal(msg)
		return nil, fmt.Errorf(msg)
	} else {
		return &HDFS{client: client}, nil
	}
}

func (this *HDFS) MkdirP(path string) error {
	log.Debugf("mkdir -p %s", path)
	return this.client.MkdirAll(path, 0755)
}

func (this *HDFS) Rm(path string) error {
	log.Debugf("remove file: %s", path)
	return this.client.Remove(path)
}

func (this *HDFS) Touch(path string) error {
	log.Debugf("touch file: %s", path)
	return this.client.CreateEmptyFile(path)
}

func (this *HDFS) IsExist(path string) (bool, error) {
	_, err := this.client.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			log.Debugf("path not exist: %s", path)
			return false, nil
		}
		return false, err
	}
	log.Debugf("path exist: %s", path)
	return true, nil
}

//===================================================================
// Private
//===================================================================
