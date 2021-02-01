package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// 设置日志格式为json格式
	log.SetFormatter(&log.JSONFormatter{})

	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	log.SetOutput(os.Stdout)

	// 设置日志级别为warn以上
	log.SetLevel(log.WarnLevel)
}

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("这是一个一般的info日志")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("这是一个warn日志")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("这是一个错误日志")
}
