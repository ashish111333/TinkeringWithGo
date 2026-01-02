package exercises

import "sync/atomic"

// lock free concurrent stack , so you will be using atomics

// each node stores stack data and a down_ptr pointing to node below
// stack (this one ) grows in bottom to up manner
type node[T interface{}] struct {
	data     T
	down_ptr *node[T]
}

type Stack[T interface{}] struct {
	head atomic.Pointer[node[T]]
}

func (s *Stack[T]) Push(x T) {
	newNode := &node[T]{
		data: x,
	}
	if !s.head.CompareAndSwap(nil, newNode) {
		currHead := s.head.Load()
		newNode.down_ptr = currHead
		s.head.Store(newNode)

	}
}

// change head to second last element from top
func (s *Stack[T]) Pop() (res T) {
	if s.head.Load() == nil {
		return
	}
	for {
		currHead := s.head.Load()
		prevNode := currHead.down_ptr
		if s.head.CompareAndSwap(currHead, prevNode) {
			return currHead.data
		}

	}

}

func (s *Stack[T]) Head() T {
	headNode := s.head.Load()
	return headNode.data
}

func NewLfStack[T any]() *Stack[T] {
	return new(Stack[T])

}
