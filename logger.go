// Logger utilities for webin2 with structured logging support
package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// LogLevel represents the severity of a log message
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

var (
	currentLogLevel = INFO
	logger          *log.Logger
)

// String returns the string representation of the log level
func (l LogLevel) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// InitLogger initializes the logger with the specified output file
func InitLogger(logfile *os.File) {
	logger = log.New(logfile, "", 0)
	currentLogLevel = INFO
}

// SetLogLevel sets the minimum log level to output
func SetLogLevel(level LogLevel) {
	currentLogLevel = level
}

// logMessage outputs a log message with the specified level
func logMessage(level LogLevel, format string, v ...interface{}) {
	if level < currentLogLevel {
		return
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	message := fmt.Sprintf(format, v...)
	logLine := fmt.Sprintf("[%s] [%s] %s", timestamp, level.String(), message)

	if logger != nil {
		logger.Println(logLine)
	} else {
		log.Println(logLine)
	}
}

// Debug logs a debug message
func Debug(format string, v ...interface{}) {
	logMessage(DEBUG, format, v...)
}

// Info logs an info message
func Info(format string, v ...interface{}) {
	logMessage(INFO, format, v...)
}

// Warn logs a warning message
func Warn(format string, v ...interface{}) {
	logMessage(WARN, format, v...)
}

// Error logs an error message
func Error(format string, v ...interface{}) {
	logMessage(ERROR, format, v...)
}

// Fatal logs a fatal message and exits
func Fatal(format string, v ...interface{}) {
	logMessage(FATAL, format, v...)
	os.Exit(1)
}

// LogStartup logs application startup information
func LogStartup(pid int, logFilename string) {
	Info("=== Webin2 Application Started ===")
	Info("Process ID: %d", pid)
	Info("Log file: %s", logFilename)
}

// LogShutdown logs application shutdown information
func LogShutdown(success bool) {
	if success {
		Info("=== Webin2 Application Completed Successfully ===")
	} else {
		Error("=== Webin2 Application Terminated with Errors ===")
	}
}

// LogArgs logs the command line arguments
func LogArgs(jsonl, yaml, asset, account string) {
	Info("Command line arguments:")
	Info("  - jsonl: %s", jsonl)
	Info("  - yaml: %s", yaml)
	Info("  - asset: %s", asset)
	Info("  - account: %s", account)
	Info("  - password: [REDACTED]")
}

// LogActionStart logs the start of an action
func LogActionStart(index int, actionType, target string, value int) {
	Info("Action %d: type=%s, target=%s, value=%d", index, actionType, target, value)
}

// LogActionComplete logs the completion of an action
func LogActionComplete(index int, actionType string) {
	Debug("Action %d completed: %s", index, actionType)
}

// LogAssetSearch logs asset search operations
func LogAssetSearch(asset, filename string) {
	Info("Searching for asset '%s' in file '%s'", asset, filename)
}

// LogAssetFound logs when an asset is found
func LogAssetFound(asset string) {
	Info("Asset '%s' found successfully", asset)
}

// LogAssetNotFound logs when an asset is not found
func LogAssetNotFound(asset, filename string) {
	Warn("Asset '%s' not found in '%s'", asset, filename)
}

// LogFileOperation logs file operations
func LogFileOperation(operation, filename string) {
	Debug("File operation: %s - %s", operation, filename)
}

// LogBrowserConfig logs browser configuration
func LogBrowserConfig(useEdge, secret, certValidation bool) {
	Info("Browser configuration:")
	Info("  - Use Edge: %t", useEdge)
	Info("  - Secret mode: %t", secret)
	Info("  - Certificate validation: %t", certValidation)
}

// LogChromedpStart logs the start of chromedp execution
func LogChromedpStart(actionCount int) {
	Info("Starting chromedp with %d actions", actionCount)
}

// LogChromedpComplete logs the completion of chromedp execution
func LogChromedpComplete() {
	Info("Chromedp execution completed successfully")
}
