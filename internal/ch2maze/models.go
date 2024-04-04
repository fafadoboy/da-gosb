package ch2maze

import (
	"fmt"
	"math/rand"
)

type Cell string

const (
	EMPTY   Cell = "_"
	BLOCKED Cell = "X"
	START   Cell = "S"
	GOAL    Cell = "G"
	PATH    Cell = "*"
)

type MazeLocation struct {
	Row, Column int
}

func (ml MazeLocation) Equal(other MazeLocation) bool {
	return ml.Row == other.Row && ml.Column == other.Column
}

type Maze struct {
	rows, columns int
	Start, Goal   MazeLocation
	grid          [][]Cell
}

func (m *Maze) GoalTest(ml MazeLocation) bool {
	return m.Goal.Equal(ml)
}

func (m *Maze) Print() {
	for i := 0; i < len(m.grid); i++ {
		for j := 0; j < len(m.grid[i]); j++ {
			fmt.Print(m.grid[i][j])
		}
		fmt.Println()
	}
}

func NewMaze(rows, columns int, sparseness float32, start, goal MazeLocation) *Maze {
	grid := make([][]Cell, 0)
	for i := 0; i < columns; i++ {
		row := make([]Cell, 0)
		for j := 0; j < rows; j++ {
			if rand.Float64() < float64(sparseness) {
				row = append(row, BLOCKED)
			} else {
				row = append(row, EMPTY)
			}
		}
		grid = append(grid, row)
	}

	grid[start.Column][start.Row] = START
	grid[goal.Column][goal.Row] = GOAL

	return &Maze{rows: rows, columns: columns, Start: start, Goal: goal, grid: grid}
}
