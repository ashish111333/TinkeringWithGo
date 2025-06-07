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
