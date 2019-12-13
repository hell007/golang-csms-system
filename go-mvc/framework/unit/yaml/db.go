package main

import (
	"fmt"
	"github.com/kataras/golog"
	//"gopkg.in/yaml.v2"
	"github.com/go-yaml/yaml"

	"go-mvc/framework/conf"
	"go-mvc/framework/utils/files"

)

type Mysql struct {
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
	LogLevel     string `yaml:"logLevel"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
}

func (d *Mysql) getConf() *Mysql {
	yamlFile, err := files.LoadFile(conf.GetConfigPath() +"conf.yaml")
	if err != nil {
		golog.Fatalf("LoadFile reids config error!! %s", yamlFile, err)
	}

	if err = yaml.Unmarshal(yamlFile, d); err != nil {
		golog.Fatalf("Unmarshal db config error!! %s", err)
	}
	return d
}

func main() {
	var d Mysql
	d.getConf()
	fmt.Println("master===", d.Master)
	fmt.Println("host===", d.Master.Host)
}