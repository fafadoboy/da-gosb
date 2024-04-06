package utils

import (
	"fmt"
	"testing"

	"github.com/fafadoboy/da-gosb/internal/chapter2/ch2maze"
	"github.com/fafadoboy/da-gosb/internal/chapter2/ch2mc"
)

func TestMazeSearch(t *testing.T) {
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

	if sol3 := algo.AStar(ch2maze.NewMazeLocation(0, 0), maze.GoalTest, maze.Successors, maze.Heuristic); sol3 != nil {
		path := sol3.ToPath()
		maze.Mark(path)
		maze.Print()
		maze.Clear(path)
	} else {
		fmt.Println("No solution found using a-star rearch")
	}
}

func TestMSRiddle(t *testing.T) {
	algo := AlgoSearch[ch2mc.MCState]{}
	start := ch2mc.NewMCState(3, 3, true)
	if sol1 := algo.BFS(start, start.GoalTest, start.Successors); sol1 != nil {
		fmt.Println("SOLL\n=====")
		for _, state := range sol1.ToPath() {
			state.Print()
		}
	} else {
		fmt.Println("No solution found using breadth-first rearch")
	}

	if sol2 := algo.DFS(start, start.GoalTest, start.Successors); sol2 != nil {
		fmt.Println("SOLL\n=====")
		for _, state := range sol2.ToPath() {
			state.Print()
		}
	} else {
		fmt.Println("No solution found using breadth-first rearch")
	}
}
