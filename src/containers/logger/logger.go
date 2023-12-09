package logger

import (
	"htmx-golang/containers/env"
	"htmx-golang/containers/time"
	"os"

	"github.com/charmbracelet/log"
)

var defaultLogger = log.NewWithOptions(os.Stderr, log.Options{
	ReportCaller:    false,
	ReportTimestamp: true,
	TimeFormat:      time.TimeLayout,
})

var debugLogger = log.NewWithOptions(os.Stderr, log.Options{
	ReportCaller:    true,
	ReportTimestamp: true,
	TimeFormat:      time.TimeLayoutInMilliseconds,
	Level:           log.DebugLevel,
})

func Debug(msg interface{}, keyvals ...interface{}) {
	if !env.IsEnv("ENVIRONMENT", env.PRODUCTION) {
		debugLogger.Debug(msg, keyvals...)
	}
}

func Info(msg interface{}, keyvals ...interface{}) {
	defaultLogger.Info(msg, keyvals...)
}

func Warn(msg interface{}, keyvals ...interface{}) {
	defaultLogger.Warn(msg, keyvals...)
}

func Error(msg interface{}, keyvals ...interface{}) {
	debugLogger.Error(msg, keyvals...)
}

func Fatal(msg interface{}, keyvals ...interface{}) {
	debugLogger.Fatal(msg, keyvals...)
}
