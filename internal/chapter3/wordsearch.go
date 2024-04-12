package chapter3

import (
	"github.com/fafadoboy/da-gosb/internal/chapter3/models"
	"github.com/fafadoboy/da-gosb/internal/utils"
)

type WordSearchConstraint struct {
	words []models.Cell
}

func (c *WordSearchConstraint) Variables() []models.Cell {
	return c.words
}

func (c *WordSearchConstraint) Satisfied(assignment map[string][]models.Cell) bool {
	allLocations := utils.FlatMap[models.Cell](assignment)
	return len(utils.Dedup[models.Cell](allLocations...)) == len(allLocations)
}

func NewWordSearchConstraint(words []models.Cell) *WordSearchConstraint {
	return &WordSearchConstraint{words: words}
}
