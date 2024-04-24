package models

type Node[T any] struct {
	Item     T
	Distance float32
}
