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

	// define a slice of functions
	var sorters = []mergesort.SortFunc{
		synchronous.MergeSort,
		channel.MergeSort,
		waitgroups.MergeSort,
	}

	var names = []string{
		"synchronous",
		"channel",
		"waitgroups",
	}

	for idx, exec := range sorters {
		start := time.Now()
		exec.MergeSort(lst)
		elapsed := time.Since(start)
		fmt.Println("Elapsed time to sort:", names[idx], elapsed)
	}
}
