package chapter3

import (
	"fmt"
	"math"
	"strconv"

	"github.com/fafadoboy/da-gosb/internal/chapter3/models"
)

type QueensConstraint struct {
	columns []models.Index
}

func (c *QueensConstraint) Variables() []models.Index {
	return c.columns
}

func (c *QueensConstraint) Satisfied(assignment map[string]models.Index) bool {
	// q1c = queen 1 column, q1r = queen 1 row
	for q1cKey, q1r := range assignment {
		q1c, _ := strconv.Atoi(q1cKey)
		// q2r = queen in 2 columns
		for q2c := q1c + 1; q2c <= len(c.columns); q2c++ {
			if q2r, ok := assignment[fmt.Sprint(q2c)]; ok {
				if q1r == q2r { // same row?
					return false
				}
				if math.Abs(float64(q1r)-float64(q2r)) == math.Abs(float64(q1c)-float64(q2c)) { // same diagonal?
					return false
				}
			}
		}
	}
	return true // no conflict
}

func NewQueensConstraint(columns []models.Index) *QueensConstraint {
	return &QueensConstraint{columns: columns}
}
