package dnasearch

func linearContains(gene Gene, key Codon) bool {
	for _, codon := range gene {
		if codon.Equal(key) {
			return true
		}
	}
	return false
}

func binaryContains(gene Gene, key Codon) bool {
	low := 0
	high := len(gene) - 1

	for low <= high {
		mid := (low + high + 1) >> 1
		if res := gene[mid].Compare(key); res < 0 {
			low = mid + 1
		} else if res > 0 {
			high = mid - 1
		} else {
			return true
		}
	}
	return false
}
