package concurrency

import (
	"hash/fnv"
	"sync"
)

type kvShard struct {
	mx sync.RWMutex
	m  map[string]int64
}

type ShardedKV struct {
	shards []kvShard
}

func NewShardedKV(numShards int) *ShardedKV {
	if numShards <= 0 {
		numShards = 1
	}

	shards := make([]kvShard, numShards)
	for i := range shards {
		shards[i].m = make(map[string]int64)
	}

	return &ShardedKV{shards: shards}
}

func (s *ShardedKV) Set(key string, val int64) {
	shard := s.getShard(key)
	shard.mx.Lock()
	shard.m[key] = val
	shard.mx.Unlock()
}

func (s *ShardedKV) Get(key string) (int64, bool) {
	shard := s.getShard(key)
	shard.mx.RLock()
	val, ok := shard.m[key]
	shard.mx.RUnlock()
	return val, ok
}

func (s *ShardedKV) Delete(key string) {
	shard := s.getShard(key)
	shard.mx.Lock()
	delete(shard.m, key)
	shard.mx.Unlock()
}

func (s *ShardedKV) Increment(key string, delta int64) int64 {
	shard := s.getShard(key)
	shard.mx.Lock()
	shard.m[key] += delta
	res := shard.m[key]
	shard.mx.Unlock()
	return res
}

func (s *ShardedKV) Len() int {
	total := 0
	for i := range s.shards {
		shard := &s.shards[i]
		shard.mx.RLock()
		total += len(shard.m)
		shard.mx.RUnlock()
	}
	return total
}

func (s *ShardedKV) getShard(key string) *kvShard {
	idx := int(hashKey(key) % uint64(len(s.shards)))
	return &s.shards[idx]
}

func hashKey(key string) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(key))
	return h.Sum64()
}
