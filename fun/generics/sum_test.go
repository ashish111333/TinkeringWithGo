package generics

import "testing"

func TestSum(t *testing.T) {
	s1 := []int{1, 2, 4}
	s2 := []string{"hello", " world"}
	t.Run("sum_test", func(t *testing.T) {
		if !assertEqual(sum(s1), 7) {
			t.Fail()
		}
		if !assertEqual(sum(s2), "hello world") {
			t.Fail()
		}

	})
}
