package dnasearch

import (
	"fmt"
	"sort"
	"testing"
)

func TestLinearContains(t *testing.T) {
	geneStr := "AGAGCAACCCTGAGATAGTCTGGCTATTTCGCAACTACGCGGTCGAAGGCAATGCAGGGCCTTGCAGTAATAAGGGCAGCCTCTCGGACGAATTAAAACC"

	myGene := NewGene(geneStr)
	sort.SliceStable(myGene, func(i, j int) bool {
		return myGene[i].Compare(myGene[i]) > -1
	})

	acg := NewCodon('A', 'C', 'G')
	gat := NewCodon('G', 'A', 'T')
	fmt.Println(linearContains(myGene, acg))
	fmt.Println(linearContains(myGene, gat))
}

func TestBinaryContains(t *testing.T) {
	geneStr := "AGAGCAACCCTGAGATAGTCTGGCTATTTCGCAACTACGCGGTCGAAGGCAATGCAGGGCCTTGCAGTAATAAGGGCAGCCTCTCGGACGAATTAAAACC"

	myGene := NewGene(geneStr)
	sort.SliceStable(myGene, func(i, j int) bool {
		return myGene[i].Less(myGene[j])
	})

	acg := NewCodon('A', 'C', 'G')
	gat := NewCodon('G', 'A', 'T')
	fmt.Println(binaryContains(myGene, acg))
	fmt.Println(binaryContains(myGene, gat))
}
