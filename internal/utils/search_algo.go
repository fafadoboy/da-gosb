package utils

import "github.com/fafadoboy/da-gosb/internal/models"

// LinearContains implements the linear search algorithm
func LinearContains[T models.Comparable](key T, sequence ...T) bool {
	for _, item := range sequence {
		if ok := key.Equal(item); ok {
			return true
		}
	}
	return false
}

// BinaryContains implements the binary search algorithm
func BinaryContains[T models.Comparable](key T, sequence ...T) bool {
	low := 0
	high := len(sequence) - 1

	for low <= high {
		mid := (low + high + 1) >> 1
		if res := sequence[mid].Compare(key); res < 0 {
			low = mid + 1
		} else if res > 0 {
			high = mid - 1
		} else {
			return true
		}
	}
	return false
}
