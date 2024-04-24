package models

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

type Graph[V, T any] struct {
	vertices   []V
	edges      [][]*Edge[T]
	comparator func(V, V) int
}

func (g *Graph[V, T]) VertexCount() int {
	return len(g.vertices)
}

func (g *Graph[V, T]) EdgeCount() int {
	sum := 0
	for _, e := range g.edges {
		sum += len(e)
	}
	return sum
}

func (g *Graph[V, T]) AddVertex(vertex V) int {
	g.vertices = append(g.vertices, vertex)
	g.edges = append(g.edges, make([]*Edge[T], 0))
	return len(g.vertices) - 1
}

func (g *Graph[V, T]) AddEdge(edge *Edge[T]) {
	g.edges[edge.U] = append(g.edges[edge.U], edge)
	g.edges[edge.V] = append(g.edges[edge.V], edge.Reversed())
}

func (g *Graph[V, T]) AddEdgeByIndices(u, v int, metadata map[string]T) {
	g.AddEdge(&Edge[T]{U: u, V: v, Meta: metadata})
}

func (g *Graph[V, T]) AddEdgeByVertices(first, second V, metadata map[string]T) {
	_, u, _ := lo.FindIndexOf[V](g.vertices, func(item V) bool { return g.comparator(item, first) == 0 })
	_, v, _ := lo.FindIndexOf[V](g.vertices, func(item V) bool { return g.comparator(item, second) == 0 })
	g.AddEdgeByIndices(u, v, metadata)
}

func (g *Graph[V, T]) VertexAt(index int) V {
	return g.vertices[index]
}

func (g *Graph[V, T]) IndexOf(item V) int {
	index := -1
	for n, vertex := range g.vertices {
		if g.comparator(vertex, item) == 0 {
			index = n
			break
		}
	}
	return index
}

func (g *Graph[V, T]) NeighborsForIndex(index int) (vertexes []V) {
	for _, e := range g.edges[index] {
		vertexes = append(vertexes, g.VertexAt(e.V))
	}
	return vertexes
}

func (g *Graph[V, T]) NeighborsForVertex(vertex V) []V {
	return g.NeighborsForIndex(g.IndexOf(vertex))
}

func (g *Graph[V, T]) EdgesForIndex(index int) []*Edge[T] {
	return g.edges[index]
}

func (g *Graph[V, T]) EdgesForVertex(vertex V) []*Edge[T] {
	return g.EdgesForIndex(g.IndexOf(vertex))
}

func (g *Graph[V, T]) ToString() string {
	var sb strings.Builder
	for i := 0; i < g.VertexCount(); i++ {
		sb.WriteString(fmt.Sprintf("%v -> %v\n", g.VertexAt(i), g.NeighborsForIndex(i)))
	}
	return sb.String()
}

func NewGraph[V, T comparable](vertices ...V) *Graph[V, T] {
	return &Graph[V, T]{
		comparator: func(v1, v2 V) int {
			if v1 == v2 {
				return 0
			}
			return 1
		},
		vertices: vertices,
		edges: func() (r [][]*Edge[T]) {
			for i := 0; i < len(vertices); i++ {
				r = append(r, make([]*Edge[T], 0))
			}
			return r
		}()}
}
