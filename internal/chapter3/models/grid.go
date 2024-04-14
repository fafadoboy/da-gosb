package models

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/samber/lo"
)

func rangeGenerator(start, end int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := start; i < end; i++ {
			ch <- i
		}
	}()
	return ch
}

func randomUppercaseLetter() string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" // equivalent to string.ascii_uppercase
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(len(letters)) // generate a random index
	return string(letters[index]) // return the letter at the random index
}

type GridLocation struct {
	row, column int
}

func (c *GridLocation) Get() (int, int) {
	return c.row, c.column
}

func (c GridLocation) Hash() string {
	return fmt.Sprintf("R%dC%d", c.row, c.column)
}

func (c GridLocation) Clone() any {
	return GridLocation{row: c.row, column: c.column}
}

type ListGL []GridLocation

func (l ListGL) Clone() any {
	clone := make(ListGL, 0)
	for _, location := range l {
		clone = append(clone, location.Clone().(GridLocation))
	}
	return clone
}

type Grid [][]string

func (grid Grid) Print() {
	for _, row := range grid {
		for _, col := range row {
			fmt.Print(col + " ")
		}
		fmt.Println() // New line after each row
	}
}

func (g Grid) GenerateDomain(word string) []ListGL {
	domain := make([]ListGL, 0)
	height := len(g)
	width := len(g[0])
	length := len(word)

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			columnsGenerator := func(yield func(int)) {
				for i := col; i < col+length; i++ {
					yield(i)
				}
			}

			rowsGenerator := func(yield func(int)) {
				for i := row; i < row+length; i++ {
					yield(i)
				}
			}

			if col+length <= width {
				// left to right
				listToAppend := make([]GridLocation, 0)
				for c := range lo.Generator(0, columnsGenerator) {
					listToAppend = append(listToAppend, GridLocation{row: row, column: c})
				}
				domain = append(domain, listToAppend)
				// diagonal towards bottom right
				if row+length <= height {
					listToAppend := make([]GridLocation, 0)
					for r := range lo.Generator(0, rowsGenerator) {
						listToAppend = append(listToAppend, GridLocation{row: r, column: col + (r - row)})
					}
					domain = append(domain, listToAppend)
				}
			}

			if row+length <= height {
				listToAppend := make([]GridLocation, 0)
				for r := range lo.Generator(0, rowsGenerator) {
					listToAppend = append(listToAppend, GridLocation{row: r, column: col})
				}
				domain = append(domain, listToAppend)
				if col-length >= 0 {
					listToAppend := make([]GridLocation, 0)
					for r := range lo.Generator(0, rowsGenerator) {
						listToAppend = append(listToAppend, GridLocation{row: r, column: col - (r - row)})
					}
					domain = append(domain, listToAppend)
				}
			}

		}
	}
	return domain
}

func NewGrid(rows, columns int) Grid {
	grid := make([][]string, rows)
	for i := range grid {
		grid[i] = make([]string, columns)
		for j := range grid[i] {
			grid[i][j] = randomUppercaseLetter()
		}
	}
	return grid
}
