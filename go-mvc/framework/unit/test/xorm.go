package test

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/core"

	//"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/kataras/golog"
	"time"
)

//定义结构体(xorm支持双向映射)
type Userinfo struct {
	Uid        int       `xorm:"int(10) pk autoincr"` //指定主键并自增
	Username   string    `xorm:"varchar(64) unique"`  //唯一的
	Department string    `xorm:"varchar(64)"`
	Created    time.Time `xorm:"created"`
	//Version    int     `xorm:"version"` //乐观锁
}

//定义orm引擎
var x *xorm.Engine

//创建orm引擎
func init() {
	var err error

	x, err = xorm.NewEngine("mysql", "root:admin@tcp(127.0.0.1:3306)/jie?charset=utf8")
	//x.ShowSQL(true) // 显示SQL的执行, 便于调试分析
	x.SetMapper(core.SameMapper{}) // 设置表名下划线转驼峰

	if err != nil {
		golog.Fatal("数据库连接失败:", err)
	}
	if err := x.Sync(new(Userinfo)); err != nil {
		golog.Fatal("数据表同步失败:", err)
	}
}

//增
func Insert(userinfo *Userinfo) {
	affected, err := x.Insert(userinfo)
	if err != nil {
		golog.Fatal("添加失败:", err)
	}
	fmt.Println(affected)
}
