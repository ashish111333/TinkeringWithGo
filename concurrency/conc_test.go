package concurrency

import (
	"testing"
)

var s []int64 = RandIntSlice(8000000, 100, 8000000, false)

func TestAddSliceItems(t *testing.T) {
	if AddSLiceItems(s) != (AddSliceItemsC(12, s)) {
		t.Fatalf("outputs don't match")
	}
}
func BenchmarkAddSliceItems(b *testing.B) {
	var sC int64
	b.Run("AddSLiceItemsC", func(b *testing.B) {
		for b.Loop() {
			sC = AddSliceItemsC(12, s)
		}
		b.Logf("sum --->%d", sC)
	})
	var sCC int64
	b.Run("AddSliceItemsCButChannels", func(b *testing.B) {
		for b.Loop() {
			sCC = addSliceItemsCChannels(12, s)
		}
		b.Logf("sum given by channels one-->%d", sCC)
	})
	var sCMx int64
	b.Run("AddSliceItemsCMx", func(b *testing.B) {

		for b.Loop() {
			sCMx = AddSliceItemsCMx(12, s)
		}
		b.Logf("sum given by mutex approach--->%d", sCMx)
	})

	b.Run("AddSliceItems", func(b *testing.B) {
		for b.Loop() {
			AddSLiceItems(s)
		}
	})

}
