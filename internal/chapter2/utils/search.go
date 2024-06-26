package utils

import (
	"container/heap"

	"github.com/fafadoboy/da-gosb/internal/chapter2/models"
)

type Hashable interface {
	Hash() string
}
type AlgoSearch[T Hashable] struct {
}

func (a *AlgoSearch[T]) DFS(initial T, GoalTest func(T) bool, successors func(T) []T) *models.Node[T] {
	frontier := models.Stack[*models.Node[T]]{}
	frontier.Init()

	explored := models.Set[string]{}
	explored.Init()

	frontier.Push(&models.Node[T]{State: initial})
	explored.Insert(initial.Hash())

	for !frontier.Empty() {
		currentNode := frontier.Pop()
		currentState := currentNode.State
		if GoalTest(currentState) {
			return currentNode
		}
		for _, child := range successors(currentState) {
			if explored.Has(child.Hash()) {
				continue
			}
			explored.Insert(child.Hash())
			frontier.Push(&models.Node[T]{State: child, Parent: currentNode})
		}
	}

	return nil
}

func (a *AlgoSearch[T]) BFS(initial T, GoalTest func(T) bool, successors func(T) []T) *models.Node[T] {
	frontier := models.Queue[*models.Node[T]]{}
	frontier.Init()

	explored := models.Set[string]{}
	explored.Init()

	frontier.Push(&models.Node[T]{State: initial})
	explored.Insert(initial.Hash())

	for !frontier.Empty() {
		currentNode := frontier.Pop()
		currentState := currentNode.State
		if GoalTest(currentState) {
			return currentNode
		}
		for _, child := range successors(currentState) {
			if explored.Has(child.Hash()) {
				continue
			}
			explored.Insert(child.Hash())
			frontier.Push(&models.Node[T]{State: child, Parent: currentNode})
		}
	}

	return nil
}

func (a *AlgoSearch[T]) AStar(initial T, GoalTest func(T) bool, successors func(T) []T, heuristic func(T) float32) *models.Node[T] {
	frontier := make(models.PriorityQueue[T], 0)
	heap.Init(&frontier)

	explored := make(map[string]float32, 0)

	heap.Push(&frontier, &models.Node[T]{State: initial, Cost: 0.0, Heuristic: heuristic(initial)})
	explored[initial.Hash()] = 0.0

	for frontier.Len() > 0 {
		currentNode := heap.Pop(&frontier).(*models.Node[T])
		currentState := currentNode.State
		if GoalTest(currentState) {
			return currentNode
		}
		for _, child := range successors(currentState) {
			newCost := currentNode.Cost + 1
			if val, ok := explored[child.Hash()]; !ok || val > newCost {
				explored[child.Hash()] = newCost
				heap.Push(&frontier, &models.Node[T]{State: child, Parent: currentNode, Cost: newCost, Heuristic: heuristic(child)})
			}
		}
	}

	return nil
}
