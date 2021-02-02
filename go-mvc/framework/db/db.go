package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/core"

	//"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/kataras/golog"
	"sync"

	"go-mvc/framework/conf"
)

var (
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
	once         sync.Once
)

/**
获取数据库连接的url
*/
func GetConnURL() (url string) {
	url = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		conf.Global.MysqlMasterUser,
		conf.Global.MysqlMasterPassword,
		conf.Global.MysqlMasterHost,
		conf.Global.MysqlMasterPort,
		conf.Global.MysqlMasterDatabase,
		conf.Global.Charset)
	return
}

/**
创建xorm Engine
*/
func newEngine() *xorm.Engine {
	engine, err := xorm.NewEngine(conf.Global.MysqlDialect, GetConnURL())
	if err != nil {
		golog.Fatalf("db.Engine, %s", err)
		return nil
	}
	//defer engine.Close()

	// Debug模式，打印全部的SQL语句，帮助对比，看ORM与SQL执行的对照关系
	engine.ShowSQL(conf.Global.MysqlShowSql)
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, conf.Global.MysqlPrefix)
	engine.SetTableMapper(tbMapper)
	engine.SetTZLocation(conf.SysTimeLocation)
	// 性能优化的时候才考虑，加上本机的SQL缓存
	//cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	//engine.SetDefaultCacher(cacher)

	return engine
}

// 主库，单例模式
func MasterEngine() *xorm.Engine {
	once.Do(func() {
		masterEngine = newEngine()
	})
	return masterEngine
}

// 从库，单例模式
func SlaveEngine() *xorm.Engine {
	once.Do(func() {
		slaveEngine = newEngine()
	})
	return slaveEngine
}
