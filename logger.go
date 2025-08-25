package uaparser

import (
	"log"
	"os"
)

// Logger is the interface for logging functionality.
// Users can implement this interface to provide their own logger.
type Logger interface {
	// Infof logs an info message with format.
	Infof(format string, args ...interface{})
}

// DefaultLogger is a simple implementation of Logger using the standard log package.
type DefaultLogger struct {
	logger *log.Logger
}

// NewDefaultLogger creates a new DefaultLogger.
func NewDefaultLogger() *DefaultLogger {
	return &DefaultLogger{
		logger: log.New(os.Stdout, "[uaparser] ", log.LstdFlags),
	}
}

// Infof implements the Logger interface.
func (l *DefaultLogger) Infof(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}

// NoOpLogger is a logger that does nothing.
// Useful when you want to disable logging completely.
type NoOpLogger struct{}

// NewNoOpLogger creates a new NoOpLogger.
func NewNoOpLogger() *NoOpLogger {
	return &NoOpLogger{}
}

// Infof implements the Logger interface but does nothing.
func (l *NoOpLogger) Infof(format string, args ...interface{}) {
	// Do nothing
}
