/*
Implement top-down merge sort.
*/
package main

import "fmt"

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

func mergeSort(l []int) []int {
	n := len(l)
	if n < 2 {
		return l
	}

	mid := int(n / 2)
	left := mergeSort(l[:mid])
	right := mergeSort(l[mid:])
	return merge(left, right)
}

func main() {
	l1 := []int{1, 3, 5, 4, 6, 8, 9, 10, 2, 7}
	fmt.Println(mergeSort(l1))
	l2 := []int{}
	fmt.Println(mergeSort(l2))
}
