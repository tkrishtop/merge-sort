# Test out goroutines

## Benchmarks with testing.B

```
$ go test -bench ^BenchmarkMergeSort$ mergesort/test
goos: linux
goarch: amd64
pkg: mergesort/test

BenchmarkMergeSort/sorter-synchronous-12                      42          25485594 ns/op
BenchmarkMergeSort/sorter-channel-12                          14          81312436 ns/op
BenchmarkMergeSort/sorter-waitgroups-12                       18          61837880 ns/op
PASS
ok      mergesort/test  5.196s
```

## Homemade Benchmarks
```
$ go run main.go
Elapsed time to sort: synchronous 27.654091ms
Elapsed time to sort: channel 109.384741ms
Elapsed time to sort: waitgroups 101.247685ms
```
