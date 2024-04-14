package models

type Hashable interface {
	Hash() string
}

type Clonable interface {
	Clone() any
}

type HashableAndClonable interface {
	Hashable
	Clonable
}
