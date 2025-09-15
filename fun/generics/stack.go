package generics

import (
	"errors"
	"fmt"
)

// a simple generic stack implementation
type Stack[T any] struct {
	stackData []T
}

// returns an empty stack
func NewStack[T any]() *Stack[T] {
	s := []T{}
	return &Stack[T]{
		stackData: s,
	}
}

// Append value at the end of stack , ss is the underlying stack slice
func (s *Stack[T]) Push(val T) {
	if s.stackData == nil {
		slc := []T{}
		slc = append(slc, val)
		s.stackData = slc
	}
	ss := s.stackData
	ss = append(ss, val)
	s.stackData = ss

}

// removes value from the top of the stack
func (s *Stack[T]) Pop() (error, T) {

	var popval T
	if s.stackData == nil {
		return cantPopEmptyStack, popval
	}
	ss := s.stackData
	popval = ss[len(ss)-1]
	new_ss := ss[0 : len(ss)-1]
	s.stackData = new_ss
	return nil, popval

}

// returns the number of elements in the stack
func (s *Stack[T]) Stacklen() int {
	return len(s.stackData)
}

// Prints the stack on screen
func (s *Stack[T]) PrinTStack() {
	fmt.Println(s.stackData)
}

var (
	cantPopEmptyStack = errors.New("can't pop out of an empty stack")
)
