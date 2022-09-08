package goquicksort

import (
	"golang.org/x/exp/constraints"
)

func Quicksort[O constraints.Ordered](slice []O) {
	quicksort(slice, 0, len(slice)-1)
}

func quicksort[O constraints.Ordered](slice []O, li, ri int) {
	if li >= ri {
		return
	}
	pivotIndex := partition(slice, li, ri)
	quicksort(slice, li, pivotIndex-1)
	quicksort(slice, pivotIndex+1, ri)
}

func partition[O constraints.Ordered](slice []O, li, ri int) int {
	startLimit := li // Limit for the right pointer.
	pivotIndex := ri // Pivot is the right edge element of a slice.
	pivot := slice[ri]

	ri-- // Move right pointer to the left to make it point not to the pivot.
	met := false
	for !met {
		for slice[li] < pivot {
			li++
		}
		for ri >= startLimit && slice[ri] > pivot {
			ri--
		}
		if li >= ri {
			met = true
		} else {
			slice[li], slice[ri] = slice[ri], slice[li] // Swap left and right pointers to provide the correct order.
			li++
		}
	}
	slice[li], slice[pivotIndex] = slice[pivotIndex], slice[li]
	return li
}
