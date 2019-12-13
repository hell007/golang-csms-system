package redis

import (
	"github.com/go-redis/redis"
	"github.com/go-yaml/yaml"
	"github.com/kataras/golog"

	"go-mvc/framework/utils/files"
	"go-mvc/framework/conf"
)

var MRedis MyRedis
//redis单机客户端
var client *redis.Client
//redis集群客户端
var redisClusterClient *redis.ClusterClient

// redis配置结构体
type MyRedis struct {
	Single  SingleInfo
	Cluster ClusterInfo
}

type SingleInfo struct {
	Addr     string   `yaml:"addr"`
	Password string   `yaml:"password"`
	Db       int      `yaml:"db"`
	PoolSize int      `yaml:"poolSize"`
}

type ClusterInfo struct {
	Addrs     []string  `yaml:"addrs"`
	Password  string    `yaml:"password"`
	State     bool      `yaml:"state"`
}

// 读取配置文件
func (redis *MyRedis) getConf() *MyRedis {
	yamlFile, err := files.LoadFile(conf.GetConfigPath() + "conf.yaml")
	if err != nil {
		golog.Errorf("LoadFile reids config error!! %s", err)
	}

	if err = yaml.Unmarshal(yamlFile, redis); err != nil {
		golog.Errorf("Unmarshal reids config error!! %s", err)
	}
	return redis
}

// 设置
func init() {
	var r MyRedis
	r.getConf()
	golog.Infof("redis 初始化信息：[%s]",r.Single)

	//判断是否为集群配置
	if r.Cluster.State {
		//ClusterClient是一个Redis集群客户机，表示一个由0个或多个底层连接组成的池。它对于多个goroutine的并发使用是安全的。
		redisClusterClient = redis.NewClusterClient(&redis.ClusterOptions{
			Password: r.Cluster.Password,
			Addrs:    r.Cluster.Addrs,
		})
		//Ping
		ping, err := redisClusterClient.Ping().Result()
		golog.Infof("Redis Ping %s %s",ping, err)

	} else {
		//Redis客户端，由零个或多个基础连接组成的池。它对于多个goroutine的并发使用是安全的。
		//更多参数参考Options结构体
		client = redis.NewClient(&redis.Options{
			Addr:     r.Single.Addr,
			Password: r.Single.Password, // no password set
			DB:       r.Single.Db,       // use default DB
		})
		//Ping
		pong, err := client.Ping().Result()
		golog.Infof("Redis Ping %s %s", pong, err)
	}
}
