package logger

import (
	"fmt"
	"log"
)

func printLog(logType, msg string) {
	log.Println("[" + logType + "] " + msg)
}

// Info : logger for information type
func Info(logTag, requestD, format string, v ...interface{}) {
	printLog(logTag+"] [info", fmt.Sprintf("[requestID: %s] - "+format, requestD, v))
}

// Warn : logger for information type
func Warn(logTag, requestD, format string, v ...interface{}) {
	printLog(logTag+"] [warn", fmt.Sprintf("[requestID: %s] - "+format, requestD, v))
}

// Error : logger for information type
func Error(logTag, requestD, format string, v ...interface{}) {
	printLog(logTag+"] [error", fmt.Sprintf("[requestID: %s] - "+format, requestD, v))
}
