package logging

import (
	"log"

	"github.com/fatih/color"
)

/*
Logger represents an instance of a log structure
*/
type Logger struct {
	ApplicationName string
	LogLevel        LogType

	colorEnabled bool
	logLevelInt  int
}

/*
Debug writes a debug entry to the log
*/
func (logger *Logger) Debug(message ...interface{}) {
	logger.writeLog(DEBUG, message)
}

/*
Debugf writes a formatted debug entry to the log
*/
func (logger *Logger) Debugf(message string, args ...interface{}) {
	logger.writeLogf(DEBUG, message, args...)
}

/*
DisableColors turns of console coloring
*/
func (logger *Logger) DisableColors() {
	logger.colorEnabled = false
}

/*
EnableColors turns on console coloring
*/
func (logger *Logger) EnableColors() {
	logger.colorEnabled = true
}

/*
Error writes an error message to the log
*/
func (logger *Logger) Error(message ...interface{}) {
	logger.writeLog(ERROR, message)
}

/*
Errorf writes a formatted error entry to the log
*/
func (logger *Logger) Errorf(message string, args ...interface{}) {
	logger.writeLogf(ERROR, message, args...)
}

/*
Fatal writes a fatal log messge
*/
func (logger *Logger) Fatal(message ...interface{}) {
	logger.writeLog(FATAL, message)
}

/*
Fatalf writes a formatted fatal entry to the log
*/
func (logger *Logger) Fatalf(message string, args ...interface{}) {
	logger.writeLogf(FATAL, message, args...)
}

/*
Info writes an information message to the log
*/
func (logger *Logger) Info(message ...interface{}) {
	logger.writeLog(INFO, message)
}

/*
Infof writes a formatted info entry to the log
*/
func (logger *Logger) Infof(message string, args ...interface{}) {
	logger.writeLogf(INFO, message, args...)
}

/*
NewLogger creates a new Logger instance with a specific application name.
The application name is a prefix used in all log messges.
*/
func NewLogger(applicationName string) *Logger {
	return &Logger{
		ApplicationName: applicationName,
		LogLevel:        DEBUG,

		colorEnabled: false,
		logLevelInt:  int(DEBUG),
	}
}

/*
NewLoggerWithMinimumLevel creates a new Logger instance that only logs
messages with a specified log type level or higher.
*/
func NewLoggerWithMinimumLevel(applicationName string, logLevel LogType) *Logger {
	return &Logger{
		ApplicationName: applicationName,
		LogLevel:        logLevel,

		colorEnabled: false,
		logLevelInt:  int(logLevel),
	}
}

/*
Warning writes a warning message to the log
*/
func (logger *Logger) Warning(message ...interface{}) {
	logger.writeLog(WARN, message)
}

/*
Warningf writes a formatted warning entry to the log
*/
func (logger *Logger) Warningf(message string, args ...interface{}) {
	logger.writeLogf(WARN, message, args...)
}

func (logger *Logger) writeLog(logType LogType, message ...interface{}) {
	logLevelInt := int(logType)

	logColor := logType.Color()

	if logger.colorEnabled {
		color.Set(logColor)
	}

	if logLevelInt >= logger.logLevelInt {
		log.SetPrefix(logger.ApplicationName + ": " + logType.String() + " - ")

		for _, item := range message {
			log.Print(item)
		}
	}

	if logger.colorEnabled {
		color.Unset()
	}
}

func (logger *Logger) writeLogf(logType LogType, message string, args ...interface{}) {
	logLevelInt := int(logType)
	logColor := logType.Color()

	if logger.colorEnabled {
		color.Set(logColor)
	}

	if logLevelInt >= logger.logLevelInt {
		log.SetPrefix(logger.ApplicationName + ": " + logType.String() + " - ")
		log.Printf(message+"\n", args...)
	}

	if logger.colorEnabled {
		color.Unset()
	}
}
