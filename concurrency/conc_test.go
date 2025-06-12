package concurrency

import (
	"testing"
)

var s []int64 = RandIntSlice(800000000000120, 50)

func TestAddSliceItems(t *testing.T) {
	if AddSLiceItems(s) != (AddSliceItemsC(6, s)) {
		t.Fatalf("outputs don't match")
	}
}
func BenchmarkAddSliceItems(b *testing.B) {
	b.Run("AddSLiceiItemsC", func(b *testing.B) {
		for b.Loop() {
			AddSliceItemsC(50, s)
		}
	})
	b.Run("AddSliceItems", func(b *testing.B) {
		for b.Loop() {
			AddSLiceItems(s)
		}
	})

}
