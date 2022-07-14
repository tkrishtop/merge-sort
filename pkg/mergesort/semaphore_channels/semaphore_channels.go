package semaphore_channels

import (
	"mergesort/pkg/mergesort"
)

type MixedSorter struct {
	sem chan int
}

func NewMixedSorter(semCapacity int) *MixedSorter {
	return &MixedSorter{
		sem: make(chan int, semCapacity),
	}
}

func (m *MixedSorter) MergeSort(l []int) []int {
	n := len(l)
	if n < 2 {
		return l
	}
	mid := int(n / 2)
	results := make(chan []int)

	select {
	case m.sem <- 1:
		go func() {
			results <- m.MergeSort(l[:mid])
			<-m.sem
		}()
	default:
		go func() {
			results <- m.MergeSort(l[:mid])
		}()
	}

	select {
	case m.sem <- 1:
		go func() {
			results <- m.MergeSort(l[mid:])
			<-m.sem
		}()
	default:
		go func() {
			results <- m.MergeSort(l[mid:])
		}()
	}

	right := <-results
	left := <-results
	return mergesort.Merge(left, right)
}
