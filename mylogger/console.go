package mylogger

import (
	"fmt"
	"os"
	"time"
)

// 往终端打印

type ConsoleLogger struct {
	level Level
}

// 文件日志结构造函数
func NewConsoleLogger(levelStr string) *ConsoleLogger {
	logLevel := parseLogLevel(levelStr)
	cl := &ConsoleLogger{
		level: logLevel,
	}
	return cl
}

// 将公用的记录日志的功能封装成一个单独的方法
func (f *ConsoleLogger) log(level Level, format string, args ...interface{}) {
	if f.level > level {
		return
	}
	msg := fmt.Sprintf(format, args...) // 得到用户要记录的日志
	// 日志格式：[时间][文件:行号][函数名][日志级别] 日志信息
	nowStr := time.Now().Format("2006-01-02 15:04:05.000")
	logLeverStr := getLevelStr(level) // 传进来的日志级别
	logMsg := fmt.Sprintf("[%s][%s]%s", nowStr, logLeverStr, msg)
	fmt.Fprintln(os.Stdout, logMsg)
	// 如果是error或者fatal级别的日志还要记录到errFile里

}

// 方法
// Debug
func (f *ConsoleLogger) Debug(format string, args ...interface{}) {
	f.log(DebugLevel, format, args...)
}

// Info
func (f *ConsoleLogger) Info(format string, args ...interface{}) {
	f.log(InfoLevel, format, args...)
}

// Warn
func (f *ConsoleLogger) Warn(format string, args ...interface{}) {
	f.log(WarningLevel, format, args...)
}

// Error
func (f *ConsoleLogger) Error(format string, args ...interface{}) {
	f.log(ErrorLevel, format, args...)
}
