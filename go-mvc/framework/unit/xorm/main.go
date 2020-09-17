package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/core"

	//"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/kataras/golog"
	"time"
	"errors"
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

//删
func Del(uid int) {
	userinfo := new(Userinfo)
	x.Id(uid).Delete(userinfo)
}

//改
func update(uid int, userinfo *Userinfo) bool {
	affected, err := x.ID(uid).Update(userinfo)
	if err != nil {
		golog.Fatal("错误:", err)
	}
	if affected == 0 {
		return false
	}
	return true
}

//查
func getinfo(uid int) *Userinfo {
	userinfo := &Userinfo{Uid: uid}
	is, err := x.Get(userinfo)
	if !is {
		golog.Fatal("搜索结果不存在!", err)
	}
	return userinfo
}

// 事务处理操作
type Test struct {
	Id         int       `xorm:"int(10) pk autoincr"`
	Name       string    `xorm:"varchar(20) unique"`
	Price      int    `xorm:"int(10)"`
}

func transfer(id1, id2 int , price int) error {

	t1 := &Test{Id: id1}
	t2 := &Test{Id: id2}

	is, err := x.Get(t1)
	if !is {
		golog.Fatal("t1不存在!", err)
	}

	ok, err := x.Get(t2)
	if !ok {
		golog.Fatal("t2不存在!", err)
	}

	t1.Price -= price
	t2.Price += price

	// 事务处理开始
	sess:= x.NewSession()
	defer sess.Close()

	if err := sess.Begin() ; err != nil {
		return errors.New("fail to session begin")
	}

	// 让1成功
	if _, err := sess.Id(id1).Cols("price").Update(t1) ; err != nil {
		sess.Rollback()
		return errors.New("fail to update 1")
	}

	// 让2不成功， 如果1操作成功就事务出错，反之，事务就正确
	if _, err := sess.Id(id2).Cols("price").Update(t2); err != nil {
		sess.Rollback()
		return errors.New("fail to update 2")
	}

	return sess.Commit()
}


func main() {

	x.ShowSQL(true) // 显示SQL的执行, 便于调试分析
	x.SetMapper(core.SameMapper{}) // 设置表名下划线转驼峰
	//x.Sync2(new(Userinfo))

	// 1.查询
	//fmt.Println("\n查询")

	// Query 最原始的也支持SQL语句查询，返回的结果类型为 []map[string][]byte
	// QueryString 返回 []map[string]string
	// QueryInterface 返回 []map[string]interface{}

	//q1, _ := x.Query("select * from userinfo")
	//fmt.Println("Query q1= ", q1, "\n")

	//q2, _ := x.QueryString("select username from userinfo where uid=8")
	//fmt.Println("Query q2= ", q2, "\n")

	//q3, _ := x.QueryInterface("select * from userinfo")
	//.Println("Query q3= ", q3, "\n")

	// 单条记录
	//uinfo := new(Userinfo)

	// q4, _ := x.Get(uinfo)
	// fmt.Println("Get q4= ", q4, "\n")

	//q5, _ := x.Where("username = ?", "曹操").Desc("uid").Get(uinfo)
	//fmt.Println("Get q5= ", q5, "\n")

	//var username string
	//q6, _ := x.Table("userinfo").Where("uid = ?", 8).Cols("username").Get(&username)
	//fmt.Println("Get q6= ", q6, "\n")

	//q7, _ := x.SQL("select username from userinfo").Get(&username)
	//fmt.Println("Get q7= ", q7, "\n")

	//var valuesMap = make(map[string]string)
	//q8, _ := x.Table("userinfo").Where("uid = ?", 8).Get(&valuesMap)
	//fmt.Println("Get q8= ", q8, "\n")

	// cols := {}
	// var valuesSlice = make([]interface{}, len(cols))
	// q9, _ := x.Table("userinfo").Where("uid = ?", 8).Cols(cols...).Get(&valuesSlice)
	// fmt.Println("Get q9= ", q9, "\n")

	// fmt.Println("\n struct查询")
	// info := getinfo(8)
	// fmt.Println("select = ", info, "\n")

	// 多条记录
	// var users []UserDetail
	// err := x.Table("user").Select("user.*, detail.*")
	// Join("INNER", "detail", "detail.user_id = user.id").
	// 	Where("user.name = ?", name).Limit(10, 0).
	// 	Find(&users)
	// SELECT user.*, detail.* FROM user INNER JOIN detail WHERE user.name = ? limit 10 offset 0

	// 2.添加
	fmt.Println("\n添加")

	// u1 := new(Userinfo)
	// u1.Username = "曹丕"
	// u1.Department = "财务部"
	// u1.Created = time.Now()
	//Insert(u1)

	// 批量添加
	// u2 := new(Userinfo)
	// u2.Username = "曹丕22"
	// u2.Department = "财务部"
	// u2.Created = time.Now()

	// u3 := new(Userinfo)
	// u3.Username = "曹丕33"
	// u3.Department = "财务部"
	// u3.Created = time.Now()

	// row, _ := x.Insert(u2, u3)
	// fmt.Println(" row = ", row)

	// 3.更改
	fmt.Println("\n更改")

	// u4 := new(Userinfo)
	// u4.Username = "刘小备1"
	// u4.Department = "刘家桥"
	// update(9, u4)

	// 4.删除
	fmt.Println("\n删除")

	// Del(10)
	//getinfo(1)

	// 5.事务

	err := transfer(1,2, 200) //800 700
	fmt.Println("====", err)

}