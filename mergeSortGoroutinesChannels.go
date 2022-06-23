/*
Implement top-down merge sort using goroutine to fasten up the execution.
*/

package main

func mergeChannels(l1 []int, l2 []int) []int {
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

func MergeSortChannels(l []int) []int {
	n := len(l)
	if n < 2 {
		return l
	}

	ch_left, ch_right := make(chan []int), make(chan []int)

	mid := int(n / 2)

	go func() {
		ch_left <- MergeSortChannels(l[:mid])
		close(ch_left)
	}()

	go func() {
		ch_right <- MergeSortChannels(l[mid:])
		close(ch_right)
	}()

	for {
		left, ok_left := <-ch_left
		right, ok_right := <-ch_right

		if ok_left && ok_right {
			return mergeChannels(left, right)
		}
	}
}
