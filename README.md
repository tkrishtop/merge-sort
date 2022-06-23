# Test out goroutines

## Trying benchmarks with testing.B

```
# no goroutines
$ go test -benchmem -run=^$ -bench ^BenchmarkMergeSortNoGoroutines$
BenchmarkMergeSortNoGoroutines-12             45          24013591 ns/op        46403407 B/op     321018 allocs/op
--- BENCH: BenchmarkMergeSortNoGoroutines-12
    mergeSort_test.go:29: Benchmark basic merge sort, array size =  100000
    mergeSort_test.go:29: Benchmark basic merge sort, array size =  100000
PASS
ok      mergesort       1.113s


# goroutines
$ go test -benchmem -run=^$ -bench ^BenchmarkMergeSort$
BenchmarkMergeSort-12                 15          69144487 ns/op        76854888 B/op     738262 allocs/op
--- BENCH: BenchmarkMergeSort-12
    mergeSort_test.go:21: Benchmark merge sort with goroutines, array size =  100000
    mergeSort_test.go:21: Benchmark merge sort with goroutines, array size =  100000
    mergeSort_test.go:21: Benchmark merge sort with goroutines, array size =  100000
PASS
ok      mergesort       2.087s


# homemade comparison
$ go run main.go mergeSortNoGoroutines.go mergeSortWithGoroutines.go 

Number of elements in the array: 100000
Elapsed time to sort without goroutines: 33.139307ms
Elapsed time to sort using goroutines: 118.175829ms
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
