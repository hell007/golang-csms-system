package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

//logger是一种相对高级的用法, 对于一个大型项目, 往往需要一个全局的logrus实例，即logger对象来记录项目所有的日志
// logrus提供了New()函数来创建一个logrus的实例。
// 项目中，可以创建任意数量的logrus实例。
var log = logrus.New()

func main() {
	// 为当前logrus实例设置消息的输出，同样地，
	// 可以设置logrus实例的输出到任意io.writer
	log.Out = os.Stdout

	// 为当前logrus实例设置消息输出格式为json格式。
	// 同样地，也可以单独为某个logrus实例设置日志级别和hook，这里不详细叙述。
	log.Formatter = &logrus.JSONFormatter{}

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("这是一种高级的用法")
}
