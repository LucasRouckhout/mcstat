// Package logger implements a very simple leveled logger.
// By its very design it is not very extendible.
// You have three logging levels: DEBUG, INFO and ERROR
// and it uses the build in logging package to log to stdout.
package logger

import (
	"log"
)

// Logger represents a leveled logger
type Logger struct {
	Level uint
}

const (
	// DEBUG Level
	DEBUG uint = 0

	// INFO Level
	INFO uint = 1

	// ERROR Level
	ERROR uint = 2
)

// NewLogger creates a new logger object with the given logging level.
//
// You can only use one of the predefined logging levels. Which
// are: DEBUG, INFO, ERROR
//
// You can then use this Logger object to call the respective
// Info, Debug and Error functions which will log a log of that
// level. Depending on your log level the log will be written to
// stdout or not. So for example if you have set your level to
// ERROR then logging with the Info() function will have no effect.
//
// There are also formatted equivalents to all the logging functions.
// So Info() --> Infof() will allow you to log with formatted strings
// just like the Print and Printf functions in the standard library.
func NewLogger(level uint) Logger {

	if level != DEBUG && level != INFO && level != ERROR {
		return Logger{
			Level: INFO,
		}
	}

	return Logger{
		Level: level,
	}
}

// Debug logs a DEBUG level log.
func (logger Logger) Debug(message string) {
	logger.Debugf("%s\n", message)
}

// Debugf logs a DEBUG level log with formated string.
func (logger Logger) Debugf(format string, v ...interface{}) {
	if logger.Level <= DEBUG {
		log.Printf("DEBUG: "+format, v...)
	}
}

// Info logs a INFO level log.
func (logger Logger) Info(message string) {
	logger.Infof("%s\n", message)
}

// Infof logs a INFO level log with formated string.
func (logger Logger) Infof(format string, v ...interface{}) {
	if logger.Level <= INFO {
		log.Printf("INFO: "+format, v...)
	}
}

// Error logs an ERROR level log.
func (logger Logger) Error(message string) {
	logger.Errorf("%s\n", message)
}

// Errorf logs a ERROR level log with formated string.
func (logger Logger) Errorf(format string, v ...interface{}) {
	log.Printf("ERROR: "+format, v...)
}
