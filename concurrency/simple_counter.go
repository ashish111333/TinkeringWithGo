package concurrency

import "sync/atomic"

// simple concurrent counter
type Counter struct {
	Count int64
}

func (c *Counter) Inc() {
	atomic.AddInt64(&c.Count, 1)
}

func NewCounter(countUpTo int64) *Counter {
	return &Counter{}
}
