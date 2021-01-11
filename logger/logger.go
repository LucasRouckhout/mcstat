// A very simple leveled logger implementation.
//
// By its very design it is not very extendible.
// You have three logging levels: DEBUG, INFO and ERROR
// and it uses the build in logging package to log to stdout.
package logger

import (
	"errors"
	"fmt"
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
func NewLogger(level uint) (Logger, error) {
    // Rather then silently resorting to the default we will throw 
    // and error if the given level does not conform to one of the 
    // specified log levels. If the user ignores the error it can
    // still use the returned Logger.
    if level != DEBUG && level != INFO && level != ERROR {
        return Logger{},
        errors.New(fmt.Sprintf("Illegal argument for Logger: %d\n", level))
    }

    return Logger{
        Level: level,
    }, nil
}

// Logs an DEBUG level log.
func (logger Logger) Debug(message string) {
    if logger.Level >= DEBUG {
        logger.Debugf("%s\n", message)
    }
}

// Logs an DEBUG level log with formated string.
func (logger Logger) Debugf(format string, v... interface{}) {
    if logger.Level >= INFO {
        log.Printf(format, v...)
    }
}

// Logs an INFO level log.
func (logger Logger) Info(message string) {
    if logger.Level >= INFO {
        logger.Infof("%s\n", message)
    }
}

// Logs an INFO level log with formated string.
func (logger Logger) Infof(format string, v... interface{}) {
    if logger.Level >= INFO {
        log.Printf(format, v...)
    }
}

// Logs an ERROR level log.
func (logger Logger) Error(message string) {
    if logger.Level >= ERROR {
        logger.Errorf("%s\n", message)
    }
}

// Logs an ERROR level log with formated string.
func (logger Logger) Errorf(format string, v... interface{}) {
    if logger.Level >= INFO {
        log.Printf(format, v...)
    }
}

