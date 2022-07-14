package semaphore_waitgroups

import (
	"mergesort/pkg/mergesort"
	"mergesort/pkg/mergesort/synchronous"
	"sync"
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

	var left []int
	var right []int

	wg := sync.WaitGroup{}
	wg.Add(2)

	select {
	case m.sem <- 1:
		go func() {
			left = m.MergeSort(l[:mid])
			<-m.sem
			wg.Done()
		}()
	default:
		left = synchronous.MergeSort(l[:mid])
		wg.Done()
	}

	select {
	case m.sem <- 1:
		go func() {
			right = m.MergeSort(l[mid:])
			<-m.sem
			wg.Done()
		}()
	default:
		right = synchronous.MergeSort(l[mid:])
		wg.Done()
	}

	wg.Wait()
	return mergesort.Merge(left, right)
}
