package redis

import (
	"github.com/go-redis/redis"
	"github.com/kataras/golog"

	"go-mvc/framework/conf"
)

//redis单机客户端
var client *redis.Client

//redis集群客户端
var redisClusterClient *redis.ClusterClient

// redis 初始化
func init() {

	//判断是否为集群配置
	if conf.GlobalConfig.RedisClusterState {
		//ClusterClient是一个Redis集群客户机，表示一个由0个或多个底层连接组成的池。它对于多个goroutine的并发使用是安全的。
		redisClusterClient = redis.NewClusterClient(&redis.ClusterOptions{
			Password: conf.GlobalConfig.RedisClusterPassword,
			Addrs:    conf.GlobalConfig.RedisClusterAddrs,
		})

		//Ping
		ping, err := redisClusterClient.Ping().Result()
		if err != nil {
			golog.Warnf("Redis Ping %s %s", ping, err)
			golog.Warn("Redis集群服务未启动，拒绝连接，连接失败...")
		}
		golog.Info("Redis集群服务已启动...")


	} else {
		//Redis客户端，由零个或多个基础连接组成的池。它对于多个goroutine的并发使用是安全的。
		//更多参数参考Options结构体
		client = redis.NewClient(&redis.Options{
			Addr:     conf.GlobalConfig.RedisSingleAddr,
			Password: conf.GlobalConfig.RedisSinglePassword, // no password set
			DB:       conf.GlobalConfig.RedisSingleDb,       // use default DB
		})

		//Ping
		pong, err := client.Ping().Result()
		if err != nil {
			golog.Warnf("Redis Ping %s %s", pong, err)
			golog.Warn("Redis服务未启动，拒绝连接，连接失败...")
		}
		golog.Info("Redis服务已启动...")
	}
}
