package mergesort

import (
	"math/rand"
	mergesort "mergesort/pkg"
	"mergesort/pkg/mergesort/channels"
	"mergesort/pkg/mergesort/semaphore_channels"
	"mergesort/pkg/mergesort/semaphore_waitgroups"
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
	var N = 1000000
	var lst = generateRandomArray(N)

	type namedSortFunc struct {
		name   string
		sorter mergesort.Sorter
	}

	var tests = []namedSortFunc{
		{
			name:   "synchronous",
			sorter: mergesort.SortFunc(synchronous.MergeSort),
		},
		{
			name:   "channels",
			sorter: mergesort.SortFunc(channels.MergeSort),
		},
		{
			name:   "waitgroups",
			sorter: mergesort.SortFunc(waitgroups.MergeSort),
		},
		{
			name:   "sem_channels_2",
			sorter: semaphore_channels.NewMixedSorter(2),
		},
		{
			name:   "sem_channels_4",
			sorter: semaphore_channels.NewMixedSorter(4),
		},
		{
			name:   "sem_channels_6",
			sorter: semaphore_channels.NewMixedSorter(6),
		},
		{
			name:   "sem_channels_8",
			sorter: semaphore_channels.NewMixedSorter(8),
		},
		{
			name:   "sem_channels_10",
			sorter: semaphore_channels.NewMixedSorter(10),
		},
		{
			name:   "sem_channels_12",
			sorter: semaphore_channels.NewMixedSorter(12),
		},
		{
			name:   "sem_channels_24",
			sorter: semaphore_channels.NewMixedSorter(24),
		},
		{
			name:   "sem_channels_100",
			sorter: semaphore_channels.NewMixedSorter(100),
		},
		{
			name:   "sem_channels_1000",
			sorter: semaphore_channels.NewMixedSorter(1000),
		},
		{
			name:   "sem_channels_10000",
			sorter: semaphore_channels.NewMixedSorter(10000),
		},
		{
			name:   "sem_channels_100000",
			sorter: semaphore_channels.NewMixedSorter(100000),
		},
		{
			name:   "sem_waitgroups_2",
			sorter: semaphore_waitgroups.NewMixedSorter(2),
		},
		{
			name:   "sem_waitgroups_4",
			sorter: semaphore_waitgroups.NewMixedSorter(4),
		},
		{
			name:   "sem_waitgroups_6",
			sorter: semaphore_waitgroups.NewMixedSorter(6),
		},
		{
			name:   "sem_waitgroups_8",
			sorter: semaphore_waitgroups.NewMixedSorter(8),
		},
		{
			name:   "sem_waitgroups_10",
			sorter: semaphore_waitgroups.NewMixedSorter(10),
		},
		{
			name:   "sem_waitgroups_12",
			sorter: semaphore_waitgroups.NewMixedSorter(12),
		},
		{
			name:   "sem_waitgroups_24",
			sorter: semaphore_waitgroups.NewMixedSorter(24),
		},
		{
			name:   "sem_waitgroups_100",
			sorter: semaphore_waitgroups.NewMixedSorter(100),
		},
		{
			name:   "sem_waitgroups_1000",
			sorter: semaphore_waitgroups.NewMixedSorter(1000),
		},
		{
			name:   "sem_waitgroups_10000",
			sorter: semaphore_waitgroups.NewMixedSorter(10000),
		},
		{
			name:   "sem_waitgroups_100000",
			sorter: semaphore_waitgroups.NewMixedSorter(100000),
		},
	}

	for _, test := range tests {
		b.Run(test.name+" tested on array of size "+strconv.Itoa(N),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					test.sorter.MergeSort(lst)
				}
			},
		)
	}
}
