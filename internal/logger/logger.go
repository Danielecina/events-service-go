package logger

import (
	"log"
	"os"
)

// AppLogger is the main application logger
var AppLogger *log.Logger

func init() {
	AppLogger = log.New(os.Stdout, "[APP] ", log.LstdFlags|log.Lshortfile)
}

// Info logs an info message
func Info(message string, args ...interface{}) {
	AppLogger.Printf("[INFO] "+message, args...)
}

// Error logs an error message
func Error(message string, args ...interface{}) {
	AppLogger.Printf("[ERROR] "+message, args...)
}

// Fatal logs an error message and exits the program
func Fatal(message string, args ...interface{}) {
	AppLogger.Printf("[FATAL] "+message, args...)
	os.Exit(1)
}

// Debug logs a debug message
func Debug(message string, args ...interface{}) {
	AppLogger.Printf("[DEBUG] "+message, args...)
}
