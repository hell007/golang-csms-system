package redis

import (
	"github.com/go-redis/redis"
	"time"
)

//set
func Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	if MRedis.Cluster.State {
		return redisClusterClient.Set(key, value, expiration)
	} else {
		return client.Set(key, value, expiration)
	}
}

//get
func Get(key string) *redis.StringCmd {
	if MRedis.Cluster.State {
		return redisClusterClient.Get(key)
	} else {
		return client.Get(key)
	}
}

//del
func Del(key string) *redis.IntCmd {
	if MRedis.Cluster.State {
		return redisClusterClient.Del(key)
	} else {
		return client.Del(key)
	}
}

//incr
func Incr(key string) *redis.IntCmd{
	if MRedis.Cluster.State {
		return redisClusterClient.Incr(key)
	} else {
		return client.Incr(key)
	}
}



//订阅
func Subscribe(channels ...string) *redis.PubSub {
	if MRedis.Cluster.State {
		return redisClusterClient.Subscribe(channels...)
	} else {
		return client.Subscribe(channels...)
	}
}

//发布
func Publish(channel string, message interface{}) *redis.IntCmd {
	if MRedis.Cluster.State {
		return redisClusterClient.Publish(channel, message)
	} else {
		return client.Publish(channel, message)
	}
}