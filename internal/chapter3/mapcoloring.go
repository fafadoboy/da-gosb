package chapter3

import (
	"github.com/fafadoboy/da-gosb/internal/chapter3/models"
)

type MapColoringConstraint struct {
	region1 models.Node
	region2 models.Node
}

func (c *MapColoringConstraint) Variables() []models.Node {
	return []models.Node{c.region1, c.region2}
}

func (c *MapColoringConstraint) Satisfied(assignment map[string]models.Node) bool {
	// Extract colors based on node hash (or direct string conversion in this case)
	color1, ok1 := assignment[c.region1.Hash()]
	color2, ok2 := assignment[c.region2.Hash()]

	// Check if both regions are assigned and have different colors
	if ok1 && ok2 && color1 == color2 {
		return false
	}
	return true
}

func NewMapColoringConstraint(region1 models.Node, region2 models.Node) *MapColoringConstraint {
	return &MapColoringConstraint{
		region1: region1,
		region2: region2,
	}
}
