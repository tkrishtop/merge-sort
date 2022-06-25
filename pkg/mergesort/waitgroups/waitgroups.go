package waitgroups

import (
	"mergesort/pkg/mergesort"
	"sync"
)

func MergeSort(l []int) []int {
	n := len(l)
	if n < 2 {
		return l
	}

	// create a new group
	// that would wait for two goroutines
	wg := new(sync.WaitGroup)
	wg.Add(2)

	mid := int(n / 2)

	var left []int
	var right []int

	go func() {
		defer wg.Done()
		left = MergeSort(l[:mid])
	}()

	go func() {
		defer wg.Done()
		right = MergeSort(l[mid:])
	}()

	wg.Wait()
	return mergesort.Merge(left, right)
}
