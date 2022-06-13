package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Here we go")

	// generate a large array with shuffled numbers
	for N := 10; N < 100000000; N *= 10 {
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
}
