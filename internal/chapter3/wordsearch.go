package chapter3

import (
	"github.com/fafadoboy/da-gosb/internal/chapter3/models"
	"github.com/fafadoboy/da-gosb/internal/utils"
)

type WordSearchConstraint struct {
	words []models.Node
}

func (c *WordSearchConstraint) Variables() []models.Node {
	return c.words
}

func (c *WordSearchConstraint) Satisfied(assignment map[string]models.ListGL) bool {
	allLocations := make(models.ListGL, 0)
	for _, values := range assignment {
		allLocations = append(allLocations, values...)
	}
	return len(utils.Dedup(allLocations...)) == len(allLocations)
}

func NewWordSearchConstraint(words ...models.Node) *WordSearchConstraint {
	return &WordSearchConstraint{words: words}
}
