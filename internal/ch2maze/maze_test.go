package ch2maze

import "testing"

func TestNewMaze(t *testing.T) {
	maze := NewMaze(10, 10, 0.3, MazeLocation{0, 0}, MazeLocation{9, 9})
	maze.Print()
}
