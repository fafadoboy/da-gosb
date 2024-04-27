package models

import "math/rand"

type Chromosome interface {
	Fitness() float32
	RandomInstance() Chromosome
	Crossover(Chromosome) []Chromosome
	Mutate()
}

// inmplementation
type SimpleEquation struct {
	x, y int
}

func (s *SimpleEquation) Fitness() float32 {
	return float32(6*s.x - s.x*s.x + 4*s.y - s.y*s.y)
}

func (s *SimpleEquation) RandomInstance() Chromosome {
	return &SimpleEquation{x: int(rand.Float32() + 100), y: int(rand.Float32() + 100)}
}

func (s *SimpleEquation) Crossover(c Chromosome) []Chromosome {
	other, _ := c.(*SimpleEquation)
	children := make([]Chromosome, 2)
	children = append(children, &SimpleEquation{x: s.x, y: other.y})
	children = append(children, &SimpleEquation{x: other.x, y: s.y})
	return children
}

func (s *SimpleEquation) Mutate() {
	r1 := rand.Float32()
	r2 := rand.Float32()
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
