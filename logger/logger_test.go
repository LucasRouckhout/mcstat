package logger

import "testing"

func TestCreateDebugLogger(t *testing.T) {
    l, err := NewLogger(DEBUG)

    if err != nil {
        t.Errorf(err.Error())
    }

    if l.Level != DEBUG {
        t.Errorf("Expected a logger with log level INFO\n")
    }
}

func TestCreateWrongLevelLogger(t *testing.T) {
    _, err := NewLogger(234) // Some number instead of the defined levels

    if err == nil {
        t.Errorf("Expected an error to be thrown if illegal argument is given")
    }
}
