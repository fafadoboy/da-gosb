package models

import (
	"github.com/golang-collections/collections/queue"
)

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
