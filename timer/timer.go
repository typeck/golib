package timer

import (
	"errors"
	"log"
	"os"
	"runtime/debug"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

type Timer struct {
	work      []*item
	redisCli  redis.Cmdable
	mux 	  sync.Mutex
	log       Logger
}

type Options interface {
	apply(t *Timer)
}

type Logger interface {
	Printf(format string, v ...interface{})
}

type entry struct {
	duration    time.Duration
	key         string
	localSerial bool
}

func (e *entry) useRedisLock() bool {
	if e.key != "" {
		return true
	}
	return false
}

type item struct {
	f     func()
	entry *entry
}

var DefaultTimer = NewTimer()

var defaultLog = log.New(os.Stdout, "[timer]", log.Lshortfile|log.Ldate|log.Ltime)

func NewTimer(opts ...Options) *Timer {
	t := &Timer{
		redisCli:  nil,
		log:       defaultLog,
	}
	for _, opt := range opts {
		opt.apply(t)
	}
	return t
}

func (t *Timer)WithOptions(opts ...Options) {
	for _, opt := range opts {
		opt.apply(t)
	}
}

func (t *Timer) Add(fn func(), dur time.Duration, opts ...EntryOption) {
	if dur <= 0 {
		return
	}
	var entry = &entry{duration: dur}
	for _, opt := range opts {
		opt.apply(entry)
	}
	t.work = append(t.work, &item{f: fn, entry: entry})
}

func (t *Timer) Start()  {
	for _, item := range t.work {
		itemTmp := item
		go func() {
			err := t.run(itemTmp)
			if err != nil {
				panic(err)
			}
		}()
	}
}

func (t *Timer) tryLock(key string, ex time.Duration) (bool, error) {

	ok, err := t.redisCli.SetNX(key, 1, ex).Result()

	if err == redis.Nil {
		return false, nil
	}

	if err != nil {
		return false, err
	}
	return ok, nil
}

func (t *Timer) run(item *item) error {
	defer t.recovery()
	var frequency = item.entry.duration
	entry := item.entry
	if entry.useRedisLock() {
		if t.redisCli == nil {
			return errors.New("nil redis")
		}
		err := t.redisCli.Ping().Err()
		if err != nil {
			return err
		}
		frequency = time.Duration(int64(entry.duration) / 20)
		if frequency == 0 {
			return errors.New("duration is to short")
		}
	}
	ticker := time.NewTicker(frequency)
	tc := ticker.C
	for {
		select {
		case <-tc:
			if entry.useRedisLock() {
				ok, err := t.tryLock(entry.key, entry.duration)
				if err != nil {
					t.log.Printf("get redis lock error: %v", err)
					continue
				}
				if !ok {
					continue
				}
			}
			if entry.localSerial {
				t.serial(item)
			}else {
				item.f()
			}
		}
	}
}

func (t *Timer)serial(item *item) {
	t.mux.Lock()
	defer t.mux.Unlock()
	item.f()

}

func (t *Timer) recovery() {
	if err := recover(); err != nil {
		t.log.Printf("recover: %s", string(debug.Stack()))
	}
}
