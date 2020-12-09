package mylogger

import "strings"

// 日志库文件
// Level 是一个自定义类型 代表日志级别
type Level uint16

// 定义具体的日志级别常量
const (
	DebugLevel Level = iota
	InfoLevel
	WarningLevel
	ErrorLevel
	FatalLevel
)

// 根据传进来的Level 获取对应的字符串
func getLevelStr(level Level) string {
	switch level {
	case DebugLevel:
		return "Debug"
	case InfoLevel:
		return "Info"
	case WarningLevel:
		return "Warn"
	case ErrorLevel:
		return "Error"
	case FatalLevel:
		return "Fatal"
	default:
		return "Debug"
	}
}

// 根据用户传入的字符串类型的日志级， 解析出对应的LEVEL
func parseLogLevel(levelStr string) Level {
	levelStr = strings.ToLower(levelStr) // 将字符串转全小写
	switch levelStr {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn":
		return WarningLevel
	case "error":
		return ErrorLevel
	case "fatal":
		return FatalLevel
	default:
		return DebugLevel
	}
}
