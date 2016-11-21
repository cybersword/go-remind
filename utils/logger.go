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

// Notice info and debug
func (sl *SimpleLogger) Notice(msg string) {
	sl.normal.Println(msg)
}

// Fatal warnning and fatal
func (sl *SimpleLogger) Fatal(msg string) {
	sl.normal.Println(msg)
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
	psl.normal = log.New(fileLog, "[Info]", log.Lshortfile|log.LstdFlags)
	psl.wf = log.New(fileLogWF, "[Fatal]", log.Lshortfile|log.LstdFlags)
	return psl
}
