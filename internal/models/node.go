package models

import (
	"sync/atomic"

	"github.com/samber/lo"
)

// TODO: implement Node struct to wrap a state, states parent, and state's properties (will be a map[string]V)
// 		 the in addition the Node will have

var gId int64

type NodesCompareFunc[T Comparable] func(this, other *Node[T]) int

type Node[T Comparable] struct {
	id         int64
	item       T
	parent     *Node[T]
	properties map[string]interface{}
	cmpFunc    func(this, other *Node[T]) int
}

func (n *Node[T]) ID() int64 {
	return n.id
}

func (n *Node[T]) Item() T {
	return n.item
}

func (n *Node[T]) AddProperty(propName string, propVal interface{}) interface{} {
	n.properties[propName] = propVal
	return propVal
}

func (n *Node[T]) AddComparison(cmp NodesCompareFunc[T]) {
	n.cmpFunc = cmp
}

func (n *Node[T]) Property(key string) interface{} {
	if p, ok := n.properties[key]; ok {
		return p
	}
	return nil
}

func (n *Node[T]) Less(other *Node[T]) bool {
	return n.cmpFunc(n, other) <= -1
}

func (n *Node[T]) Compare(other *Node[T]) int {
	return n.cmpFunc(n, other)
}

func (n *Node[T]) ToPath() []T {
	path := make([]T, 0)
	node := n

	for node.parent != nil {
		node = node.parent
		path = append(path, node.item)
	}

	path = append(path, node.item)
	return lo.Reverse[T](path)
}

func NewNode[T Comparable](item T, parent *Node[T]) *Node[T] {
	id := atomic.AddInt64(&gId, 1)

	return &Node[T]{item: item,
		id:         id,
		parent:     parent,
		cmpFunc:    func(this, other *Node[T]) int { return 1 },
		properties: make(map[string]interface{})}
}
