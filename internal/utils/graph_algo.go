package utils

import (
	"github.com/fafadoboy/da-gosb/internal/models"
	"github.com/samber/lo"
)

func astarCmpFunc[T any](this, other *models.Node[T]) int {
	tsum := lo.ReduceRight[string, float32]([]string{"cost", "heuristic"}, func(agg float32, property string, index int) float32 {
		if val, ok := this.Property(property).(float32); ok {
			agg += val
		}
		return agg
	}, 0.0)

	osum := lo.ReduceRight[string, float32]([]string{"cost", "heuristic"}, func(agg float32, property string, index int) float32 {
		if val, ok := other.Property(property).(float32); ok {
			agg += val
		}
		return agg
	}, 0.0)

	return CompareFloat32(tsum, osum)
}

// DFS performs a depth-first search over a space defined by a start node, goal condition, and successors function.
// The function returns a pointer to a Node of type T if the goal is reached, or nil if no solution is found.
func DFS[T models.Hashable](initial T, GoalTest func(T) bool, successors func(T) []T) *models.Node[T] {
	// Initialize a heap to manage the frontier; using a heap for DFS is unconventional,
	// typically a stack (or a simple list with LIFO order) is used.
	heap := models.NewHeap[*models.Node[T]]()

	// Set to keep track of explored states to avoid revisiting them.
	explored := models.Set[string]{}
	explored.Init()

	// Start by pushing the initial state onto the heap and marking it as explored.
	heap.Push(models.NewNode(initial, nil))
	explored.Insert(initial.Hash())

	// Continue searching while there are nodes to explore.
	for heap.Len() > 0 {
		// Pop the node from the heap (this simulates stack behavior if the heap is actually a priority queue).
		currentNode := heap.Pop()
		currentState := currentNode.Item()

		// Check if the current node meets the goal criteria.
		if GoalTest(currentState) {
			return currentNode
		}

		// Expand the current node to find its successors.
		for _, child := range successors(currentState) {
			// Check if the child has already been explored.
			if explored.Has(child.Hash()) {
				continue
			}

			// Mark the child as explored and add it to the heap.
			explored.Insert(child.Hash())
			heap.Push(models.NewNode(child, currentNode))
		}
	}

	// Return nil if no goal state is found after exploring all possibilities.
	return nil
}

// BFS performs a breadth-first search over a space defined by a start node, goal condition, and successors function.
// The function returns a pointer to a Node of type T if the goal is reached, or nil if no solution is found.
func BFS[T models.Hashable](initial T, GoalTest func(T) bool, successors func(T) []T) *models.Node[T] {
	// Initialize a queue to manage the frontier; this is typical for BFS which explores the nearest nodes first.
	frontier := models.Queue[*models.Node[T]]{}
	frontier.Init()

	// Set to keep track of explored states to avoid cycles and revisiting nodes.
	explored := models.Set[string]{}
	explored.Init()

	// Start by enqueuing the initial state and marking it as explored.
	frontier.Push(models.NewNode(initial, nil))
	explored.Insert(initial.Hash())

	// Continue searching while there are nodes to explore in the queue.
	for !frontier.Empty() {
		// Dequeue the first node in the queue.
		currentNode := frontier.Pop()
		currentState := currentNode.Item()

		// Check if the current node satisfies the goal criteria.
		if GoalTest(currentState) {
			return currentNode // Return the current node as the search goal has been achieved.
		}

		// Generate successors of the current node and process each one.
		for _, child := range successors(currentState) {
			// Check if the child node has already been explored.
			if explored.Has(child.Hash()) {
				continue // Skip processing this child if it has been explored.
			}

			// Mark the child as explored and add it to the queue for further exploration.
			explored.Insert(child.Hash())
			frontier.Push(models.NewNode(child, currentNode))
		}
	}

	// Return nil if no goal state is found after exploring all possible paths.
	return nil
}

// AStar performs the A* search algorithm over a graph defined by initial nodes, goal conditions, and successor functions.
// It also requires a heuristic function to estimate the cost from any node to the goal.
// The function returns a pointer to a Node of type T if the goal is reached, or nil if no solution is found.
func AStar[T models.Hashable](initial T, GoalTest func(T) bool, successors func(T) []T, heuristic func(T) float32) *models.Node[T] {
	// Initialize a priority queue (heap) to manage the frontier nodes with the lowest cost first.
	heap := models.NewHeap[*models.Node[T]]()

	// A map to keep track of the lowest cost to reach all explored nodes, avoiding redundant paths.
	explored := make(map[string]float32, 0)

	// Prepare the initial node with necessary properties.
	node := models.NewNode(initial, nil)
	node.AddComparison(astarCmpFunc)                  // Assuming this sets the node comparison function for the heap.
	node.AddProperty("cost", 0.0)                     // Start with a cost of 0 for the initial node.
	node.AddProperty("heuristic", heuristic(initial)) // Estimate the cost from the initial node to the goal.

	// Add the initial node to the heap and mark it as explored with a cost of 0.
	heap.Push(node)
	explored[initial.Hash()] = 0.0

	// Continue the search as long as there are nodes in the frontier.
	for heap.Len() > 0 {
		// Pop the node with the lowest cost from the heap.
		currentNode := heap.Pop()
		currentState := currentNode.Item()

		// Check if the current node meets the goal criteria.
		if GoalTest(currentState) {
			return currentNode // Return the current node as the search goal has been achieved.
		}

		// Expand the current node to find its successors.
		for _, child := range successors(currentState) {
			// Calculate the new cost to reach the child node.
			newCost := currentNode.Property("cost").(float32) + 1 // Assume each step costs 1.

			// Check if the child node has been explored or if a cheaper path to it is found.
			if val, ok := explored[child.Hash()]; !ok || val > newCost {
				// Update the cost for reaching the child node.
				explored[child.Hash()] = newCost
				// Create a new node for the child with updated cost and heuristic.
				node := models.NewNode(child, currentNode)
				node.AddProperty("cost", newCost)
				node.AddProperty("heuristic", heuristic(child))
				node.AddComparison(astarCmpFunc) // Set the comparison function for the new node.
				// Add the new node to the heap.
				heap.Push(node)
			}
		}
	}

	// Return nil if no goal state is found after exploring all possibilities.
	return nil
}
