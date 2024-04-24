package utils

import (
	"container/heap"

	"github.com/fafadoboy/da-gosb/internal/chapter4/models"
	"github.com/samber/lo"
)

type AlgoGraph[V, T any] struct {
}

// MST stands for minimal spaning tree (Jerkin's algorithm)
func (*AlgoGraph[V, T]) MST(wg *models.Graph[V, T], start int, converter func(T) float32) (path []*models.Edge[T]) {
	if start > (wg.VertexCount()-1) || start < 0 {
		return nil
	}

	pq := models.PriorityQueue[*models.Edge[T]]{}
	heap.Init(&pq)

	visited := make([]bool, wg.VertexCount())

	visit := func(index int) {
		visited[index] = true
		for _, edge := range wg.EdgesForIndex(index) {
			if !visited[edge.V] {
				pq.Push(&models.Node[*models.Edge[T]]{Item: edge, Distance: converter(edge.Meta["weight"])})
			}
		}
	}

	visit(start) // the first vertex is where everything begins

	for pq.Len() > 0 {
		edge := pq.Pop().(*models.Node[*models.Edge[T]]).Item
		if visited[edge.V] {
			continue // don't ever revisit
		}

		// this is the current smallest, so add it to solution
		path = append(path, edge)
		visit(edge.V) // visit where this connects
	}
	return path
}

func (*AlgoGraph[V, T]) Dijkstra(wg *models.Graph[V, T], root V, converter func(T) float32) (distances []*float32, path map[int]*models.Edge[T]) {
	distances = make([]*float32, wg.VertexCount())
	path = make(map[int]*models.Edge[T])

	first := wg.IndexOf(root)
	distances[first] = lo.ToPtr[float32](0.0)
	pq := models.PriorityQueue[int]{}
	heap.Init(&pq)

	pq.Push(&models.Node[int]{Item: first, Distance: 0.0})

	for pq.Len() > 0 {

		// explore the next closest vertex
		n := pq.Pop().(*models.Node[int])
		u := n.Item

		distU := distances[u] // should already have seed it

		// look at every edge/vertex from the vertex in question
		for _, we := range wg.EdgesForIndex(u) {
			// extract the wight

			// the old distance to this vertex
			distV := distances[we.V]
			// no old distance or found shorter path
			weight := converter(we.Meta["weight"])
			sum := weight + *distU
			if distV == nil || *distV > sum {
				// update distance to this vertex
				distances[we.V] = lo.ToPtr(sum)
				// update the edge on the shortest path to this vertex
				path[we.V] = we
				// explore it soon
				pq.Push(&models.Node[int]{Item: we.V, Distance: sum})
			}
		}

	}

	return distances, path
}
