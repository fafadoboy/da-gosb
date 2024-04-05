package models

type Node[T any] struct {
	State     T
	Parent    *Node[T]
	Cost      float32
	Heuristic float32
}

func (n *Node[T]) Less(other *Node[T]) bool {
	return (n.Cost + n.Heuristic) < (other.Cost + other.Heuristic)
}

func (n *Node[T]) ToPath() (path []T) {
	node := n

	path = append(path, node.State)
	for node.Parent != nil {
		node = node.Parent
		path = append(path, node.State)
	}
	return path[:len(path)-1]
}
