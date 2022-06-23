package main

import (
	"math/rand"
	"testing"
	"time"
)

var N = 100000

func generateRandomArray() []int {
	rand.Seed(time.Now().UnixNano())
	lst := rand.Perm(N)

	return lst
}

var lst = generateRandomArray()

func BenchmarkMergeSort(b *testing.B) {
	b.Log("Benchmark merge sort with goroutines, array size = ", N)

	for i := 0; i < b.N; i++ {
		MergeSort(lst)
	}
}

func BenchmarkMergeSortNoGoroutines(b *testing.B) {
	b.Log("Benchmark basic merge sort, array size = ", N)

	for i := 0; i < b.N; i++ {
		MergeSortNoGoroutines(lst)
	}
}
