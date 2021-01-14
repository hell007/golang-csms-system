package test

/**
压力测试
*/

import (
	"testing"
)

// 注意命名规范 Benchmark+首字母大写的方法名 参数固定
func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		//Insert(4, 5)
	}
}

//go test -bench BenchmarkAdd 只压测BenchmarkAdd方法
//go test -bench BenchmarkSub 只压测BenchmarkSub方法
//go test -bench .    点表示测试该路径下所有压力测试
func Benchmark_TimeConsumingFunction(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能

	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		//Division(4, 5)
	}
}
