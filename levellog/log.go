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
 * @file log.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Thu Nov 20 11:37:37 2014
 *
 **/

package levellog

import (
	"io"
	"log"
	"os"
	"sync"
)

//===================================================================
// Public APIs
//===================================================================

const (
	LOG_LEVEL_DEBUG = 1 << (iota + 1)
	LOG_LEVEL_TRACE
	LOG_LEVEL_INFO
	LOG_LEVEL_WARN
	LOG_LEVEL_ERROR
	LOG_LEVEL_FATAL
	LOG_LEVEL_NUM = 6
	LOG_LEVEL_ALL = LOG_LEVEL_DEBUG | LOG_LEVEL_TRACE | LOG_LEVEL_INFO |
		LOG_LEVEL_WARN | LOG_LEVEL_ERROR | LOG_LEVEL_FATAL
)

type LevelLogger interface {
	Debug(v ...interface{})
	Debugf(fmt string, v ...interface{})
	Trace(v ...interface{})
	Tracef(fmt string, v ...interface{})
	Info(v ...interface{})
	Infof(fmt string, v ...interface{})
	Warn(v ...interface{})
	Warnf(fmt string, v ...interface{})
	Error(v ...interface{})
	Errorf(fmt string, v ...interface{})
	Fatal(v ...interface{})
	Fatalf(fmt string, v ...interface{})
}

func New(writer io.Writer, prefix string, logLevel int) LevelLogger {
	ret := new(levelLogger)
	ret.prefix = prefix
	ret.init(writer, prefix, logLevel)
	return ret
}

//===================================================================
// Private
//===================================================================

type underlayLogger interface {
	Print(v ...interface{})
	Printf(fmt string, v ...interface{})
	Fatal(v ...interface{})
	Fatalf(fmt string, v ...interface{})
}

type nullLogger struct{}

func (this *nullLogger) Print(v ...interface{})              {}
func (this *nullLogger) Printf(fmt string, v ...interface{}) {}
func (this *nullLogger) Fatal(v ...interface{})              {}
func (this *nullLogger) Fatalf(fmt string, v ...interface{}) {}

type levelLogger struct {
	mu      sync.Mutex
	prefix  string
	loggers map[int]underlayLogger
}

var levelToStr = map[int]string{
	LOG_LEVEL_DEBUG: "DEBUG",
	LOG_LEVEL_TRACE: "TRACE",
	LOG_LEVEL_INFO:  "INFO",
	LOG_LEVEL_WARN:  "WARN",
	LOG_LEVEL_ERROR: "ERROR",
	LOG_LEVEL_FATAL: "FATAL",
}

var std = New(os.Stdout, "", LOG_LEVEL_ALL)

func (this *levelLogger) init(writer io.Writer, prefix string, logLevel int) {
	this.loggers = make(map[int]underlayLogger)
	null := new(nullLogger)
	var i uint
	for i = 1; i <= LOG_LEVEL_NUM; i++ {
		level := 1 << i
		if level&logLevel != 0 {
			levelName, _ := levelToStr[level]
			this.loggers[level] =
				log.New(writer, levelName+" "+prefix,
					log.LstdFlags|log.Lshortfile)
		} else {
			this.loggers[level] = null
		}
	}
}

func (this *levelLogger) getLogger(level int) underlayLogger {
	l, ok := this.loggers[level]
	if !ok {
		panic("logger 404")
	}
	return l
}

func (this *levelLogger) print(level int, v ...interface{}) {
	l := this.getLogger(level)
	this.mu.Lock()
	defer this.mu.Unlock()
	l.Print(v...)
}

func (this *levelLogger) printf(level int, fmt string, v ...interface{}) {
	l := this.getLogger(level)
	this.mu.Lock()
	defer this.mu.Unlock()
	l.Printf(fmt, v...)
}

func Debug(v ...interface{}) {
	std.Debug(v...)
}

func Debugf(fmt string, v ...interface{}) {
	std.Debugf(fmt, v...)
}

func Trace(v ...interface{}) {
	std.Trace(v...)
}

func Tracef(fmt string, v ...interface{}) {
	std.Tracef(fmt, v...)
}

func Info(v ...interface{}) {
	std.Info(v...)
}

func Infof(fmt string, v ...interface{}) {
	std.Infof(fmt, v...)
}

func Warn(v ...interface{}) {
	std.Warn(v...)
}

func Warnf(fmt string, v ...interface{}) {
	std.Warnf(fmt, v...)
}

func Error(v ...interface{}) {
	std.Error(v...)
}

func Errorf(fmt string, v ...interface{}) {
	std.Errorf(fmt, v...)
}

func Fatal(v ...interface{}) {
	std.Fatal(v...)
}

func Fatalf(fmt string, v ...interface{}) {
	std.Fatalf(fmt, v...)
}

func (this *levelLogger) Debug(v ...interface{}) {
	this.print(LOG_LEVEL_DEBUG, v...)
}

func (this *levelLogger) Debugf(fmt string, v ...interface{}) {
	this.printf(LOG_LEVEL_DEBUG, fmt, v...)
}

func (this *levelLogger) Trace(v ...interface{}) {
	this.print(LOG_LEVEL_TRACE, v...)
}

func (this *levelLogger) Tracef(fmt string, v ...interface{}) {
	this.printf(LOG_LEVEL_TRACE, fmt, v...)
}

func (this *levelLogger) Info(v ...interface{}) {
	this.print(LOG_LEVEL_INFO, v...)
}

func (this *levelLogger) Infof(fmt string, v ...interface{}) {
	this.printf(LOG_LEVEL_INFO, fmt, v...)
}

func (this *levelLogger) Warn(v ...interface{}) {
	this.print(LOG_LEVEL_WARN, v...)
}

func (this *levelLogger) Warnf(fmt string, v ...interface{}) {
	this.printf(LOG_LEVEL_WARN, fmt, v...)
}

func (this *levelLogger) Error(v ...interface{}) {
	this.print(LOG_LEVEL_ERROR, v...)
}

func (this *levelLogger) Errorf(fmt string, v ...interface{}) {
	this.printf(LOG_LEVEL_ERROR, fmt, v...)
}

func (this *levelLogger) Fatal(v ...interface{}) {
	this.print(LOG_LEVEL_FATAL, v...)
}

func (this *levelLogger) Fatalf(fmt string, v ...interface{}) {
	this.printf(LOG_LEVEL_FATAL, fmt, v...)
}
