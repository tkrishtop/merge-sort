# Test out goroutines

## Trying benchmarks with testing.B

```
# no goroutines
$ go test -benchmem -run=^$ -bench ^BenchmarkMergeSortNoGoroutines$ mergesort

BenchmarkMergeSortNoGoroutines-12    	       1	3496454416 ns/op	8059488768 B/op	33033479 allocs/op
--- BENCH: BenchmarkMergeSortNoGoroutines-12
    bench_test.go:21: Benchmark basic merge sort, array size =  10000000
PASS
ok  	mergesort	4.063s


# goroutines / channels
go test -benchmem -run=^$ -bench ^BenchmarkMergeSortChannels$ mergesort

BenchmarkMergeSortChannels-12    	       1	9464648884 ns/op	11638110384 B/op	75388294 allocs/op
--- BENCH: BenchmarkMergeSortChannels-12
    bench_test.go:29: Benchmark merge sort with goroutines / channels, array size =  10000000
PASS
ok  	mergesort	10.357s


# goroutines / wait groups
go test -benchmem -run=^$ -bench ^BenchmarkMergeSortWaitGroups$ mergesort

BenchmarkMergeSortWaitGroups-12    	       1	8786366103 ns/op	10445542384 B/op	84567488 allocs/op
--- BENCH: BenchmarkMergeSortWaitGroups-12
    bench_test.go:37: Benchmark merge sort with goroutines / wait groups, array size =  10000000
PASS
ok  	mergesort	9.603s


# homemade comparison
$ go run main.go mergeSort*.go

Number of elements in the array: 10000000
Elapsed time to sort without goroutines: 3.298257048s
Elapsed time to sort using goroutines / channels: 9.448487501s
Elapsed time to sort using goroutines / waitGroups: 7.567289726s
```


## Homemade Benchmarks
```
$ go run *.go

Number of elements in the array: 10
Elapsed time to sort without goroutines: 5.674µs
Elapsed time to sort using goroutines: 104.985µs

Number of elements in the array: 100
Elapsed time to sort without goroutines: 22.667µs
Elapsed time to sort using goroutines: 175.58µs

Number of elements in the array: 1000
Elapsed time to sort without goroutines: 243.66µs
Elapsed time to sort using goroutines: 745.597µs

Number of elements in the array: 10000
Elapsed time to sort without goroutines: 3.436244ms
Elapsed time to sort using goroutines: 9.635354ms

Number of elements in the array: 100000
Elapsed time to sort without goroutines: 40.213296ms
Elapsed time to sort using goroutines: 103.150552ms

Number of elements in the array: 1000000
Elapsed time to sort without goroutines: 337.191223ms
Elapsed time to sort using goroutines: 852.668807ms

Number of elements in the array: 10000000
Elapsed time to sort without goroutines: 3.980890384s
Elapsed time to sort using goroutines: 9.509865839s
```
