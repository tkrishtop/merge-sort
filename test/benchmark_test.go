package mergesort

import (
	"math/rand"
	mergesort "mergesort/pkg"
	"mergesort/pkg/mergesort/channel"
	"mergesort/pkg/mergesort/synchronous"
	"mergesort/pkg/mergesort/waitgroups"
	"mergesort/pkg/mergesort/workerpool"
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
	var N = 1000
	var lst = generateRandomArray(N)

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
		{
			name:     "workerpool",
			function: workerpool.MergeSort,
		},
	}

	for _, sorter := range sorters {
		b.Run(sorter.name+" tested on array of size "+strconv.Itoa(N),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					sorter.function(lst)
				}
			},
		)
	}
}
