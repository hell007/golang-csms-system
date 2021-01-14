package main

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"github.com/kataras/golog"

	"go-mvc/framework/utils/files"
)

// redis配置结构体
type Redis struct {
	Single  SingleInfo
	Cluster ClusterInfo
}

type SingleInfo struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
	PoolSize int    `yaml:"poolSize"`
}

type ClusterInfo struct {
	Addrs    []string `yaml:"addrs"`
	Password string   `yaml:"password"`
	State    bool     `yaml:"state"`
}

// 读取配置文件
func (redis *Redis) getConf() *Redis {
	yamlFile, err := files.LoadFile("./conf.yaml")
	if err != nil {
		golog.Fatalf("LoadFile reids config error!! %s", err)
	}

	if err = yaml.Unmarshal(yamlFile, redis); err != nil {
		golog.Fatalf("Unmarshal reids config error!! %s", err)
	}
	return redis
}

func main() {
	var redis Redis
	redis.getConf()
	fmt.Println("Single===", redis.Single)
	fmt.Println("Addrs===", redis.Cluster.Addrs)
}
