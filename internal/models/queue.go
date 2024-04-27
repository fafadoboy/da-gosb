package models

import (
	"container/heap"

	"github.com/golang-collections/collections/queue"
)

// type priorityQueue[T any] []*Node[T]
type priorityQueue[T Comparable] []T

func (pq priorityQueue[T]) Len() int { return len(pq) }

func (pq priorityQueue[T]) Less(i, j int) bool {
	// Use Node's Less method for comparison.
	return pq[i].Compare(pq[j]) < 0
}

func (pq priorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue[T]) Push(x interface{}) {
	item := x.(T)
	*pq = append(*pq, item)
}

func (pq *priorityQueue[T]) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type Heap[T Comparable] struct {
	frontier *priorityQueue[T]
}

func (h *Heap[T]) Len() int {
	return h.frontier.Len()
}

func (h *Heap[T]) Push(item T) {
	heap.Push(h.frontier, item)
}

func (h *Heap[T]) Pop() T {
	return heap.Pop(h.frontier).(T)
}

func NewHeap[T Comparable]() *Heap[T] {
	frontier := make(priorityQueue[T], 0)
	heap.Init(&frontier)
	return &Heap[T]{frontier: &frontier}
}

type Queue[T any] struct {
	container *queue.Queue
}

func (q *Queue[T]) Empty() bool {
	return q.container.Len() == 0
}

func (q *Queue[T]) Push(item T) {
	q.container.Enqueue(item)
}

func (q *Queue[T]) Pop() T {
	return q.container.Dequeue().(T)
}

func (q *Queue[T]) Init() {
	q.container = queue.New()
}
