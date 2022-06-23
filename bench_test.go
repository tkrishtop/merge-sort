package main

import (
	"math/rand"
	"testing"
	"time"
)

var N = 10000000

func generateRandomArray() []int {
	rand.Seed(time.Now().UnixNano())
	lst := rand.Perm(N)

	return lst
}

var lst = generateRandomArray()

func BenchmarkMergeSortNoGoroutines(b *testing.B) {
	b.Log("Benchmark basic merge sort, array size = ", N)

	for i := 0; i < b.N; i++ {
		MergeSortNoGoroutines(lst)
	}
}

func BenchmarkMergeSortChannels(b *testing.B) {
	b.Log("Benchmark merge sort with goroutines / channels, array size = ", N)

	for i := 0; i < b.N; i++ {
		MergeSortChannels(lst)
	}
}

func BenchmarkMergeSortWaitGroups(b *testing.B) {
	b.Log("Benchmark merge sort with goroutines / wait groups, array size = ", N)

	for i := 0; i < b.N; i++ {
		MergeSortWaitGroups(lst)
	}
}
