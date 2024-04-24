package models

import (
	"fmt"
)

type Edge[T any] struct {
	U, V int
	Meta map[string]T
}

func (e *Edge[T]) Reversed() *Edge[T] {
	return &Edge[T]{U: e.V, V: e.U, Meta: e.Meta}
}

func (e *Edge[T]) ToString() string {
	return fmt.Sprintf("%d -> %d (%v)", e.U, e.V, e.Meta)
}
