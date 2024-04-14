package models

type Constraint[V, D any] interface {
	Variables() []V
	Satisfied(assignment map[string]D) bool
}
