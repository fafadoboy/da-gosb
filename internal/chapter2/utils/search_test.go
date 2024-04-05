package utils

import (
	"fmt"
	"testing"

	"github.com/fafadoboy/da-gosb/internal/chapter2/ch2maze"
)

func TestDfsBfs(t *testing.T) {
	algo := AlgoSearch[ch2maze.MazeLocation]{}

	maze := ch2maze.NewMaze(10, 10, 0.2, ch2maze.NewMazeLocation(0, 0), ch2maze.NewMazeLocation(9, 9))
	maze.Print()
	if sol1 := algo.DFS(ch2maze.NewMazeLocation(0, 0), maze.GoalTest, maze.Successors); sol1 != nil {
		path := sol1.ToPath()
		maze.Mark(path)
		maze.Print()
		maze.Clear(path)
	} else {
		fmt.Println("No solution found using depth-first rearch")
	}

	if sol2 := algo.BFS(ch2maze.NewMazeLocation(0, 0), maze.GoalTest, maze.Successors); sol2 != nil {
		path := sol2.ToPath()
		maze.Mark(path)
		maze.Print()
		maze.Clear(path)
	} else {
		fmt.Println("No solution found using breadth-first rearch")
	}
}
