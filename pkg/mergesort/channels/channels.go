package channels

import (
	"mergesort/pkg/mergesort"
)

func MergeSort(l []int) []int {
	n := len(l)
	if n < 2 {
		return l
	}

	ch_left, ch_right := make(chan []int), make(chan []int)

	mid := int(n / 2)

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
