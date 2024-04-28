package models

type Chromosome interface {
	Fitness() float32
	Crossover(Chromosome) []Chromosome
	Mutate()
	ToString() string
}
