package models

import (
	"fmt"
	"math/rand"
)

type SimpleEquation struct {
	x, y int
}

func (s *SimpleEquation) Fitness() float32 {
	return float32(6*s.x - s.x*s.x + 4*s.y - s.y*s.y)
}

func (s *SimpleEquation) Crossover(c Chromosome) []Chromosome {
	other, _ := c.(*SimpleEquation)
	return []Chromosome{&SimpleEquation{x: s.x, y: other.y}, &SimpleEquation{x: other.x, y: s.y}}
}

func (s *SimpleEquation) Mutate() {
	r1, r2 := rand.Float32(), rand.Float32()
	switch {
	case r1 > 0.5 && r2 > 0.5:
		s.x += 1
	case r1 > 0.5 && r2 <= 0.5:
		s.x -= 1
	case r1 <= 0.5 && r2 > 0.5:
		s.y += 1
	case r1 <= 0.5 && r2 <= 0.5:
		s.y -= 1
	}
}

func (s *SimpleEquation) ToString() string {
	return fmt.Sprintf("x: %d, y: %d", s.x, s.y)
}

func SampleEquationInstance() Chromosome {
	return &SimpleEquation{x: rand.Intn(100), y: rand.Intn(100)}
}
