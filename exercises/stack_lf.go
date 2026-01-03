package exercises

import (
	"sync"
	"sync/atomic"
)

// lock free concurrent stack , so you will be using atomics

// each node stores stack data and a down_ptr pointing to node below
// stack (this one ) grows in bottom to up manner
type node[T comparable] struct {
	data     T
	down_ptr *node[T]
}

type Stack[T comparable] struct {
	head atomic.Pointer[node[T]]
}

func (s *Stack[T]) Push(x T) {
	newNode := &node[T]{
		data: x,
	}
	if !s.head.CompareAndSwap(nil, newNode) {

		for {
			currHead := s.head.Load()
			newNode.down_ptr = currHead
			if s.head.CompareAndSwap(currHead, newNode) {
				return
			}
		}

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

// check if an element exists or not(not safe for concurrent use)
func (s *Stack[T]) exists(element T) bool {
	currNode := s.head.Load()
	if currNode == nil {
		return false
	}
	var found bool

	for {
		data := currNode.data
		if data == element {
			found = true
			return found
		}
		if currNode.down_ptr == nil {
			return found

		}
		currNode = currNode.down_ptr
	}

}

func (s *Stack[T]) Head() T {
	headNode := s.head.Load()
	return headNode.data
}

func NewLfStack[T comparable]() *Stack[T] {
	return new(Stack[T])

}

type StackMx[T comparable] struct {
	head *node[T]
	mx   sync.Mutex
}

func (s *StackMx[T]) Push(element T) {
	s.mx.Lock()
	defer s.mx.Unlock()
	newNode := &node[T]{
		data: element,
	}
	if s.head == nil {
		s.head = newNode
		return

	}
	currHead := s.head
	s.head = newNode
	newNode.down_ptr = currHead

}
func (s *StackMx[T]) Pop() (res T) {
	s.mx.Lock()
	defer s.mx.Unlock()
	if s.head == nil {
		return
	}
	res = s.head.data
	preNode := s.head.down_ptr
	s.head = preNode
	return res

}

func NewStackMx[T comparable]() *StackMx[T] {
	return new(StackMx[T])
}
