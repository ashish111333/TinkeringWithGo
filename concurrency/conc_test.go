package concurrency

import (
	"testing"
)

var s []int64 = RandIntSlice(80000, 100, 80000, true)

func TestAddSliceItems(t *testing.T) {
	if AddSLiceItems(s) != (AddSliceItemsC(6, s)) {
		t.Fatalf("outputs don't match")
	}
}
func BenchmarkAddSliceItems(b *testing.B) {
	var sC int64
	b.Run("AddSLiceiItemsC", func(b *testing.B) {
		for b.Loop() {
			sC = AddSliceItemsC(3, s)
		}
	})
	b.Logf("sum --->%d", sC)
	var sCC int64
	b.Run("AddSliceItemsCButChannels", func(b *testing.B) {
		for b.Loop() {
			sCC = addSliceItemsCChannels(3, s)
		}
	})
	b.Logf("sum given by channels one-->%d", sCC)
	b.Run("AddSliceItems", func(b *testing.B) {
		for b.Loop() {
			AddSLiceItems(s)
		}
	})

}
