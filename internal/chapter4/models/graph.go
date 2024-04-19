package models

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

type Edge struct {
	u, v int
}

func (e *Edge) Reversed() Edge {
	return Edge{u: e.v, v: e.u}
}

func (e *Edge) Hash() string {
	return fmt.Sprintf("%d -> %d", e.u, e.v)
}

type Graph[V comparable] struct {
	vertices []V
	edges    [][]Edge
}

func (g *Graph[V]) VertexCount() int {
	return len(g.vertices)
}

func (g *Graph[V]) EdgeCount() int {
	sum := 0
	for _, e := range g.edges {
		sum += len(e)
	}
	return sum
}

func (g *Graph[V]) AddVertex(vertex V) int {
	g.vertices = append(g.vertices, vertex)
	g.edges = append(g.edges, make([]Edge, 0))
	return len(g.vertices) - 1
}

func (g *Graph[V]) AddEdge(edge Edge) {
	g.edges[edge.u] = append(g.edges[edge.u], edge)
	g.edges[edge.v] = append(g.edges[edge.v], edge.Reversed())
}

func (g *Graph[V]) AddEdgeByIndeces(u, v int) {
	g.AddEdge(Edge{u, v})
}

func (g *Graph[V]) AddEdgeByVertices(first, second V) {
	_, u, _ := lo.FindIndexOf[V](g.vertices, func(item V) bool { return item == first })
	_, v, _ := lo.FindIndexOf[V](g.vertices, func(item V) bool { return item == second })
	g.AddEdgeByIndeces(u, v)
}

func (g *Graph[V]) VertexAt(index int) V {
	return g.vertices[index]
}

func (g *Graph[V]) IndexOf(item V) int {
	index := -1
	for n, vertex := range g.vertices {
		if vertex == item {
			index = n
			break
		}
	}
	return index
}

func (g *Graph[V]) NeighborsForIndex(index int) (vertexes []V) {
	for _, e := range g.edges[index] {
		vertexes = append(vertexes, g.VertexAt(e.v))
	}
	return vertexes
}

func (g *Graph[V]) NeighborsForVertex(vertex V) []V {
	return g.NeighborsForIndex(g.IndexOf(vertex))
}

func (g *Graph[V]) EdgesForIndex(index int) []Edge {
	return g.edges[index]
}

func (g *Graph[V]) EdgesForVertex(vertex V) []Edge {
	return g.EdgesForIndex(g.IndexOf(vertex))
}

func (g *Graph[V]) ToString() string {
	var sb strings.Builder
	for i := 0; i < g.VertexCount(); i++ {
		sb.WriteString(fmt.Sprintf("%v -> %v\n", g.VertexAt(i), g.NeighborsForIndex(i)))
	}
	return sb.String()
}

func NewGraph[V comparable](vertices ...V) *Graph[V] {
	return &Graph[V]{
		vertices: vertices,
		edges: func() (r [][]Edge) {
			for i := 0; i < len(vertices); i++ {
				r = append(r, make([]Edge, 0))
			}
			return r
		}()}
}
