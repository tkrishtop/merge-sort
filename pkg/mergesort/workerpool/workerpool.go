package workerpool

import (
	"mergesort/pkg/mergesort"
)

func MergeSort(l []int) []int {
	n := len(l)
	if n < 2 {
		return l
	}
	mid := int(n / 2)

	sem := make(chan int, 90)
	if len(sem) < cap(sem) {
		sem <- 1
		go func() {
			l = asyncJob(mid, l)
			<-sem
		}()
	} else {
		l = syncJob(mid, l)
		// time.Sleep(10 * time.Nanosecond)
	}

	// waiting for all async jobs to finish
	for len(sem) != 0 {
	}

	close(sem)
	return l
}

func asyncJob(mid int, l []int) []int {
	ch_left, ch_right := make(chan []int), make(chan []int)

	go func() {
		ch_left <- MergeSort(l[:mid])
		close(ch_left)
	}()

	go func() {
		ch_right <- MergeSort(l[mid:])
		close(ch_right)
	}()

	for {
		left, ok_left := <-ch_left
		right, ok_right := <-ch_right

		if ok_left && ok_right {
			return mergesort.Merge(left, right)
		}
	}
}

func syncJob(mid int, l []int) []int {
	left := MergeSort(l[:mid])
	right := MergeSort(l[mid:])
	return mergesort.Merge(left, right)
}
