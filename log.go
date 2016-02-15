package log

import (
	"os"
	ilog "log"
)

var (
	debug   *ilog.Logger
	info    *ilog.Logger
	error   *ilog.Logger
)

var ShowDebug bool

func init() {
	debug = ilog.New(os.Stdout, "", 0)
	info = ilog.New(os.Stdout, "", 0)
	error = ilog.New(os.Stderr, "", 0)
}

func Debug(v ...interface{}) {
	if ShowDebug {
		debug.Print(v...)
	}
}

func Debugf(format string, v ...interface{}) {
	if ShowDebug {
		debug.Printf(format, v...)
	}
}

func Info(v ...interface{}) {
	info.Print(v...)
}

func Infof(format string, v ...interface{}) {
	info.Printf(format, v...)
}

func Error(v ...interface{}) {
	error.Print(v...)
}

func Errorf(format string, v ...interface{}) {
	error.Printf(format, v...)
}