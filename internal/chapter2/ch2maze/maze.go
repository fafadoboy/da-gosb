package ch2maze

import (
	"fmt"
	"math/rand"
	"time"
)

type Cell string

const (
	EMPTY   Cell = "_"
	BLOCKED Cell = "X"
	START   Cell = "S"
	GOAL    Cell = "G"
	PATH    Cell = "*"
)

func (ml MazeLocation) Equal(other MazeLocation) bool {
	return ml.row == other.row && ml.column == other.column
}

type Maze struct {
	rows, columns int
	start, goal   MazeLocation
	grid          [][]Cell
}

func (m *Maze) setStartGoal() {
	m.grid[m.start.column][m.start.row] = START
	m.grid[m.goal.column][m.goal.row] = GOAL
}

func (m *Maze) Successors(ml MazeLocation) (locations []MazeLocation) {
	if ml.row+1 < m.rows && m.grid[ml.row+1][ml.column] != BLOCKED {
		locations = append(locations, MazeLocation{ml.row + 1, ml.column})
	}
	if ml.row-1 >= 0 && m.grid[ml.row-1][ml.column] != BLOCKED {
		locations = append(locations, MazeLocation{ml.row - 1, ml.column})
	}
	if ml.column+1 < m.columns && m.grid[ml.row][ml.column+1] != BLOCKED {
		locations = append(locations, MazeLocation{ml.row, ml.column + 1})
	}
	if ml.column-1 >= 0 && m.grid[ml.row][ml.column-1] != BLOCKED {
		locations = append(locations, MazeLocation{ml.row, ml.column - 1})
	}
	return locations
}

func (m *Maze) GoalTest(ml MazeLocation) bool {
	return m.goal.Equal(ml)
}

func (m *Maze) Mark(path []MazeLocation) {
	for _, ml := range path {
		m.grid[ml.row][ml.column] = PATH
	}
	m.setStartGoal()
}

func (m *Maze) Clear(path []MazeLocation) {
	for _, ml := range path {
		m.grid[ml.row][ml.column] = EMPTY
	}
	m.setStartGoal()
}

func (m *Maze) Print() {
	fmt.Println("\nMAZE\n=====")
	for i := 0; i < len(m.grid); i++ {
		for j := 0; j < len(m.grid[i]); j++ {
			fmt.Print(m.grid[i][j])
		}
		fmt.Println()
	}
}

func NewMaze(rows, columns int, sparseness float32, start, goal MazeLocation) *Maze {
	rand.Seed(time.Now().UnixNano())

	grid := make([][]Cell, 0)
	for i := 0; i < columns; i++ {
		row := make([]Cell, 0)
		for j := 0; j < rows; j++ {
			if rand.Float32() < sparseness {
				row = append(row, BLOCKED)
			} else {
				row = append(row, EMPTY)
			}
		}
		grid = append(grid, row)
	}

	maze := &Maze{rows: rows, columns: columns, start: start, goal: goal, grid: grid}
	maze.setStartGoal()
	return maze
}
