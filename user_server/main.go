package main

import (
	"github.com/xiaoyaolaotou/logger/mylogger"
)

// 一个使用自定义日志库的用户程序

func main() {
	logger := mylogger.NewFileLogger("info", "./", "xxx.log")
	//logger := mylogger.NewConsoleLogger("info")

	for {
		logger.Debug("debug")
		logger.Info("info")
		logger.Error("error")
	}

}
