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

// if fixed is set to true
func RandIntSlice(l, n, cap int64, fixed bool) []int64 {
	s := make([]int64, cap)
	var i int64
	if fixed {
		num := mr.Int63n(n)
		for range l {
			s = append(s, num)
		}
		return s
	}
	for i = 0; i < l; i++ {
		s = append(s, mr.Int63n(n))
	}
	return s
}
func OsThreadsUnderRuntime() {

}
func GetRuntimeStats() {

}
