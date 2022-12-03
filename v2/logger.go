package v2

import (
	"errors"
	"io"
	"os"
)

type Logger struct {
	logLevel   string
	key        string
	message    string
	err        string
	loggerPath string
	log        io.Writer
}

func NewLogger(logLevel string) *Logger {
	l := Logger{
		logLevel: logLevel,
		key:      "",
		message:  "",
		err:      "",
		log:      os.Stderr,
	}

	return &l
}

func (l *Logger) SetLogFile(log *os.File) error {
	if log == nil {
		return errors.New("no log file provided")
	}

	l.log = log

	return nil
}

func (l *Logger) Err(err error) *Logger {
	if err == nil {
		return l
	}

	l.err = err.Error()

	return l
}

func (l *Logger) Msg(key string, msg string) *Logger {
	l.key = key
	l.message = msg

	return l
}

func (l *Logger) Log() {
	// TODO: build out log message and write it to l.log
}
