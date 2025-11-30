package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

// Level represents log level
type Level string

const (
	LevelDebug Level = "DEBUG"
	LevelInfo  Level = "INFO"
	LevelWarn  Level = "WARN"
	LevelError Level = "ERROR"
)

// Logger provides structured logging
type Logger struct {
	serviceName string
	level       Level
}

// LogEntry represents a structured log entry
type LogEntry struct {
	Timestamp   string                 `json:"timestamp"`
	Level       Level                  `json:"level"`
	Service     string                 `json:"service"`
	Message     string                 `json:"message"`
	TraceID     string                 `json:"trace_id,omitempty"`
	MessageID   string                 `json:"message_id,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// NewLogger creates a new logger instance
func NewLogger(serviceName string, level Level) *Logger {
	return &Logger{
		serviceName: serviceName,
		level:       level,
	}
}

// log writes a log entry
func (l *Logger) log(level Level, message string, fields map[string]interface{}) {
	if !l.shouldLog(level) {
		return
	}

	entry := LogEntry{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Level:     level,
		Service:   l.serviceName,
		Message:   message,
		Metadata:  fields,
	}

	// Extract special fields
	if traceID, ok := fields["trace_id"].(string); ok {
		entry.TraceID = traceID
		delete(fields, "trace_id")
	}
	if messageID, ok := fields["message_id"].(string); ok {
		entry.MessageID = messageID
		delete(fields, "message_id")
	}

	jsonBytes, err := json.Marshal(entry)
	if err != nil {
		log.Printf("Error marshaling log entry: %v", err)
		return
	}

	fmt.Fprintln(os.Stdout, string(jsonBytes))
}

// shouldLog determines if a message should be logged based on level
func (l *Logger) shouldLog(level Level) bool {
	levels := map[Level]int{
		LevelDebug: 0,
		LevelInfo:  1,
		LevelWarn:  2,
		LevelError: 3,
	}
	return levels[level] >= levels[l.level]
}

// Debug logs a debug message
func (l *Logger) Debug(message string, fields map[string]interface{}) {
	l.log(LevelDebug, message, fields)
}

// Info logs an info message
func (l *Logger) Info(message string, fields map[string]interface{}) {
	l.log(LevelInfo, message, fields)
}

// Warn logs a warning message
func (l *Logger) Warn(message string, fields map[string]interface{}) {
	l.log(LevelWarn, message, fields)
}

// Error logs an error message
func (l *Logger) Error(message string, fields map[string]interface{}) {
	l.log(LevelError, message, fields)
}

// WithTraceID returns a new logger with trace ID
func (l *Logger) WithTraceID(traceID string) *Logger {
	// This is a simplified version; in production, you'd want to maintain context
	return l
}

// WithMessageID returns a new logger with message ID
func (l *Logger) WithMessageID(messageID string) *Logger {
	// This is a simplified version; in production, you'd want to maintain context
	return l
}
