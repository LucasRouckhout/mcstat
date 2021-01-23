// A very simple leveled logger implementation.
//
// By its very design it is not very extendible.
// You have three logging levels: DEBUG, INFO and ERROR
// and it uses the build in logging package to log to stdout.
package logger

import (
	"log"
)

type Logger struct {
	Level uint
}

const (
	DEBUG uint = 0
	INFO  uint = 1
	ERROR uint = 2
)

// Creates a new logger object with the given logging level.
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

// Logs a DEBUG level log.
func (logger Logger) Debug(message string) {
	logger.Debugf("%s\n", message)
}

// Logs a DEBUG level log with formated string.
func (logger Logger) Debugf(format string, v ...interface{}) {
	if logger.Level <= DEBUG {
		log.Printf("DEBUG: "+format, v...)
	}
}

// Logs a INFO level log.
func (logger Logger) Info(message string) {
	logger.Infof("%s\n", message)
}

// Logs a INFO level log with formated string.
func (logger Logger) Infof(format string, v ...interface{}) {
	if logger.Level <= INFO {
		log.Printf("INFO: "+format, v...)
	}
}

// Logs a ERROR level log.
func (logger Logger) Error(message string) {
	logger.Errorf("%s\n", message)
}

// Logs a ERROR level log with formated string.
func (logger Logger) Errorf(format string, v ...interface{}) {
	log.Printf("ERROR: "+format, v...)
}
