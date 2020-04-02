package mylogger

import (
	"errors"
	"strings"
)

type LogLevel uint16

const (
	UNKNOWN LogLevel = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

type LogType uint8

const (
	CONSOLE LogType = iota
	FILE
)

type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
}

func getLogString(lv LogLevel) string {
	switch lv {
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
		return "INFO"
	}
}

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "warn":
		return WARN, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		return UNKNOWN, errors.New("unknown log level")
	}
}

//func NewLogger(lt LogType, args ...interface{}) (*Logger, error) {
//	switch lt {
//	case CONSOLE:
//		if len(args) != 1 {
//			return nil, errors.New("error parameters for console log type")
//		}
//		return NewConsoleLogger(args[0].(string)), nil
//	case FILE:
//		if len(args) != 4 {
//			return nil, errors.New("error parameters for file log type")
//		}
//		return NewFileLogger(args[0], args[1], args[2], args[3]), nil
//	default:
//		return nil, errors.New("invalid log type")
//	}
//}
