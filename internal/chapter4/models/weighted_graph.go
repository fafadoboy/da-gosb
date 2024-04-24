package models

// import (
// 	"fmt"
// 	"strings"
// )

// type VertexWeight[V any] struct {
// 	vertex V
// 	weight float32
// }

// type WeightedGraph[V any] struct {
// 	Graph[V]
// }

// func (w *WeightedGraph[V]) NeighborsForIndexWithWeights(index int) (distances []VertexWeight[V]) {
// 	for _, edge := range w.EdgesForIndex(index) {
// 		if weight, ok := edge.Meta["weight"]; ok {
// 			distances = append(distances, VertexWeight[V]{vertex: w.VertexAt(index), weight: weight})
// 		}
// 	}
// 	return distances
// }

// func (w *WeightedGraph[V]) ToString() string {
// 	var sb strings.Builder
// 	for i := 0; i < w.VertexCount(); i++ {
// 		sb.WriteString(fmt.Sprintf("%v -> %v\n", w.VertexAt(i), w.NeighborsForIndexWithWeights(i)))
// 	}
// 	return sb.String()
// }

// func NewWightedGraph[V comparable](vertices ...V) *WeightedGraph[V] {
// 	graph := NewGraph(vertices...)
// 	graph.comparator = func(v1, v2 V) int { return 0 }
// 	return &WeightedGraph[V]{Graph: *graph}
// }
