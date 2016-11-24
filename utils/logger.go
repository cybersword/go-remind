// Package utils implements some helper functions, which are esay to use.

// logger.go - defines a type, SimpleLogger, with methods for output log.
// The Notice function writes the log message to xxx.log with [Notice] prefix.
// The Fatal function writes the log message to both xxx.wf.log and xxx.log with [Fatal] prefix.
// The difference between utils.Notice with log.Println is write to file or echo on screen.

package utils

import (
	"log"
	"os"
	"path"
)

// SimpleLogger write log
type SimpleLogger struct {
	normal *log.Logger
	wf     *log.Logger
}

var psl *SimpleLogger

// Notice writes the log message to xxx.log with [Notice] prefix.
func Notice(msg string) {
	GetSimpleLogger().Notice(msg)
}

// Fatal  writes the log message to both xxx.wf.log and xxx.log with [Fatal] prefix.
func Fatal(v ...interface{}) {
	GetSimpleLogger().Fatal(v)
}

// Notice info and debug
func (sl *SimpleLogger) Notice(msg string) {
	sl.normal.Println(msg)
}

// Fatal warnning and fatal
func (sl *SimpleLogger) Fatal(v ...interface{}) {
	sl.wf.Println(v)
	sl.normal.SetPrefix("[Fatal]")
	sl.normal.Println(v)
	sl.normal.SetPrefix("[Notice]")
}

// SetOutput reset log writer
func (sl *SimpleLogger) SetOutput(l *log.Logger, level int) {
	switch level {
	case 1:
		sl.normal = l
	case 2:
		sl.wf = l
	}
}

// GetSimpleLogger single instance
func GetSimpleLogger() *SimpleLogger {
	if psl != nil {
		return psl
	}
	dirLog := "."
	fileLog, _ := os.Create(path.Join(dirLog, "debug.log"))
	fileLogWF, err := os.Create(path.Join(dirLog, "debug.wf.log"))
	if err != nil {
		panic(err)
	}

	pl1 := log.New(fileLog, "[Notice]", log.Lshortfile|log.LstdFlags)
	pl2 := log.New(fileLogWF, "[Fatal]", log.Lshortfile|log.LstdFlags)
	psl = &SimpleLogger{pl1, pl2}
	return psl
}
