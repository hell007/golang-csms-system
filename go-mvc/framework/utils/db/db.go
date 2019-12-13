package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/core"

	//"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/go-yaml/yaml"
	"github.com/kataras/golog"
	"sync"

	"go-mvc/framework/conf"
	"go-mvc/framework/utils/files"
)

var (
	Db MySql // for casbin to user
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
	once         sync.Once
)

// db配置结构体
type MySql struct {
	Master DbInfo
	Slave  DbInfo
}

type DbInfo struct {
	Dialect      string `yaml:"dialect"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Database     string `yaml:"database"`
	Charset      string `yaml:"charset"`
	ShowSql      bool   `yaml:"showSql"`
	Prefix		 string `yaml:"prefix"`
	LogLevel     string `yaml:"logLevel"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
}

// 读取配置文件
func (d *MySql) GetConf() *MySql {
	yamlFile, err := files.LoadFile(conf.GetConfigPath() + "conf.yaml")
	if err != nil {
		golog.Errorf("LoadFile db config error!! %s", err)
	}

	if err = yaml.Unmarshal(yamlFile, d); err != nil {
		golog.Errorf("Unmarshal db config error!! %s", err)
	}
	return d
}

// 获取数据库连接的url
// true：master主库
func GetConnURL(c *DbInfo) (url string) {
	url = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
		c.Charset)
	return
}

// 创建xorm
func newEngine() *xorm.Engine{
	var d MySql
	master := d.GetConf().Master
	engine, err := xorm.NewEngine(master.Dialect, GetConnURL(&master))
	if err != nil {
		golog.Fatalf("db.Engine, %s", err)
		return nil
	}

	// Debug模式，打印全部的SQL语句，帮助对比，看ORM与SQL执行的对照关系
	engine.ShowSQL(master.ShowSql)
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, master.Prefix)
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