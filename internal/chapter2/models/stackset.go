package models

import (
	"github.com/golang-collections/collections/set"
	"github.com/golang-collections/collections/stack"
)

type Stack[T any] struct {
	container *stack.Stack
}

func (s *Stack[T]) Empty() bool {
	return s.container.Len() == 0
}

func (s *Stack[T]) Push(item T) {
	s.container.Push(item)
}

func (s *Stack[T]) Pop() T {
	return s.container.Pop().(T)
}

func (s *Stack[T]) Init() {
	s.container = stack.New()
}

/*
	SET
*/

type Set[T any] struct {
	container *set.Set
}

func (s *Set[T]) Has(el T) bool {
	return s.container.Has(el)
}

func (s *Set[T]) Insert(el T) {
	s.container.Insert(el)
}

func (s *Set[T]) Init() {
	s.container = set.New()
}
