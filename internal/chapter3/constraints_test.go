package chapter3

import (
	"fmt"
	"testing"

	"github.com/fafadoboy/da-gosb/internal/chapter3/models"
	"github.com/fafadoboy/da-gosb/internal/chapter3/utils"
)

func TestMapColoring(t *testing.T) {
	variables := []models.Node{
		models.WESTAUSTRALIA, models.NORTHERNTERRITORY, models.SOUTHAUSTRALIA, models.QUEENSLAND, models.NEWSOUTHWALES,
		models.VICTORIA, models.TASMANIA}
	domains := make(map[string][]models.Node, 0)

	for _, variable := range variables {
		domains[variable.Hash()] = append(domains[variable.Hash()], "red", "green", "blue")
	}

	cps, error := utils.NewCPS[models.Node, models.Node](variables, domains)
	if error != nil {
		fmt.Println("Failed to create CPS")
	}
	cps.AddConstraint(NewMapColoringConstraint(models.WESTAUSTRALIA, models.NORTHERNTERRITORY))
	cps.AddConstraint(NewMapColoringConstraint(models.WESTAUSTRALIA, models.SOUTHAUSTRALIA))

	cps.AddConstraint(NewMapColoringConstraint(models.SOUTHAUSTRALIA, models.NORTHERNTERRITORY))

	cps.AddConstraint(NewMapColoringConstraint(models.QUEENSLAND, models.NORTHERNTERRITORY))
	cps.AddConstraint(NewMapColoringConstraint(models.QUEENSLAND, models.SOUTHAUSTRALIA))
	cps.AddConstraint(NewMapColoringConstraint(models.QUEENSLAND, models.NEWSOUTHWALES))

	cps.AddConstraint(NewMapColoringConstraint(models.NEWSOUTHWALES, models.SOUTHAUSTRALIA))

	cps.AddConstraint(NewMapColoringConstraint(models.VICTORIA, models.SOUTHAUSTRALIA))
	cps.AddConstraint(NewMapColoringConstraint(models.VICTORIA, models.NEWSOUTHWALES))
	cps.AddConstraint(NewMapColoringConstraint(models.VICTORIA, models.TASMANIA))

	if sol := cps.BacktrackingSearch(make(map[string]models.Node)); sol == nil {
		fmt.Printf("No solution found!")
	} else {
		for k, v := range sol {
			fmt.Printf("key: %s, val: %v\n", k, v)
		}
	}

}

func TestQueens(t *testing.T) {
	columns := []models.Index{1, 2, 3, 4, 5, 6, 7, 8}
	rows := make(map[string][]models.Index, 0)

	for _, column := range columns {
		rows[column.Hash()] = []models.Index{1, 2, 3, 4, 5, 6, 7, 8}
	}

	cps, error := utils.NewCPS[models.Index, models.Index](columns, rows)
	if error != nil {
		fmt.Println("Failed to create CPS")
	}
	cps.AddConstraint(NewQueensConstraint(columns))
	if sol := cps.BacktrackingSearch(make(map[string]models.Index)); sol == nil {
		fmt.Printf("No solution found!")
	} else {
		for k, v := range sol {
			fmt.Printf("(%s,%v)\n", k, v)
		}
	}
}

func TestWordSearch(t *testing.T) {
	grid := models.NewGrid(9, 9)
	// grid.Print()
	variables := []models.Node{"MATTHEW", "JOE", "MARY", "SARAH", "SALLY"}
	domains := make(map[string][]models.ListGL, 0)

	for _, variable := range variables {
		domains[variable.Hash()] = grid.GenerateDomain(variable.Hash())
	}

	csp, error := utils.NewCPS(variables, domains)
	if error != nil {
		fmt.Println("Failed to create CPS")
	}
	csp.AddConstraint(NewWordSearchConstraint(variables...))

	if sol := csp.BacktrackingSearch(make(map[string]models.ListGL)); sol == nil {
		fmt.Printf("No solution found!")
	} else {
		for word, locations := range sol {
			for i := 0; i < len(word); i++ {
				row, column := locations[i].Get()
				grid[row][column] = word[i : i+1]
			}
		}
		grid.Print()
	}
}
