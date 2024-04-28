package chapter5

import (
	"fmt"
	"testing"

	"github.com/fafadoboy/da-gosb/internal/chapter5/models"
)

func TestGeneticAlgo(t *testing.T) {
	initialPopulation := make([]models.Chromosome, 0)
	for i := 0; i < 20; i++ {
		instance := models.SampleEquationInstance()
		fmt.Printf("%d) %+v\n", i+1, instance)
		initialPopulation = append(initialPopulation, instance)
	}

	result := Run(models.Roulette, 100, 0.1, 0.7, 13.0, initialPopulation...)
	fmt.Printf("result is %+v\n", result)

	result2 := Run(models.Tournament, 100, 0.1, 0.7, 13.0, initialPopulation...)
	fmt.Printf("result is %+v\n", result2)
}

func TestGeneticAlgo2(t *testing.T) {
	n := 1000

	initialPopulation := make([]models.Chromosome, 0)
	for i := 0; i < n; i++ {
		instance := models.SendMoreMoneyInstance()
		initialPopulation = append(initialPopulation, instance)
	}

	result := Run(models.Roulette, n, 0.2, 0.7, 1.0, initialPopulation...)
	fmt.Printf("result is \"%+v\"\n", result.ToString())

	result2 := Run(models.Tournament, 1000, 0.2, 0.7, 1.0, initialPopulation...)
	fmt.Printf("result is \"%+v\"\n", result2.ToString())
}
