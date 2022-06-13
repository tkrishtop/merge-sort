/*
Implement top-down merge sort.
*/
package main

func mergeNoGoroutines(l1 []int, l2 []int) []int {
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

func MergeSortNoGoroutines(l []int) []int {
	n := len(l)
	if n < 2 {
		return l
	}

	mid := int(n / 2)
	left := MergeSortNoGoroutines(l[:mid])
	right := MergeSortNoGoroutines(l[mid:])
	return mergeNoGoroutines(left, right)
}
