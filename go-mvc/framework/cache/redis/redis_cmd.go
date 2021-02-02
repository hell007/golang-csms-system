package redis

import (
	"github.com/go-redis/redis"
	"go-mvc/framework/conf"
	"time"
)

//set
func Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	if conf.Global.RedisClusterState {
		return redisClusterClient.Set(key, value, expiration)
	} else {
		return client.Set(key, value, expiration)
	}
}

//get
func Get(key string) *redis.StringCmd {
	if conf.Global.RedisClusterState {
		return redisClusterClient.Get(key)
	} else {
		return client.Get(key)
	}
}

//del
func Del(key string) *redis.IntCmd {
	if conf.Global.RedisClusterState {
		return redisClusterClient.Del(key)
	} else {
		return client.Del(key)
	}
}

//incr
func Incr(key string) *redis.IntCmd {
	if conf.Global.RedisClusterState {
		return redisClusterClient.Incr(key)
	} else {
		return client.Incr(key)
	}
}

//订阅
func Subscribe(channels ...string) *redis.PubSub {
	if conf.Global.RedisClusterState {
		return redisClusterClient.Subscribe(channels...)
	} else {
		return client.Subscribe(channels...)
	}
}

//发布
func Publish(channel string, message interface{}) *redis.IntCmd {
	if conf.Global.RedisClusterState {
		return redisClusterClient.Publish(channel, message)
	} else {
		return client.Publish(channel, message)
	}
}
