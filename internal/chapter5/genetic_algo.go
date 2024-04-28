package chapter5

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/fafadoboy/da-gosb/internal/chapter5/models"
	"github.com/fafadoboy/da-gosb/internal/utils"
	"github.com/samber/lo"
)

func reproduceAndReplace(selection models.SelectionType, crossoverChance float32, population ...models.Chromosome) []models.Chromosome {
	newPopulation := make([]models.Chromosome, 0)

	for len(newPopulation) < len(population) {

		var parents []models.Chromosome
		switch selection {
		case models.Roulette:
			weights := lo.Map(population, func(item models.Chromosome, _ int) float32 {
				return item.Fitness()
			})
			parents, _ = utils.RandomChoices(population, weights, 2)
		case models.Tournament:
			sample := utils.Sample(population, int(len(population)/2))
			sort.Slice(sample, func(i, j int) bool {
				return sample[i].Fitness() < sample[j].Fitness()
			})
			parents = sample[len(sample)-2:]
		}

		if rand.Float32() < crossoverChance {
			crossover := parents[0].Crossover(parents[1])
			// fmt.Printf("%v->%v, %v->%v\n", parents[0].ToString(), crossover[0].ToString(), parents[1].ToString(), crossover[1].ToString())

			newPopulation = append(newPopulation, crossover...)
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

func Run(selection models.SelectionType, maxGenerations int, mutationsChance, crossoverChance, threshold float32, population ...models.Chromosome) models.Chromosome {

	best := lo.MaxBy(population, func(a, b models.Chromosome) bool { return a.Fitness() > b.Fitness() })
	for i := 0; i < maxGenerations; i++ {
		if best.Fitness() >= threshold {
			return best
		}

		sum := lo.ReduceRight[models.Chromosome](population, func(agg float32, item models.Chromosome, _ int) float32 {
			return agg + item.Fitness()
		}, 0.0)
		fmt.Printf("Generation %d Best %s Avg %f\n", i, best.ToString(), sum/float32(len(population)))
		population := reproduceAndReplace(selection, crossoverChance, population...)
		mutate(mutationsChance, population...)
		highest := lo.MaxBy(population, func(a, b models.Chromosome) bool { return a.Fitness() > b.Fitness() })
		if highest.Fitness() > best.Fitness() {
			best = highest
		}
	}
	return best
}
