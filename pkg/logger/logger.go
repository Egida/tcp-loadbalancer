package logger

import (
	log "github.com/amupxm/xmus-logger"
)

type (
	// Logger is an interface for logging.
	Logger interface {
		// Info logs the given information.
		Info(intr ...interface{})
		// Infof formats and logs the given information.
		Infof(format string, intr ...interface{})
		// Error logs the given error.
		Error(err error)
		// Errorf formats and logs the given error.
		Errorf(format string, err error)
		// Warn logs the given warning.
		Warn(intr ...interface{})
		// Warnf formats and logs the given warning.
		Warnf(format string, intr ...interface{})
	}
	logger struct {
		l log.Logger
	}
)

var l log.Logger

func init() {
	l = log.CreateLogger(&log.Options{
		LogLevel: log.Info,
	})
}

func NewLogger(prefix string) Logger {
	return &logger{
		l: l.BeginWithPrefix(prefix),
	}
}

func (l *logger) Info(intr ...interface{}) {
	l.l.Info(intr...)
}

func (l *logger) Infof(format string, intr ...interface{}) {
	l.l.Infof(format, intr...)
}

func (l *logger) Error(err error) {
	l.l.Error(err)
}

func (l *logger) Errorf(format string, err error) {
	l.l.Errorf(format, err)
}

func (l *logger) Warn(intr ...interface{}) {
	l.l.Warn(intr...)
}

func (l *logger) Warnf(format string, intr ...interface{}) {
	l.l.Warnf(format, intr...)
}

func AddWhiteList(str ...string) {
	l.AddToWhitelist(str...)
}
