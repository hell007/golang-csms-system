package rateLimiter

import (
	"sync"
	"time"
)

type TokenBucket struct {
	rate      int64 //固定的token放入速率, r/s
	capacity  int64 //桶的容量
	tokens    int64 //桶中当前token数量
	timestamp int64 //桶上次放token的时间戳 s

	lock sync.Mutex
}

func NewBucket(r, c int64) *TokenBucket {
	return &TokenBucket{
		rate:      r,
		capacity:  c,
		tokens:    0,
		timestamp: time.Now().Unix(),
	}
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
