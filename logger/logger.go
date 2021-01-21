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
    DEBUG uint   = 0
    INFO uint    = 1
    ERROR uint   = 2
)

// Creates a new logger object with the given logging level.
// 
// You can only use one of the predefined logging levels. Which
// are: DEBUG, INFO, ERROR
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

// Logs an DEBUG level log.
func (logger Logger) Debug(message string) {
    logger.Debugf("%s\n", message)
}

// Logs an DEBUG level log with formated string.
func (logger Logger) Debugf(format string, v... interface{}) {
    if logger.Level <= DEBUG {
        log.Printf("DEBUG: " + format, v...)
    }
}

// Logs an INFO level log.
func (logger Logger) Info(message string) {
    logger.Infof("%s\n", message)
}

// Logs an INFO level log with formated string.
func (logger Logger) Infof(format string, v... interface{}) {
    if logger.Level <= INFO {
        log.Printf("INFO: " + format, v...)
    }
}

// Logs an ERROR level log.
func (logger Logger) Error(message string) {
    logger.Errorf("%s\n", message)
}

// Logs an ERROR level log with formated string.
func (logger Logger) Errorf(format string, v... interface{}) {
    log.Printf("ERROR: " + format, v...)
}
