package banjo

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"
)

// Logger struct
//
// Provides logging with timestamps and different logging levels
// With saving all messages to log file
//
type Logger struct {
	filePath string
}

// CreateLogger function
//
// Returns Logger with default filePath
//
func CreateLogger() Logger {
	path := os.Getenv("BANJO_LOG")

	if path == "" {
		path = "banjo_out.log"
	}

	return Logger{filePath: path}
}

// Info function
//
// Log message with log level INFO
//
// Params:
// - message {string}
//
// Response:
// - None
//
func (logger Logger) Info(message string) {
	logger.log("INFO", message)
}

// Critical function
//
// Log message with log level CRITICAL
//
// Params:
// - message {string}
//
// Response:
// - None
//
func (logger Logger) Critical(message string) {
	logger.log("CRITICAL", message)
}

// Warning function
//
// Log message with log level WARNING
//
// Params:
// - message {string}
//
// Response:
// - None
//
func (logger Logger) Warning(message string) {
	logger.log("WARNING", message)
}

// Error function
//
// Log message with log level ERROR
//
// Params:
// - message {string}
//
// Response:
// - None
//
func (logger Logger) Error(message string) {
	logger.log("ERROR", message)
}

// log function
//
// Puts logging information to file and to io
//
// Params:
// - logLevel {string} Message Log Level
// - message 	{string} Message to log
//
// Response:
// - None
//
func (logger Logger) log(level string, message string) {
	logLine := formatMessage(level, message)
	err := logger.writeToLogFile(logLine + "\n")

	if err != nil {
		str := fmt.Sprintf("Error while trying save logs to file:\nError: %v", err)
		fmt.Println(formatMessage("ERROR", str))
	}

	fmt.Println(logLine)
}

// writeToLogFile function
//
// Saving log message to file
//
// Params:
// - line {string}
//
// Response:
// - err {Error} returns error if something went wrong
//
func (logger Logger) writeToLogFile(line string) error {
	_, err := os.Stat(logger.filePath)

	if os.IsNotExist(err) {
		_, err := os.Create(logger.filePath)
		if err != nil {
			return err
		}
	}

	f, err := os.OpenFile(logger.filePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err := f.WriteString(line); err != nil {
		return err
	}

	return nil
}

// formatMessage function
//
// Create Log line with timestampz, level and message
//
// Params:
// - level {string}
// - message  {string}
//
// Response:
// - response {string}
//
func formatMessage(level string, message string) string {
	var buffer bytes.Buffer
	t := time.Now()

	buffer.WriteString("[BANjO] ")
	buffer.WriteString(t.Format(time.RFC3339))
	buffer.WriteString(" | ")
	buffer.WriteString(strings.ToUpper(level))
	buffer.WriteString(" | ")
	buffer.WriteString(message)

	return buffer.String()
}
