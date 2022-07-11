package main

import (
	"fmt"
	"math/rand"
	"mergesort/pkg/mergesort/mix_channel"
	"time"
)

func generateRandomArray(N int) []int {
	rand.Seed(time.Now().UnixNano())
	lst := rand.Perm(N)

	return lst
}

func main() {
	var N = 1000
	var lst = generateRandomArray(N)
	fmt.Println("Testing on array of size", N)

	// for workerpool_waitgroups.Max = 1; workerpool_waitgroups.Max < 20; workerpool_waitgroups.Max++ {
	// 	start := time.Now()
	// 	l := workerpool_waitgroups.MergeSort(lst)
	// 	elapsed := time.Since(start)
	// 	fmt.Println("result =", l)
	// 	fmt.Println("MaxGorouties, elapsed time to sort with workerpool: ", workerpool_waitgroups.Max, elapsed)
	// }

	// workerpool_waitgroups.Max = 4
	// start := time.Now()
	// workerpool_waitgroups.MergeSort(lst)
	// elapsed := time.Since(start)
	// fmt.Println("MaxGorouties, elapsed time to sort with workerpool_waitgroups: ", workerpool_waitgroups.Max, elapsed)

	// for mix_waitgroups.Max = 0; mix_waitgroups.Max < 50; mix_waitgroups.Max++ {
	// 	start := time.Now()
	// 	mix_waitgroups.MergeSort(lst)
	// 	elapsed := time.Since(start)
	// 	fmt.Println("MaxGorouties, elapsed time to sort with mix_waitgroups: ", mix_waitgroups.Max, elapsed)
	// }

	for mix_channel.Max = 0; mix_channel.Max < 50; mix_channel.Max++ {
		start := time.Now()
		mix_channel.MergeSort(lst)
		elapsed := time.Since(start)
		fmt.Println("MaxGorouties, elapsed time to sort with mix_channel: ", mix_channel.Max, elapsed)
	}
}
