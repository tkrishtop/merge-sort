# Test out goroutines

## Benchmarks with testing.B

```
$ go test -bench ^BenchmarkMergeSort$ mergesort/test
goos: linux
goarch: amd64
pkg: mergesort/test

go test -benchmem -run=^$ -bench ^BenchmarkMergeSort$ mergesort/test

BenchmarkMergeSort/synchronous_tested_on_array_of_size_100000-12         	      46	  24593932 ns/op	46403184 B/op	  321014 allocs/op
BenchmarkMergeSort/channel_tested_on_array_of_size_100000-12             	      16	  73703149 ns/op	76896723 B/op	  738542 allocs/op
BenchmarkMergeSort/waitgroups_tested_on_array_of_size_100000-12          	      20	  56929758 ns/op	66605624 B/op	  831452 allocs/op
PASS
ok  	mergesort/test	5.517s
```

## Homemade Benchmarks
```
$ go run main.go
Testing on array of size 100000
Elapsed time to sort: synchronous 35.125593ms
Elapsed time to sort: channel 104.685877ms
Elapsed time to sort: waitgroups 78.890767ms
```
