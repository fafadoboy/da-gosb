package utils

import (
	"container/heap"

	"github.com/fafadoboy/da-gosb/internal/chapter4/models"
)

type AlgoGraph[V, T any] struct {
}

// MST stands for minimal spaning tree (Jerkin's algorithm)
func (a *AlgoGraph[V, T]) MST(wg *models.Graph[V, T], start int) (path []*models.Edge[float32]) {
	if start > (wg.VertexCount()-1) || start < 0 {
		return nil
	}

	pq := models.NewWeightPriorityQueue(func(f1, f2 float32) bool { return f1 < f2 })
	heap.Init(pq)

	visited := make([]bool, wg.VertexCount())

	visit := func(index int) {
		visited[index] = true
		for _, edge := range wg.EdgesForIndex(index) {
			if !visited[edge.V] {
				pq.Push(edge)
			}
		}
	}

	visit(start) // the first vertex is where everything begins

	for pq.Len() > 0 {
		edge := pq.Pop().(*models.Edge[float32])
		if visited[edge.V] {
			continue // don't ever revisit
		}

		// this is the current smallest, so add it to solution
		path = append(path, edge)
		visit(edge.V) // visit where this connects
	}
	return path
}

func (AlgoGraph[V, T]) Dijkstra(wg *models.Graph[V, T], root V) (distances []float32, path map[int]*models.Edge[T]) {
	distances = make([]float32, 0)

	first := wg.IndexOf(root)
	distances[first] = 0.0
	pq := models.NewWeightPriorityQueue(func(f1, f2 float32) bool { return f1 < f2 })
	heap.Init(pq)

	// TBD
	// pq.Push()

	return distances, path
}
