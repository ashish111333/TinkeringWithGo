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
			sC = AddSliceItemsC(8000, s)
		}
		b.Logf("sum --->%d", sC)
	})
	var sCC int64
	b.Run("AddSliceItemsCButChannels", func(b *testing.B) {
		for b.Loop() {
			sCC = addSliceItemsCChannels(8000, s)
		}
		b.Logf("sum given by channels one-->%d", sCC)
	})
	var sCMx int64
	b.Run("AddSliceItemsCMx", func(b *testing.B) {

		for b.Loop() {
			sCMx = AddSliceItemsCMx(8000, s)
		}
		b.Logf("sum given by mutex approach--->%d", sCMx)
	})

	b.Run("AddSliceItems", func(b *testing.B) {
		for b.Loop() {
			AddSLiceItems(s)
		}
	})

}

func BenchmarkUpdateVar(b *testing.B) {
	a := 0
	times := 500
	b.Run("updateVar", func(b *testing.B) {
		UpdateVar(&a, times)
	})
	e := 0
	b.Run("UpdateVarCh", func(b *testing.B) {
		UpdateVarCh(&e, times)
	})
	c := 0
	b.Run("UpdateVarMx", func(b *testing.B) {
		UpdateVarMx(&c, times)
	})
}
