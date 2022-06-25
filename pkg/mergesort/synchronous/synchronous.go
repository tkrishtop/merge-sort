package synchronous

import (
	"mergesort/pkg/mergesort"
)

func MergeSort(l []int) []int {
	n := len(l)
	if n < 2 {
		return l
	}

	mid := int(n / 2)
	left := MergeSort(l[:mid])
	right := MergeSort(l[mid:])
	return mergesort.Merge(left, right)
}
