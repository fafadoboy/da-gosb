package ch2maze

import (
	"fmt"
	"hash/fnv"
)

type MazeLocation struct {
	row, column int
}

func (ml MazeLocation) Hash() string {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("C%dR%d", ml.column, ml.row)))
	// If your T type is more complex or doesn't implement fmt.Stringer,
	// you might need a more sophisticated approach to generate a hash.
	return fmt.Sprintf("%x", h.Sum32())
}

func NewMazeLocation(row, column int) MazeLocation {
	return MazeLocation{row: row, column: column}
}
