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
	MergeSortChannels(l)
	elapsed = time.Since(start)
	fmt.Println("Elapsed time to sort using goroutines / channels:", elapsed)

	start = time.Now()
	MergeSortWaitGroups(l)
	elapsed = time.Since(start)
	fmt.Println("Elapsed time to sort using goroutines / waitGroups:", elapsed)
}

func main() {
	homemadeBenchmark(10000000)
}
