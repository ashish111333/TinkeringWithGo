package exercises

import (
	"fmt"
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
			fmt.Println()

		}
		if stackInt.Pop() != 2 {
			t.FailNow()
		}
		if stackInt.Pop() != 1 {
			t.FailNow()
		}

	})

}
