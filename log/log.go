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

package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
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
	LOG_LEVEL_NUM      = 6
	LOG_LEVEL_NONDEBUG = LOG_LEVEL_TRACE | LOG_LEVEL_INFO |
		LOG_LEVEL_WARN | LOG_LEVEL_ERROR | LOG_LEVEL_FATAL
	LOG_LEVEL_ALL = LOG_LEVEL_DEBUG | LOG_LEVEL_TRACE | LOG_LEVEL_INFO |
		LOG_LEVEL_WARN | LOG_LEVEL_ERROR | LOG_LEVEL_FATAL
)

type LevelLogger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Trace(v ...interface{})
	Tracef(format string, v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
}

func New(writer io.Writer, prefix string, logLevel int) LevelLogger {
	ret := new(levelLogger)
	ret.prefix = prefix
	ret.callerpath = 3
	ret.init(writer, ret.prefix, logLevel)
	return ret
}

func NewDefault(writer io.Writer, prefix string, logLevel int) LevelLogger {
	ret := new(levelLogger)
	ret.prefix = prefix
	ret.callerpath = 4
	ret.init(writer, prefix, logLevel)
	return ret
}

//===================================================================
// Private
//===================================================================

type underlayLogger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
}

type nullLogger struct{}

func (this *nullLogger) Print(v ...interface{})                 {}
func (this *nullLogger) Printf(format string, v ...interface{}) {}
func (this *nullLogger) Fatal(v ...interface{})                 {}
func (this *nullLogger) Fatalf(format string, v ...interface{}) {}

type levelLogger struct {
	mu         sync.Mutex
	prefix     string
	loggers    map[int]underlayLogger
	callerpath int
}

var levelToStr = map[int]string{
	LOG_LEVEL_DEBUG: "DEBUG",
	LOG_LEVEL_TRACE: "TRACE",
	LOG_LEVEL_INFO:  "INFO",
	LOG_LEVEL_WARN:  "WARN",
	LOG_LEVEL_ERROR: "ERROR",
	LOG_LEVEL_FATAL: "FATAL",
}

var StdLogger = NewDefault(os.Stdout, "", LOG_LEVEL_ALL)

func (this *levelLogger) init(writer io.Writer, prefix string, logLevel int) {
	this.loggers = make(map[int]underlayLogger)
	null := new(nullLogger)
	var i uint
	for i = 1; i <= LOG_LEVEL_NUM; i++ {
		level := 1 << i
		if level&logLevel != 0 {
			levelName, _ := levelToStr[level]
			this.loggers[level] =
				log.New(writer, levelName+" "+prefix+" ", log.LstdFlags)
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

func (this *levelLogger) print(level int, verbose bool, v ...interface{}) {
	l := this.getLogger(level)
	this.mu.Lock()
	defer this.mu.Unlock()
	if verbose {
		v = append([]interface{}{getCallerMoreInfo(this.callerpath) + ": "}, v...)
	} else {
		v = append([]interface{}{getCallerInfo(this.callerpath) + ": "}, v...)
	}
	l.Print(v...)
}

func (this *levelLogger) printf(level int, verbose bool, format string, v ...interface{}) {
	l := this.getLogger(level)
	this.mu.Lock()
	defer this.mu.Unlock()
	if verbose {
		format = getCallerMoreInfo(this.callerpath) + ": " + format
	} else {
		format = getCallerInfo(this.callerpath) + ": " + format
	}
	l.Printf(format, v...)
}

func Debug(v ...interface{}) {
	StdLogger.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	StdLogger.Debugf(format, v...)
}

func Trace(v ...interface{}) {
	StdLogger.Trace(v...)
}

func Tracef(format string, v ...interface{}) {
	StdLogger.Tracef(format, v...)
}

func Info(v ...interface{}) {
	StdLogger.Info(v...)
}

func Infof(format string, v ...interface{}) {
	StdLogger.Infof(format, v...)
}

func Warn(v ...interface{}) {
	StdLogger.Warn(v...)
}

func Warnf(format string, v ...interface{}) {
	StdLogger.Warnf(format, v...)
}

func WarnErrf(format string, v ...interface{}) error {
	err := fmt.Errorf(format, v)
	StdLogger.Warn(err)
	return err
}

func Error(v ...interface{}) {
	StdLogger.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	StdLogger.Errorf(format, v...)
}

func ErrorErrf(format string, v ...interface{}) error {
	err := fmt.Errorf(format, v)
	StdLogger.Error(err)
	return err
}

func Fatal(v ...interface{}) {
	StdLogger.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	StdLogger.Fatalf(format, v...)
}

func FatalErrf(format string, v ...interface{}) error {
	err := fmt.Errorf(format, v)
	StdLogger.Fatal(err)
	return err
}

func (this *levelLogger) Debug(v ...interface{}) {
	this.print(LOG_LEVEL_DEBUG, true, v...)
}

func (this *levelLogger) Debugf(format string, v ...interface{}) {
	this.printf(LOG_LEVEL_DEBUG, true, format, v...)
}

func (this *levelLogger) Trace(v ...interface{}) {
	this.print(LOG_LEVEL_TRACE, true, v...)
}

func (this *levelLogger) Tracef(format string, v ...interface{}) {
	this.printf(LOG_LEVEL_TRACE, true, format, v...)
}

func (this *levelLogger) Info(v ...interface{}) {
	this.print(LOG_LEVEL_INFO, false, v...)
}

func (this *levelLogger) Infof(format string, v ...interface{}) {
	this.printf(LOG_LEVEL_INFO, false, format, v...)
}

func (this *levelLogger) Warn(v ...interface{}) {
	this.print(LOG_LEVEL_WARN, true, v...)
}

func (this *levelLogger) Warnf(format string, v ...interface{}) {
	this.printf(LOG_LEVEL_WARN, true, format, v...)
}

func (this *levelLogger) Error(v ...interface{}) {
	this.print(LOG_LEVEL_ERROR, true, v...)
}

func (this *levelLogger) Errorf(format string, v ...interface{}) {
	this.printf(LOG_LEVEL_ERROR, true, format, v...)
}

func (this *levelLogger) Fatal(v ...interface{}) {
	this.print(LOG_LEVEL_FATAL, true, v...)
}

func (this *levelLogger) Fatalf(format string, v ...interface{}) {
	this.printf(LOG_LEVEL_FATAL, true, format, v...)
}

func getCallerInfo(callpath int) string {
	pc, _, _, _ := runtime.Caller(callpath)
	return shorten(runtime.FuncForPC(pc).Name())
}

func getCallerMoreInfo(callpath int) string {
	pc, file, line, ok := runtime.Caller(callpath)
	if !ok {
		file = "???"
		line = 0
	}

	file = shorten(file)

	return file + ":" + strconv.Itoa(line) + " " +
		shorten(runtime.FuncForPC(pc).Name())
}

func shorten(long string) string {
	short := long
	for i := len(long) - 1; i > 0; i-- {
		if long[i] == '/' {
			short = long[i+1:]
			break
		}
	}
	return short
}
