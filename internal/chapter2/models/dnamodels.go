package models

import (
	"fmt"

	"github.com/fafadoboy/da-gosb/internal/models"
)

type Nucleotide int

const (
	A Nucleotide = iota
	C
	G
	T
)

// runeToNucleotide converts a rune to a Nucleotide.
func runeToNucleotide(r rune) Nucleotide {
	switch r {
	case 'A':
		return A
	case 'C':
		return C
	case 'G':
		return G
	case 'T':
		return T
	default:
		// Handle invalid nucleotide
		fmt.Printf("Warning: '%c' is not a valid nucleotide.\n", r)
		return A // Return a default value or handle this case as needed
	}
}

type Codon [3]Nucleotide

func (c Codon) value() int {
	return int(c[0]*100 + c[1]*10 + c[2])
}

func (c Codon) Compare(other models.Comparable) int {
	otherCodon, _ := other.(Codon)

	otherSum := otherCodon.value()
	if thisSum := c.value(); thisSum > otherSum {
		return 1
	} else if thisSum < otherSum {
		return -1
	}
	return 0
}

func (c Codon) Less(other Codon) bool {
	return c.value() < other.value()
}

func (c Codon) Equal(other models.Comparable) bool {
	if res := c.Compare(other); res != 0 {
		return false
	}
	return true
}

func NewCodon(r1, r2, r3 rune) Codon {
	return Codon{runeToNucleotide(r1), runeToNucleotide(r2), runeToNucleotide(r3)}
}

type Gene []Codon

func NewGene(s string) Gene {
	gene := make([]Codon, 0)

	for i := 0; i < len(s); i += 3 {
		if i+2 > len(s) {
			break
		}
		gene = append(gene, NewCodon(rune(s[i]), rune(s[i+1]), rune(s[i+2])))
	}

	return gene
}
