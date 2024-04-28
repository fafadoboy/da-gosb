package models

import (
	"fmt"
	"math"
	"strings"

	"github.com/fafadoboy/da-gosb/internal/utils"
	"github.com/samber/lo"
)

type SendMoreMoney struct {
	letters []string
}

func (s *SendMoreMoney) Fitness() float32 {
	_s := lo.IndexOf(s.letters, "S")
	_e := lo.IndexOf(s.letters, "E")
	_n := lo.IndexOf(s.letters, "N")
	_d := lo.IndexOf(s.letters, "D")
	_m := lo.IndexOf(s.letters, "M")
	_o := lo.IndexOf(s.letters, "O")
	_r := lo.IndexOf(s.letters, "R")
	_y := lo.IndexOf(s.letters, "Y")
	send := _s*1000 + _e*100 + _n*10 + _d
	more := _m*1000 + _o*100 + _r*10 + _e
	money := _m*10000 + _o*1000 + _n*100 + _e*10 + _y
	difference := math.Abs(float64(money - (send + more)))
	return float32(1 / (difference + 1))
}

func (s *SendMoreMoney) Crossover(c Chromosome) []Chromosome {
	other, _ := c.(*SendMoreMoney)
	child1 := &SendMoreMoney{letters: make([]string, len(s.letters))}
	copy(child1.letters, s.letters)

	child2 := &SendMoreMoney{letters: make([]string, len(other.letters))}
	copy(child2.letters, other.letters)

	indexes := make([]int, len(s.letters))
	for n := range indexes {
		indexes[n] = n
	}
	indexes = utils.Sample(indexes, 2)

	idx1, idx2 := indexes[0], indexes[1]
	l1, l2 := child1.letters[idx1], child2.letters[idx2]
	child1.letters[lo.IndexOf(child1.letters, l2)], child1.letters[idx2] = child1.letters[idx2], l2
	child2.letters[lo.IndexOf(child2.letters, l1)], child2.letters[idx1] = child2.letters[idx1], l1

	return []Chromosome{child1, child2}
}

func (s *SendMoreMoney) Mutate() {
	indexes := make([]int, len(s.letters))
	for n := range indexes {
		indexes[n] = n
	}
	indexes = utils.Sample(indexes, 2)
	idx1, idx2 := indexes[0], indexes[1]

	s.letters[idx1], s.letters[idx2] = s.letters[idx2], s.letters[idx1]
}

func (s *SendMoreMoney) ToString() string {
	return fmt.Sprintf("%s (%f)", strings.Join(s.letters, ""), s.Fitness())
}

func SendMoreMoneyInstance() Chromosome {
	letters := []string{"S", "E", "N", "D", "M", "O", "R", "Y", " ", " "}
	return &SendMoreMoney{letters: utils.Sample(letters, len(letters))}
}
