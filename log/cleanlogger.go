/***************************************************************
 *
 * Copyright (c) 2015, Menglong TAN <tanmenglong@gmail.com>
 *
 **************************************************************/

/**
 *
 *
 * @file cleanlogger.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sat Oct 17 00:23:33 2015
 *
 **/

package log

import (
	"fmt"
	"io"
	"sync"
)

//===================================================================
// Public APIs
//===================================================================

type CleanLogger struct {
	mu         sync.Mutex
	level      int
	prefix     string
	callerpath int
}

func NewDefaultCleanLogger(writer io.Writer, prefix string, logLevel int) LevelLogger {
	return &CleanLogger{
		prefix:     prefix,
		level:      logLevel,
		callerpath: 4,
	}
}

func (this *CleanLogger) Debug(v ...interface{}) {
	this.print(LOG_LEVEL_DEBUG, false, v...)
}

func (this *CleanLogger) Debugf(fmt string, v ...interface{}) {
	this.printf(LOG_LEVEL_DEBUG, false, fmt, v...)
}

func (this *CleanLogger) Trace(v ...interface{}) {
	this.print(LOG_LEVEL_TRACE, false, v...)
}

func (this *CleanLogger) Tracef(fmt string, v ...interface{}) {
	this.printf(LOG_LEVEL_TRACE, false, fmt, v...)
}

func (this *CleanLogger) Info(v ...interface{}) {
	this.print(LOG_LEVEL_INFO, false, v...)
}

func (this *CleanLogger) Infof(fmt string, v ...interface{}) {
	this.printf(LOG_LEVEL_INFO, false, fmt, v...)
}

func (this *CleanLogger) Warn(v ...interface{}) {
	this.print(LOG_LEVEL_WARN, false, v...)
}

func (this *CleanLogger) Warnf(fmt string, v ...interface{}) {
	this.printf(LOG_LEVEL_WARN, false, fmt, v...)
}

func (this *CleanLogger) Error(v ...interface{}) {
	this.print(LOG_LEVEL_ERROR, false, v...)
}

func (this *CleanLogger) Errorf(fmt string, v ...interface{}) {
	this.printf(LOG_LEVEL_ERROR, false, fmt, v...)
}

func (this *CleanLogger) Fatal(v ...interface{}) {
	this.print(LOG_LEVEL_FATAL, false, v...)
}

func (this *CleanLogger) Fatalf(fmt string, v ...interface{}) {
	this.printf(LOG_LEVEL_FATAL, false, fmt, v...)
}

//===================================================================
// Private
//===================================================================

func (this *CleanLogger) print(level int, verbose bool, v ...interface{}) {
	if level&this.level == 0 {
		return
	}
	this.mu.Lock()
	defer this.mu.Unlock()
	if verbose {
		v = append([]interface{}{getCallerMoreInfo(this.callerpath) + ": "}, v...)
	} else {
		v = append([]interface{}{getCallerInfo(this.callerpath) + ": "}, v...)
	}
	fmt.Print(v...)
}

func (this *CleanLogger) printf(level int, verbose bool, format string, v ...interface{}) {
	if level&this.level == 0 {
		return
	}
	this.mu.Lock()
	defer this.mu.Unlock()
	if verbose {
		format = getCallerMoreInfo(this.callerpath) + ": " + format
	} else {
		format = getCallerInfo(this.callerpath) + ": " + format
	}
	fmt.Printf(format+"\n", v...)
}
