package utils

import "github.com/fafadoboy/da-gosb/internal/models"

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
//   DedupedSlice := Dedup(hashableItems...)
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
//   combinedItems := FlatMap(mapOfSlices)
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
