package concurrency

import (
	"testing"
)

var s []int64 = RandIntSlice(800, 100, 800, false)

func TestAddSliceItems(t *testing.T) {
	if AddSLiceItems(s) != (AddSliceItemsC(6, s)) {
		t.Fatalf("outputs don't match")
	}
}
func BenchmarkAddSliceItems(b *testing.B) {
	var sC int64
	b.Run("AddSLiceItemsC", func(b *testing.B) {
		for b.Loop() {
			sC = AddSliceItemsC(8, s)
		}
		b.Logf("sum --->%d", sC)
	})

	var sCC int64
	b.Run("AddSliceItemsCButChannels", func(b *testing.B) {
		for b.Loop() {
			sCC = addSliceItemsCChannels(8, s)
		}
		b.Logf("sum given by channels one-->%d", sCC)
	})

	b.Run("AddSliceItems", func(b *testing.B) {
		for b.Loop() {
			AddSLiceItems(s)
		}
	})

}
