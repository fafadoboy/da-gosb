package models

import "fmt"

type Edge struct {
	u, v int
}

func (e *Edge) Reversed() Edge {
	return Edge{u: e.v, v: e.u}
}

func (e *Edge) Hash() string {
	return fmt.Sprintf("%d -> %d", e.u, e.v)
}

type WeightedEdge struct {
	Edge
	weight float32
}

func (w *WeightedEdge) Compare(other *WeightedEdge) int {

	if w.weight > other.weight {
		return 1
	}

	if w.weight < other.weight {
		return -1
	}
	return 0
}

func (w *WeightedEdge) Reversed() WeightedEdge {
	return WeightedEdge{Edge: w.Edge.Reversed(), weight: w.weight}
}

func (w *WeightedEdge) ToString() string {
	return fmt.Sprintf("%d %f > %d", w.u, w.weight, w.v)
}
