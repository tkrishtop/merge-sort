# Test out goroutines

## Trying benchmarks with testing.B

```
$ go test -bench ^BenchmarkMergeSort$ mergesort/test
goos: linux
goarch: amd64
pkg: mergesort/test

BenchmarkMergeSort/sorter-0-12                43          24985546 ns/op
BenchmarkMergeSort/sorter-1-12                14          77368753 ns/op
BenchmarkMergeSort/sorter-2-12                18          62596688 ns/op
PASS
ok      mergesort/test  5.170s
```

## Homemade Benchmarks
```
$ go run main.go
Elapsed time to sort: 0 29.920647ms
Elapsed time to sort: 1 112.736571ms
Elapsed time to sort: 2 91.542557ms
```
