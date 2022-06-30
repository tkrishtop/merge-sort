package main

import (
	"fmt"
	"math/rand"
	mergesort "mergesort/pkg"
	"mergesort/pkg/mergesort/channel"
	"mergesort/pkg/mergesort/synchronous"
	"mergesort/pkg/mergesort/waitgroups"
	"time"
)

func generateRandomArray(N int) []int {
	rand.Seed(time.Now().UnixNano())
	lst := rand.Perm(N)

	return lst
}

func main() {
	var N = 100000
	var lst = generateRandomArray(N)
	fmt.Println("Testing on array of size", N)

	type namedSortFunc struct {
		name     string
		function mergesort.SortFunc
	}

	var sorters = []namedSortFunc{
		{
			name:     "synchronous",
			function: synchronous.MergeSort,
		},
		{
			name:     "channel",
			function: channel.MergeSort,
		},
		{
			name:     "waitgroups",
			function: waitgroups.MergeSort,
		},
	}

	for _, sorter := range sorters {
		start := time.Now()
		sorter.function.MergeSort(lst)
		elapsed := time.Since(start)
		fmt.Println("Elapsed time to sort:", sorter.name, elapsed)
	}
}
