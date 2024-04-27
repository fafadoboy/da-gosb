package models

import (
	"github.com/golang-collections/collections/set"
)

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
