package generics

import (
	"testing"
)

func TestStack(t *testing.T) {
	slc := []int{1, 2, 3}
	s := NewStack[int]()
	for _, v := range slc {
		s.Push(v)
	}
	if !assertEqual(s.Stacklen(), len(slc)) {
		t.Fail()
	}
	for i := len(slc) - 1; i > 0; i-- {
		if _, v := s.Pop(); !assertEqual(v, slc[i]) {
			t.Fail()
		}
	}

}
