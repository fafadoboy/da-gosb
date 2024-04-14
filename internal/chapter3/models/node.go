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

type Index int

func (i Index) Hash() string {
	return fmt.Sprintf("%d", i)
}

func (i Index) Clone() any {
	newNode := i
	return newNode
}
