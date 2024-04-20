package models

import (
	"fmt"
	"strings"
)

type VertexWeight[V any] struct {
	first  V
	second float32
}

type WeightedGraph[V comparable] struct {
	Graph[V]
	edges [][]WeightedEdge
}

func (w *WeightedGraph[V]) AddEdge(edge WeightedEdge) {
	w.edges[edge.u] = append(w.edges[edge.u], edge)
	w.edges[edge.v] = append(w.edges[edge.v], edge)
}

func (w *WeightedGraph[V]) EdgesForIndex(index int) []WeightedEdge {
	return w.edges[index]
}

func (w *WeightedGraph[V]) AddEdgeByIndices(u, v int, weight float32) {
	edge := WeightedEdge{Edge: Edge{u: u, v: v}, weight: weight}
	w.AddEdge(edge)
}

func (w *WeightedGraph[V]) AddEdgeByVertices(first, second V, weight float32) {
	w.AddEdgeByIndices(w.IndexOf(first), w.IndexOf(second), weight)
}

func (w *WeightedGraph[V]) NeighborsForIndexWithWeights(index int) (distances []VertexWeight[V]) {
	for _, edge := range w.EdgesForIndex(index) {
		distances = append(distances, VertexWeight[V]{first: w.VertexAt(index), second: edge.weight})
	}
	return distances
}

func (w *WeightedGraph[V]) ToString() string {
	var sb strings.Builder
	for i := 0; i < w.VertexCount(); i++ {
		sb.WriteString(fmt.Sprintf("%v -> %v\n", w.VertexAt(i), w.NeighborsForIndexWithWeights(i)))
	}
	return sb.String()
}

func NewWightedGraph[V comparable](vertices ...V) *WeightedGraph[V] {
	return &WeightedGraph[V]{
		Graph: *NewGraph(vertices...),
		edges: func() (r [][]WeightedEdge) {
			for i := 0; i < len(vertices); i++ {
				r = append(r, make([]WeightedEdge, 0))
			}
			return r
		}()}
}
