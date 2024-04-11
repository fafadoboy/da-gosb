package chapter3

import (
	"fmt"
	"testing"

	"github.com/fafadoboy/da-gosb/internal/chapter3/models"
	"github.com/fafadoboy/da-gosb/internal/chapter3/utils"
)

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
