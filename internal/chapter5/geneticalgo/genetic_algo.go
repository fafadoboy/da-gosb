package geneticalgo

import (
	"math/rand"

	"github.com/fafadoboy/da-gosb/internal/chapter5/models"
	"github.com/fafadoboy/da-gosb/internal/utils"
	"github.com/samber/lo"
)

func reproduceAndReplace(crossoverChance float32, population ...models.Chromosome) []models.Chromosome {
	newPopulation := make([]models.Chromosome, 0)

	for len(newPopulation) < len(population) {
		parents, _ := utils.RandomChoices(population, lo.Map(population, func(item models.Chromosome, _ int) float32 {
			return item.Fitness()
		}), 2)

		if rand.Float32() < crossoverChance {
			newPopulation = append(newPopulation, parents[0].Crossover(parents[1])...)
		} else {
			newPopulation = append(newPopulation, parents...)
		}
	}

	// if we had an odd number, we’ll have 1 extra, so we remove it
	if len(newPopulation) > len(population) {
		newPopulation = newPopulation[:len(newPopulation)-1]
	}

	return newPopulation
}

func mutate(mutationChance float32, population ...models.Chromosome) {
	for _, individual := range population {
		if rand.Float32() < mutationChance {
			individual.Mutate()
		}
	}
}

func Run(generationsCount int, mutationsChance, crossoverChance, threshold float32, population ...models.Chromosome) models.Chromosome {

	best := lo.MaxBy(population, func(a, b models.Chromosome) bool { return a.Fitness() > b.Fitness() })
	for i := 0; i < generationsCount; i++ {
		if best.Fitness() >= threshold {
			return best
		}

		reproduceAndReplace(crossoverChance, population...)
		mutate(mutationsChance, population...)
		highest := lo.MaxBy(population, func(a, b models.Chromosome) bool { return a.Fitness() > b.Fitness() })
		if highest.Fitness() > best.Fitness() {
			best = highest
		}
	}
	return best
}
