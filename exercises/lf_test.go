package exercises

import (
	"slices"
	"sync"
	"testing"
)

func TestLfStack(t *testing.T) {
	// let's test integer stack
	stackInt := NewLfStack[int]()
	t.Run("stack_test", func(t *testing.T) {
		stackInt.Push(1)
		stackInt.Push(2)
		stackInt.Push(3)
		if stackInt.Pop() != 3 {
			t.FailNow()

		}
		if stackInt.Pop() != 2 {
			t.FailNow()
		}
		if stackInt.Pop() != 1 {
			t.FailNow()
		}

	})

}

// test lock free stack on concurrent environment
func TestLfStackConc(t *testing.T) {
	stck := NewLfStack[int]()

	var wg sync.WaitGroup
	wg.Add(3)
	for i := 1; i < 4; i++ {
		go func(i int) {
			defer wg.Done()
			stck.Push(i)

		}(i)

	}
	wg.Wait()
	// now check if all elements exist
	if !stck.exists(1) {
		t.Fail()

	}
	if !stck.exists(2) {
		t.Fail()
	}
	if !stck.exists(3) {
		t.Fail()
	}

}

// test stack with mutex
func TestStackMx(t *testing.T) {
	stck := NewStackMx[int]()
	stck.Push(1)
	stck.Push(2)
	stck.Push(3)
	if stck.Pop() != 3 {
		t.FailNow()
	}
	if stck.Pop() != 2 {
		t.Failed()

	}
	if stck.Pop() != 1 {
		t.FailNow()
	}

}

func TestStackMxConc(t *testing.T) {
	stc := NewStackMx[int]()
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 1; i < 4; i++ {
		go func(i int) {
			defer wg.Done()
			stc.Push(i)

		}(i)
	}
	wg.Wait()
	slc := []int{1, 2, 3}
	if !slices.Contains(slc, stc.Pop()) {
		t.FailNow()

	}
	if !slices.Contains(slc, stc.Pop()) {
		t.FailNow()
	}
	if !slices.Contains(slc, stc.Pop()) {
		t.FailNow()
	}

}
