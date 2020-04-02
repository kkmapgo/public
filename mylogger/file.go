package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLogger struct {
	Level       LogLevel
	filePath    string // 日志文件路径
	fileName    string // 日志文件名
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64 // 最大文件大小
}

func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	// 按照文件路径和文件名将文件打开
	err = fl.initFile()
	if err != nil {
		panic(err)
	}
	return fl
}

func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:%v\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:%v\n", err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

func (f *FileLogger) enable(logLevel LogLevel) bool {
	return f.Level <= logLevel
}

func (f *FileLogger) log(lv LogLevel, format string, args ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, args...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		if f.checkSize(f.fileObj) {
			newFileObj, err := f.splitFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFileObj
		}
		fmt.Fprintf(f.fileObj, "%s %s %s:%s:%s %s\n",
			now.Format("2006-01-02 15:04:05"),
			getLogString(lv),
			funcName, fileName, lineNo, msg)
		if lv >= ERROR {
			if f.checkSize(f.errFileObj) {
				newErrFileObj, err := f.splitFile(f.errFileObj)
				if err != nil {
					return
				}
				f.errFileObj = newErrFileObj
			}
			// 如果是错误日志级别以上，还需要在error日志文件中记录一份
			fmt.Fprintf(f.errFileObj, "%s %s %s:%s:%s %s\n",
				now.Format("2006-01-02 15:04:05"),
				getLogString(lv),
				funcName, fileName, lineNo, msg)
		}
	}
}

func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	// 需要切割
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed, err:%v\n", err)
		return nil, err
	}
	nowStr := time.Now().Format("20060102150405000")
	logName := path.Join(f.filePath, fileInfo.Name())
	newLogName := fmt.Sprintf("%s/%s.bak%s", logName, nowStr)
	// 1. 关闭当前日志文件
	file.Close()
	// 2. 重命名
	os.Rename(logName, newLogName)
	// 3. 打开新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file failed, err:%v\n", err)
		return nil, err
	}
	// 4. 打开的新文件赋值给f.fileObj
	return fileObj, nil
}

func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed, err:%v\n", err)
		return false
	}
	return fileInfo.Size() > f.maxFileSize
}

func (f *FileLogger) close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

func (f *FileLogger) Debug(format string, args ...interface{}) {
	f.log(DEBUG, format, args)
}

func (f *FileLogger) Info(format string, args ...interface{}) {
	f.log(INFO, format, args)
}

func (f *FileLogger) Warn(format string, args ...interface{}) {
	f.log(WARN, format, args)
}

func (f *FileLogger) Error(format string, args ...interface{}) {
	f.log(ERROR, format, args)
}

func (f *FileLogger) Fatal(format string, args ...interface{}) {
	f.log(FATAL, format, args)
}
