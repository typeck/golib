package timer

import "github.com/go-redis/redis"

type OptionFunc func(t *Timer)

func (f OptionFunc) apply(t *Timer) {
	f(t)
}

//设置redis，如果用到分布式锁，则需要设置
func SetRedis(r redis.Cmdable) Options {
	return OptionFunc(func(t *Timer) {
		t.redisCli = r
	})
}

type EntryOption interface {
	apply(entry *entry)
}

type entryOptFunc func(*entry)

func (f entryOptFunc) apply(e *entry) {
	f(e)
}

// 分布式锁
func WithRedisLock(key string) EntryOption {
	return entryOptFunc(func(e *entry) {
		e.key = key
	})
}

// 任务本地串行执行
func WithLocalSerial() EntryOption {
	return entryOptFunc(func(e *entry) {
		e.localSerial = true
	})
}
