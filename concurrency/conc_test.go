package concurrency

import (
	"testing"
)

var s []int64 = RandIntSlice(80000000, 100000)

func TestAddSliceItems(t *testing.T) {
	if AddSLiceItems(s) != (AddSliceItemsC(2, s)) {
		t.Fatalf("outputs don't match")
	}
}
func BenchmarkName(b *testing.B) {
	b.Run("AddSLiceiItemsC", func(b *testing.B) {
		for b.Loop() {
			AddSliceItemsC(12, s)
		}
	})
	b.Run("AddSliceItems", func(b *testing.B) {
		for b.Loop() {
			AddSLiceItems(s)
		}
	})

}
