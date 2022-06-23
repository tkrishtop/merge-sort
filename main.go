package main

import (
	"fmt"
	"math/rand"
	"time"
)

func homemadeBenchmark(N int) {
	// generate a large array with shuffled numbers
	fmt.Println("\nNumber of elements in the array:", N)
	rand.Seed(time.Now().UnixNano())
	l := rand.Perm(N)

	start := time.Now()
	MergeSortNoGoroutines(l)
	elapsed := time.Since(start)
	fmt.Println("Elapsed time to sort without goroutines:", elapsed)

	start = time.Now()
	MergeSort(l)
	elapsed = time.Since(start)
	fmt.Println("Elapsed time to sort using goroutines:", elapsed)
}

func main() {
	homemadeBenchmark(100000)
}
