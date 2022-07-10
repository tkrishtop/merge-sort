# Test out goroutines

## Benchmarks with testing.B

```
$ $ go test -bench ^BenchmarkMergeSort$ mergesort/test
goos: linux
goarch: amd64
pkg: mergesort/test

BenchmarkMergeSort/synchronous_tested_on_array_of_size_1000-12              5379            233282 ns/op
BenchmarkMergeSort/channel_tested_on_array_of_size_1000-12                  3038            344868 ns/op
BenchmarkMergeSort/waitgroups_tested_on_array_of_size_1000-12               3606            291932 ns/op
BenchmarkMergeSort/workerpool_tested_on_array_of_size_1000-12                  1        7829637523 ns/op
```

## Homemade Benchmarks
```
$ go run main.go
Testing on array of size 1000
Elapsed time to sort: synchronous 248.589µs
Elapsed time to sort: channel 1.209889ms
Elapsed time to sort: waitgroups 355.875µs
Elapsed time to sort: workerpool 9.357412224s
```
