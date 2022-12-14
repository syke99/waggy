package waggy

import (
	"errors"
	"fmt"
	"os"
)

// ParentLoggerOverrider overrides the parent *Logger in a Handler
type ParentLoggerOverrider = func() bool

// OverrideParentLogger
func OverrideParentLogger() ParentLoggerOverrider {
	return func() bool {
		return true
	}
}

// Logger is used for writing to a log
type Logger struct {
	logLevel string
	key      string
	message  string
	err      string
	vals     map[string]interface{}
	log      *os.File
}

// LogLevel allows you to set the level
// to be used in a *Logger
type LogLevel int

const (
	Info LogLevel = iota
	Debug
	Warning
	Fatal
	Error
	Warn
	All
	Off
)

func (l LogLevel) level() string {
	return []string{
		"INFO",
		"DEBUG",
		"WARNING",
		"FATAL",
		"ERROR",
		"WARN",
		"ALL",
		"OFF",
	}[l]
}

// NewLogger returns a new *Logger with the provided log file (if
// log is not nil) and the provided logLevel.
func NewLogger(logLevel LogLevel, log *os.File) *Logger {
	l := Logger{
		logLevel: logLevel.level(),
		key:      "",
		message:  "",
		err:      "",
		vals:     make(map[string]interface{}),
		log:      log,
	}

	return &l
}

// SetLogFile set a specific file for the logger to Write to.
// You must mount the volume that this file resides in whenever
// you configure your WAGI server via your modules.toml file
// for a *Logger to be able to write to the provided file
func (l *Logger) SetLogFile(log *os.File) error {
	if log == nil {
		return errors.New("no log file provided")
	}

	l.log = log

	return nil
}

// Level update the level of a *Logger
func (l *Logger) Level(level LogLevel) *Logger {
	l.logLevel = level.level()

	return l
}

// Err provide an error to the *Logger to be logged
func (l *Logger) Err(err error) *Logger {
	if err == nil {
		return l
	}

	l.err = err.Error()

	return l
}

// Val add a value with the corresponding key to be logged by the *Logger
func (l *Logger) Val(key string, val any) *Logger {
	l.vals[key] = val

	return l
}

// Msg provide a message with a key to be logged and then
// logs the constructed log messed to the set *os.File (or default to os.Stderr)
func (l *Logger) Msg(key string, msg string) (int, error) {
	l.key = key
	l.message = msg

	lm := make(map[string]string)

	lm["level"] = l.logLevel

	if l.key != "" {
		lm[l.key] = l.message
	}

	for k, v := range l.vals {
		if k != "" {
			lm[k] = fmt.Sprintf("%s,%v", lm[k], v)
		}
	}

	if l.err != "" {
		lm["error"] = l.err
	}

	return l.log.Write([]byte(buildLogJSON(lm)))
}

func buildLogJSON(log map[string]string) string {
	logJSON := "{"

	for key, value := range log {
		if key != "" {
			if logJSON[:1] != "{" {
				logJSON = fmt.Sprintf("%s,", logJSON)
			}

			logJSON = fmt.Sprintf("%[1]s \"%[2]s\": \"%[3]s\"", logJSON, key, value)
		}
	}

	return fmt.Sprintf("%s }", logJSON)
}
