package json_markd

import (
	"context"
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

func WithContext(ctx context.Context, entry *logrus.Entry) context.Context {
	return context.WithValue(ctx, logTracerKey, entry)
}

func FromContext(ctx context.Context) *logrus.Entry {
	if log, ok := ctx.Value(logTracerKey).(*logrus.Entry); ok {
		return log
	}
	return logrus.NewEntry(Log.Logger)
}
