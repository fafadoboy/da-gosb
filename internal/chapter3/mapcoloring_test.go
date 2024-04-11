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
	// variables := []MapColoringConstraint{}
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
