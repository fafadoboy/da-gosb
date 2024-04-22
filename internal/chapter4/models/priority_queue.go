package models

type WeightPriorityQueue[T comparable] struct {
	edges  []*Edge[T]
	isLess func(T, T) bool
}

func (pq WeightPriorityQueue[T]) Len() int { return len(pq.edges) }

func (pq WeightPriorityQueue[T]) Less(i, j int) bool {
	// Use Edge's Less method for comparison.
	return pq.isLess(pq.edges[i].Meta["weight"], pq.edges[j].Meta["weight"])
}

func (pq WeightPriorityQueue[T]) Swap(i, j int) {
	pq.edges[i], pq.edges[j] = pq.edges[j], pq.edges[i]
}

func (pq *WeightPriorityQueue[T]) Push(x interface{}) {
	item := x.(*Edge[T])
	pq.edges = append(pq.edges, item)
}

func (pq *WeightPriorityQueue[T]) Pop() interface{} {
	item := pq.edges[pq.Len()-1]
	pq.edges = pq.edges[:pq.Len()-1]
	return item
}

func NewWeightPriorityQueue[T comparable](isLess func(T, T) bool) *WeightPriorityQueue[T] {
	return &WeightPriorityQueue[T]{
		edges:  make([]*Edge[T], 0),
		isLess: isLess,
	}
}
