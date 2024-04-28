package utils

import (
	"math/rand"

	"github.com/fafadoboy/da-gosb/internal/models"
	"github.com/samber/lo"
)

// Dedup removes duplicate items from a slice based on their hash values.
// It takes a variable number of items of any type that implements the models.Hashable interface,
// which requires a Hash method that returns a unique string identifier for each item.
//
// Parameters:
//   - items ...T: A variadic slice of items where T is constrained to types that implement the models.Hashable interface.
//
// Returns:
//   - deduped []T: A slice containing only the unique items, preserving the order they appear in the original input.
//
// Example Usage:
//
//	DedupedSlice := Dedup(hashableItems...)
//
// This function operates in linear time relative to the number of items, O(n),
// with respect to both time and space complexity, where n is the number of items in the input slice.
func Dedup[T models.Hashable](items ...T) (deduped []T) {
	allKeys := make(map[string]bool, 0)

	for _, cell := range items {
		if _, ok := allKeys[cell.Hash()]; !ok {
			allKeys[cell.Hash()] = true
			deduped = append(deduped, cell)
		}
	}
	return
}

// FlatMap concatenates the slices of items stored in a map into a single slice.
// It takes a map with string keys and values that are slices of any type (T).
//
// Parameters:
//   - itemsMap map[string][]T: A map where each key is a string and the value is a slice of type T.
//
// Returns:
//   - items []T: A single slice containing all the elements from all slices in the map, concatenated in the order they appear.
//
// Example Usage:
//
//	combinedItems := FlatMap(mapOfSlices)
//
// This function is useful for combining data from a grouped or categorized collection into a single list.
// It operates in linear time relative to the total number of individual items across all slices, O(n),
// where n is the cumulative count of all items in all the slices within the map.
func FlatMap[T any](itemsMap map[string][]T) (items []T) {
	for _, values := range itemsMap {
		items = append(items, values...)
	}
	return
}

// RandomChoices selects n elements from the provided slice using the given weights.
// If weights are nil, it assumes uniform distribution.
func RandomChoices[T any](elements []T, weights []float32, n int) ([]T, error) {
	// Initialize the result slice
	result := make([]T, n)

	// Select randomly if no weights are provided
	if weights == nil {
		for i := 0; i < n; i++ {
			result[i] = elements[rand.Intn(len(elements))]
		}
		return result, nil
	}

	// Otherwise
	// normalize weights
	min := lo.Min(weights)
	if min < 0 {
		min *= -1
	}
	weights = lo.Map(weights, func(item float32, _ int) float32 { return item + min })

	// use the cumulative weight method if weights are provided
	cumulativeWeights := make([]float32, len(weights))
	cumulativeWeights[0] = weights[0]
	for i := 1; i < len(weights); i++ {
		cumulativeWeights[i] = cumulativeWeights[i-1] + weights[i]
	}
	totalWeight := cumulativeWeights[len(cumulativeWeights)-1]

	for i := 0; i < n; i++ {
		r := rand.Float32() * totalWeight
		for j, w := range cumulativeWeights {
			if r <= w {
				result[i] = elements[j]
				break
			}
		}
	}

	return result, nil
}

func Sample[T any](elements []T, k int) []T {
	if k > len(elements) {
		k = len(elements)
	}

	lenElements := len(elements)
	result := make([]T, 0)
	setOfIndexes := make(map[int]bool)

	for k > 0 {
		r := rand.Intn(lenElements)
		if _, ok := setOfIndexes[r]; !ok {
			result = append(result, elements[r])
			setOfIndexes[r] = true
			k--
		}
	}
	return result
}
