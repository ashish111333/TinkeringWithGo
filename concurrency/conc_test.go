package concurrency

import (
	"testing"
)

func TestAddSliceItems(t *testing.T) {
	s := RandIntSlice(40, 50)
	if AddSLiceItems(s) != (AddSliceItemsC(2, s)) {
		t.Fatalf("outputs don't match")
	}
}
func BenchmarkAddSLiceItems(b *testing.B) {
	s := RandIntSlice(10000, 100000)
	for i := 0; i < b.N; i++ {
		res := AddSLiceItems(s)
		b.Logf("Iteration -- %d result --- %d", i, res)
	}
}
