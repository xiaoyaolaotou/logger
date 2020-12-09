package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 这是往文件里面写日志代码

// FileLogger 文件日志结构体
type FileLogger struct {
	level    Level // 日志级别
	fileName string
	filePath string
	file     *os.File
	errFile  *os.File
}

// 文件日志结构造函数
func NewFileLogger(levelStr, fileName, filePath string) *FileLogger {
	logLevel := parseLogLevel(levelStr)
	fl := &FileLogger{
		level:    logLevel,
		fileName: fileName,
		filePath: filePath,
	}
	fl.initFfile() // 根据上面的文件路径和文件名打开日志文件， 把文件句柄赋值给构体
	return fl
}

// 将指定的日志文件打开, 赋值给结构体
func (f *FileLogger) initFfile() {
	logName := path.Join(f.filePath, f.fileName)
	// 打开文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Errorf("打开日志文件%s失败", logName))
	}
	f.file = fileObj
	// 打开错误的日志文件
	errLogName := fmt.Sprintf("%s.err", logName)
	errFileObj, err := os.OpenFile(errLogName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Errorf("打开日志文件%s失败", errLogName))
	}
	f.errFile = errFileObj
}

// 将公用的记录日志的功能封装成一个单独的方法
func (f *FileLogger) log(level Level, format string, args ...interface{}) {
	if f.level > level {
		return
	}
	msg := fmt.Sprintf(format, args...) // 得到用户要记录的日志
	// 日志格式：[时间][文件:行号][函数名][日志级别] 日志信息
	nowStr := time.Now().Format("2006-01-02 15:04:05.000")
	logLeverStr := getLevelStr(level) // 传进来的日志级别
	logMsg := fmt.Sprintf("[%s][%s]%s", nowStr, logLeverStr, msg)
	fmt.Fprintln(f.file, logMsg)
	// 如果是error或者fatal级别的日志还要记录到errFile里
	if level >= ErrorLevel {
		fmt.Fprintln(f.errFile, logMsg)
	}
}

// 方法
// Debug
func (f *FileLogger) Debug(format string, args ...interface{}) {
	f.log(DebugLevel, format, args...)
}

// Info
func (f *FileLogger) Info(format string, args ...interface{}) {
	f.log(InfoLevel, format, args...)
}

// Warn
func (f *FileLogger) Warn(format string, args ...interface{}) {
	f.log(WarningLevel, format, args...)
}

// Error
func (f *FileLogger) Error(format string, args ...interface{}) {
	f.log(ErrorLevel, format, args...)
}
