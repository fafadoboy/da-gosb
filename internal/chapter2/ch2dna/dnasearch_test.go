package dnasearch

import (
	"fmt"
	"sort"
	"testing"

	"github.com/fafadoboy/da-gosb/internal/utils"
)

const geneStr = "AGAGCAACCCTGAGATAGTCTGGCTATTTCGCAACTACGCGGTCGAAGGCAATGCAGGGCCTTGCAGTAATAAGGGCAGCCTCTCGGACGAATTAAAACC"

func newGene() Gene {
	myGene := NewGene(geneStr)
	sort.SliceStable(myGene, func(i, j int) bool {
		return myGene[i].Less(myGene[j])
	})
	return myGene
}

func TestLinearContains(t *testing.T) {
	myGene := newGene()

	acg := NewCodon('A', 'C', 'G')
	gat := NewCodon('G', 'A', 'T')
	fmt.Println(utils.LinearContains(acg, myGene...))
	fmt.Println(utils.LinearContains(gat, myGene...))
}

func TestBinaryContains(t *testing.T) {
	myGene := newGene()
	acg := NewCodon('A', 'C', 'G')
	gat := NewCodon('G', 'A', 'T')
	fmt.Println(utils.BinaryContains(acg, myGene...))
	fmt.Println(utils.BinaryContains(gat, myGene...))
}
