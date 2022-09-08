package goquicksort_test

import (
	"testing"

	goquicksort "github.com/Asideron/go-quicksort"
)

func TestSort(t *testing.T) {
	slice := []int{0, 1, 3, 8, 5, 2, 3}
	sortedSlice := []int{0, 1, 2, 3, 3, 5, 8}

	goquicksort.Quicksort(slice)

	if len(slice) != len(sortedSlice) {
		t.Errorf("Incorrect slice len.\ngot: %v\nwanted: %v", len(slice), len(sortedSlice))
	}
	for i := range slice {
		if slice[i] != sortedSlice[i] {
			t.Errorf("Slices are not equal.\ngot: %v\nwanted: %v", slice, sortedSlice)
		}
	}
}
