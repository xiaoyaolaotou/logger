package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type logger struct {
	fileName string
	filePath string
	file     *os.File
}

func NewFileLogger(fileName, filePath string) *logger {
	fl := &logger{
		fileName: fileName,
		filePath: filePath,
	}
	fl.initFile()
	return fl
}

// 打开文件
func (l *logger) initFile() {
	logName := filepath.Join(l.filePath, l.fileName)
	fileObj, err := os.OpenFile(logName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(fmt.Sprintf("打开日志文件错误: %s", err))
	}
	l.file = fileObj
}

func (l *logger) Debug(format string, args ...interface{}) {
	fmt.Fprintln(l.file, format)
}

func main() {
	x := NewFileLogger("./", "cunzhang.txt")
	x.Debug("123")
}
