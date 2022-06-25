package mergesort

import (
	"math/rand"
	mergesort "mergesort/pkg"
	"mergesort/pkg/mergesort/channel"
	"mergesort/pkg/mergesort/synchronous"
	"mergesort/pkg/mergesort/waitgroups"
	"strconv"
	"testing"
	"time"
)

func generateRandomArray(N int) []int {
	rand.Seed(time.Now().UnixNano())
	lst := rand.Perm(N)

	return lst
}

func BenchmarkMergeSort(b *testing.B) {
	var N = 100000
	var lst = generateRandomArray(N)

	b.Log("Benchmark, array size = ", N)

	// define a slice of functions
	var sorters = []mergesort.SortFunc{
		synchronous.MergeSort,
		channel.MergeSort,
		waitgroups.MergeSort,
	}

	for idx, exec := range sorters {
		b.Run("sorter-"+strconv.Itoa(idx),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					exec.MergeSort(lst)
				}
			},
		)
	}
}
