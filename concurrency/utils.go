package concurrency

import (
	"crypto/rand"
	mr "math/rand"
)

func RandString(prefix string) string {
	if prefix == "" {
		return rand.Text()
	}
	return prefix + rand.Text()
}

func RandIntSlice(l, n int64) []int64 {
	s := []int64{}
	var i int64
	for i = 0; i < l; i++ {
		s = append(s, mr.Int63n(n))
	}
	return s
}
