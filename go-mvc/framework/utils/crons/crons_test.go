// https://www.cnblogs.com/jssyjam/p/11910851.html
package crons

import (
	"fmt"
	"os"
	"testing"
)

type testTask struct {
}

func (t *testTask) Run() {
	fmt.Println("hello world")
}

// $ go test -v -cover crons_test.go crons.go
func TestNewCrontab(t *testing.T) {
	crontab := NewCrontab()
	// 实现接口的方式添加定时任务
	task := &testTask{}
	if err := crontab.AddByID("1", "* * * * *", task); err != nil {
		fmt.Printf("error to add crontab task 1:%s", err)
		os.Exit(-1)
	}

	// 添加函数作为定时任务
	taskFunc := func() {
		fmt.Println("hello world")
	}
	if err := crontab.AddByFunc("2", "* * * * *", taskFunc); err != nil {
		fmt.Printf("error to add crontab task 2:%s", err)
		os.Exit(-1)
	}
	crontab.Start()
	select {}
}

//自定义封装
//在上述的使用方法基础上,基于我的实际需求,我对cron库进行了简单封装,主要为实现下面几个需求:
//1、管理所有的定时任务,需要记录定时任务的编号和相关信息
//2、停止一个定时任务
//3、支持添加函数类型和接口类型任务

//注:
//task id 是唯一的,实际开发可以使用 uuid
//这个封装是并发安全的

//不足:
//未支持初始化参数
//定时任务的错误采集,如果某个定时任务出错,应该能够获取到错误信息(这里指的是错误不是 panic)
//panic 恢复操作可以参考 withChain和cron.Recover