package models

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	pq := make(PriorityQueue[string], 0)
	heap.Init(&pq)

	heap.Push(&pq, &Node[string]{State: "A", Cost: 2, Heuristic: 1})
	heap.Push(&pq, &Node[string]{State: "B", Cost: 1, Heuristic: 2})

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node[string])
		fmt.Println(node.State)
	}

}
