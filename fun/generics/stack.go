package generics

import (
	"errors"
	"fmt"
)

// a simple generic stack implementation
type Stack[T any] struct {
	stackData *[]T
}

// returns an empty stack
func NewStack[T any]() *Stack[T] {
	s := []T{}
	return &Stack[T]{
		stackData: &s,
	}
}

// Append value at the end of stacka , ss is the underlying stack slice
func (s *Stack[T]) Push(val T) {
	if s.stackData == nil {
		s.stackData = &[]T{}
	}
	ss := *s.stackData
	ss = append(ss, val)
	*s.stackData = ss

}

// removes value from the top of the stack
func (s *Stack[T]) Pop() (error, T) {
	var pop_val T
	stack_len := s.Stacklen()
	if s.Stacklen() == 0 {
		return errors.New("can't pop from an empty stack"), pop_val
	}
	ss := *s.stackData
	pop_val = ss[stack_len-1]
	*s.stackData = ss[0 : stack_len-1]
	return nil, pop_val

}

// returns the number of elements in the stack
func (s *Stack[T]) Stacklen() int {

	return len(*s.stackData)
}

// Prints the stack on screen
func (s *Stack[T]) PrinTStack() {
	ss := *s.stackData
	fmt.Println(ss)
}
