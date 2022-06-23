/*
Implement top-down merge sort using goroutine to fasten up the execution.
*/

package main

import (
	"sync"
)

func merge(l1 []int, l2 []int) []int {
	var output []int
	p1, p2 := 0, 0
	for p1 < len(l1) && p2 < len(l2) {
		if l1[p1] <= l2[p2] {
			output = append(output, l1[p1])
			p1 += 1
		} else {
			output = append(output, l2[p2])
			p2 += 1
		}
	}

	if p1 < len(l1) {
		output = append(output, l1[p1:]...)
	} else {
		output = append(output, l2[p2:]...)
	}

	return output
}

func MergeSortWaitGroups(l []int) []int {
	n := len(l)
	if n < 2 {
		return l
	}

	// create a new group
	// that must wait for two goroutines
	wg := new(sync.WaitGroup)
	wg.Add(2)

	mid := int(n / 2)

	var left []int
	var right []int

	go func() {
		defer wg.Done()
		left = MergeSortWaitGroups(l[:mid])
	}()

	go func() {
		defer wg.Done()
		right = MergeSortWaitGroups(l[mid:])
	}()

	wg.Wait()
	return merge(left, right)
}
