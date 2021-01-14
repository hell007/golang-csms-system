package test

/**
单元测试
*/

import (
	"testing"
	"time"
)

//go test -v    显示详细过程,也可以配合-run使用
//go test -run TestInsert go测试TestInsert方法
//go test -v -cover main_test.go xorm.go 测试多个方法
func TestInsert(t *testing.T) {
	u1 := new(Userinfo)
	u1.Username = "单元测试"
	u1.Department = "财务部"
	u1.Created = time.Now()
	Insert(u1)
}
