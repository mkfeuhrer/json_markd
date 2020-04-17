package json_markd

import (
	"os"

	"github.com/sirupsen/logrus"
)

type key int

const (
	logTracerKey key = iota
)

type Logger struct {
	*logrus.Logger
}

var Log *Logger

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

// SetupLogger sets up logger for use. You can use all the logrus library functions eg. logger.Log.Error()
//
// Please refer to https://github.com/sirupsen/logrus
func SetupLogger() {
	level, err := logrus.ParseLevel("DEBUG")
	panicIfError(err)

	logrusVar := &logrus.Logger{
		Out:       os.Stderr,
		Formatter: &logrus.JSONFormatter{},
		Hooks:     make(logrus.LevelHooks),
		Level:     level,
	}

	Log = &Logger{logrusVar}
}
