package crons

import (
	"fmt"
	"os"
	"testing"
)

type testTask struct {
}

func (t *testTask) Run() {
	fmt.Println("tast1:实现接口的方式添加定时任务")
}

// $ go test -v -cover crons_test.go crons.go
func TestNewCronManager(t *testing.T) {
	cronManager := NewCronManager()
	// 1、实现接口的方式添加定时任务
	task := &testTask{}
	if err := cronManager.AddByID("1", "* * * * *", task); err != nil {
		fmt.Printf("添加接口定时任务出错:%s", err)
		os.Exit(-1)
	}

	///2、添加函数作为定时任务
	taskFunc := func() {
		fmt.Println("tast2:添加函数作为定时任务")
	}
	if err := cronManager.AddByFunc("2", "* * * * *", taskFunc); err != nil {
		fmt.Printf("添加函数定时任务出错:%s", err)
		os.Exit(-1)
	}
	cronManager.Start()
	select {}
}
