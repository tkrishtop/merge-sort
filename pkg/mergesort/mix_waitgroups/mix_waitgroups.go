package mix_waitgroups

import (
	"mergesort/pkg/mergesort"
	"sync"
)

var Max = 4

func MergeSort(l []int) []int {
	n := len(l)
	if n < 2 {
		return l
	}
	mid := int(n / 2)

	sem := make(chan int, Max)
	var left []int
	var right []int

	wg := sync.WaitGroup{}
	wg.Add(2)

	select {
	case sem <- 1:
		go func() {
			left = MergeSort(l[:mid])
			<-sem
			wg.Done()
		}()
	default:
		left = MergeSort(l[:mid])
		wg.Done()
	}

	select {
	case sem <- 1:
		go func() {
			right = MergeSort(l[mid:])
			<-sem
			wg.Done()
		}()
	default:
		right = MergeSort(l[:mid])
		wg.Done()
	}

	wg.Wait()
	return mergesort.Merge(left, right)
}
