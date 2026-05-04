package concurrency

import "sync/atomic"

// Problem: Sharded Counter

type Sharded_counter struct {
	shards []int64
}

func (sc *Sharded_counter) Inc(worker_id int64) {
	go_routine_idx := int64(worker_id) % int64(len(sc.shards))
	atomic.AddInt64(&sc.shards[go_routine_idx], 1)
}
func (sc *Sharded_counter) Total() int64 {
	var result int64
	for _, v := range sc.shards {
		result += v
	}
	return result
}

func NewShardedCounter(shards int64) *Sharded_counter {
	sc := &Sharded_counter{}
	sc.shards = []int64{}

	for range shards {
		sc.shards = append(sc.shards, 0)
	}
	return sc

}
