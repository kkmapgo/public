package mylogger

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

// Logger 日志结构体
type ConsoleLogger struct {
	Level LogLevel
}

func getInfo(skip int) (funcName, fileName string, lineNo int) {
	caller, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime.Caller failed")
		return
	}
	funcName = runtime.FuncForPC(caller).Name()
	return strings.Split(funcName, ".")[1], path.Base(file), line
}

func NewConsoleLogger(levelStr string) *ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return &ConsoleLogger{
		Level: level,
	}
}

func (c *ConsoleLogger) log(lv LogLevel, format string, args ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, args...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Printf("%s %s %s:%s:%s %s\n", now.Format("2006-01-02 15:04:05"),
			getLogString(lv),
			funcName, fileName, lineNo, msg)
	}
}

func (c *ConsoleLogger) enable(logLevel LogLevel) bool {
	return c.Level <= logLevel
}

func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	c.log(DEBUG, format, args)
}

func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	c.log(INFO, format, args)
}

func (c *ConsoleLogger) Warn(format string, args ...interface{}) {
	c.log(WARN, format, args)
}

func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	c.log(ERROR, format, args)
}

func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	c.log(FATAL, format, args)
}
