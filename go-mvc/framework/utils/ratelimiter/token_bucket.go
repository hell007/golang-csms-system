package ratelimiter

import (
	"sync"
	"time"
)

/*
令牌桶算法（Token Bucket）是网络流量整形（Traffic Shaping）和速率限制（Rate Limiting）中最常使用的一种算法。典型情况下，令牌桶算法用来控制发送到网络上的数据的数目，并允许突发数据的发送。

我们有一个固定的桶，桶里存放着令牌（token）。一开始桶是空的，系统按固定的时间（rate）往桶里添加令牌，直到桶里的令牌数满，多余的请求会被丢弃。当请求来的时候，从桶里移除一个令牌，如果桶是空的则拒绝请求或者阻塞。

令牌桶有以下特点：
令牌按固定的速率被放入令牌桶中
桶中最多存放 B 个令牌，当桶满时，新添加的令牌被丢弃或拒绝
如果桶中的令牌不足 N 个，则不会删除令牌，且请求将被限流（丢弃或阻塞等待）
令牌桶限制的是平均流入速率（允许突发请求，只要有令牌就可以处理，支持一次拿3个令牌，4个令牌...），并允许一定程度突发流量。
*/
type TokenBucket struct {
	rate      int64 //固定的token放入速率, r/s
	capacity  int64 //桶的容量
	tokens    int64 //桶中当前token数量
	timestamp int64 //桶上次放token的时间戳 s

	lock sync.Mutex
}

func (tb *TokenBucket) Allow() bool {
	tb.lock.Lock()
	defer tb.lock.Unlock()

	now := time.Now().Unix()
	tb.tokens = tb.tokens + (now-tb.timestamp)*tb.rate // 先添加令牌
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}
	tb.timestamp = now

	if tb.tokens > 0 {
		// 还有令牌，领取令牌
		tb.tokens--
		return true
	} else {
		// 没有令牌,则拒绝
		return false
	}
}

func (tb *TokenBucket) Set(r, c int64) {
	tb.rate = r
	tb.capacity = c
	tb.tokens = 0
	tb.timestamp = time.Now().Unix()
}
