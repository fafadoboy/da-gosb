package models

import "fmt"

type Cell struct {
	row, column int
}

func (c Cell) Hash() string {
	return fmt.Sprintf("R%dC%d", c.row, c.column)
}

func (c Cell) Clone() any {
	return Cell{row: c.row, column: c.column}
}

type Grid [][]string
