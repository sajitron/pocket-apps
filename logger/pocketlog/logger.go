package pocketlog

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Logger is used to log information
type Logger struct {
	threshold Level
	output    io.Writer
}

// Debugf formats and prints a message if the log level is debug or higher.
func (l *Logger) Debugf(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}

	if l.threshold <= LevelDebug {
		l.logf(LevelDebug, format, args...)
	}
}

// Infof formats and prints a message if the log level is info or higher.
func (l *Logger) Infof(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}

	if l.threshold <= LevelInfo {
		l.logf(LevelInfo, format, args...)
	}
}

// Errorf formats and prints a message if the log level is error or higher.
func (l *Logger) Errorf(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}

	if l.threshold <= LevelError {
		l.logf(LevelError, format, args...)
	}
}

// New returns a logger, ready to log at the required threshold.
// Give it a list of configuration functions to tune it at your will.
// The default output is Stdout.
func New(threshold Level, opts ...Option) *Logger {
	lgr := &Logger{
		threshold: threshold,
		output:    os.Stdout,
	}

	for _, configFunc := range opts {
		configFunc(lgr)
	}

	return lgr
}

// logf prints the message to the output.
// Add decorations here, if any.
func (l *Logger) logf(logLevel Level, format string, args ...any) {
	contents := fmt.Sprintf(format, args...)
	msg := message{
		Level:   logLevel.String(),
		Message: contents,
	}

	formattedMessage, err := json.Marshal(msg)

	if err != nil {
		_, _ = fmt.Fprintf(l.output, "unable to format message for %v\n", contents)
	}

	_, _ = fmt.Fprintln(l.output, string(formattedMessage))
}

// Logf formats and prints a message if the log level is high enough
func (l *Logger) Logf(logLevel Level, format string, args ...any) {
	if l.threshold <= logLevel {
		l.logf(logLevel, format, args...)
	}
}

// Message represents the JSON structure of the logged messages.
type message struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}
