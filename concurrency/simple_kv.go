package concurrency

import "sync"

type SimpleKV struct {
	mx sync.RWMutex
	m  map[string]int64
}

func NewSimpleKV() *SimpleKV {
	return &SimpleKV{m: make(map[string]int64)}
}

func (s *SimpleKV) Set(key string, val int64) {
	s.mx.Lock()
	s.m[key] = val
	s.mx.Unlock()
}

func (s *SimpleKV) Get(key string) (int64, bool) {
	s.mx.RLock()
	val, ok := s.m[key]
	s.mx.RUnlock()
	return val, ok
}

func (s *SimpleKV) Delete(key string) {
	s.mx.Lock()
	delete(s.m, key)
	s.mx.Unlock()
}

func (s *SimpleKV) Increment(key string, delta int64) int64 {
	s.mx.Lock()
	s.m[key] += delta
	res := s.m[key]
	s.mx.Unlock()
	return res
}

func (s *SimpleKV) Len() int {
	s.mx.RLock()
	res := len(s.m)
	s.mx.RUnlock()
	return res
}
