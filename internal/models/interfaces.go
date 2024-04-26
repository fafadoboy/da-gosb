package models

type Hashable interface {
	Hash() string
}

type Clonable interface {
	Clone() any
}

type Comparable interface {
	Compare(Comparable) int
	Equal(Comparable) bool
}

type HashableAndClonable interface {
	Hashable
	Clonable
}

type AenimaItem interface {
	Hashable
	Clonable
	Comparable
}

type Comperator[T any] interface {
	Compare(T, T) int
}
