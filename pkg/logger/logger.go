package logger

import "log"

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
	logger struct{}
)

func NewLogger() Logger {
	return &logger{}
}

func (l *logger) Info(intr ...interface{}) {
	log.Print(intr...)
}

func (l *logger) Infof(format string, intr ...interface{}) {
	log.Printf(format, intr...)
}

func (l *logger) Error(err error) {
	log.Print(err)
}

func (l *logger) Errorf(format string, err error) {
	log.Printf(format, err)
}

func (l *logger) Warn(intr ...interface{}) {
	log.Print(intr...)
}

func (l *logger) Warnf(format string, intr ...interface{}) {
	log.Printf(format, intr...)
}
