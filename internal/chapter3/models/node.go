package models

import "fmt"

type Node string

func (n Node) Hash() string {
	return fmt.Sprintf("%s", n)
}

func (n Node) Clone() any {
	newNode := n
	return newNode
}
