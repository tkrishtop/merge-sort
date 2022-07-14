# Test out goroutines

## Testing semaphores

```
$ go test -bench ^BenchmarkMergeSort$ mergesort/test
goos: linux
goarch: amd64
pkg: mergesort/test
cpu: Intel(R) Core(TM) i7-10850H CPU @ 2.70GHz
BenchmarkMergeSort/synchronous_tested_on_array_of_size_1000000-12                      4         278863082 ns/op
BenchmarkMergeSort/channels_tested_on_array_of_size_1000000-12                         2         647550513 ns/op
BenchmarkMergeSort/waitgroups_tested_on_array_of_size_1000000-12                       2         527465514 ns/op

BenchmarkMergeSort/sem_channels_2_tested_on_array_of_size_1000000-12                   2         503141063 ns/op
BenchmarkMergeSort/sem_channels_4_tested_on_array_of_size_1000000-12                   2         529184726 ns/op
BenchmarkMergeSort/sem_channels_6_tested_on_array_of_size_1000000-12                   2         537195784 ns/op
BenchmarkMergeSort/sem_channels_8_tested_on_array_of_size_1000000-12                   2         528916088 ns/op
BenchmarkMergeSort/sem_channels_10_tested_on_array_of_size_1000000-12                  2         524825271 ns/op
BenchmarkMergeSort/sem_channels_12_tested_on_array_of_size_1000000-12                  2         508332486 ns/op
BenchmarkMergeSort/sem_channels_24_tested_on_array_of_size_1000000-12                  2         524019418 ns/op
BenchmarkMergeSort/sem_channels_100_tested_on_array_of_size_1000000-12                 2         549371054 ns/op
BenchmarkMergeSort/sem_channels_1000_tested_on_array_of_size_1000000-12                2         524972001 ns/op
BenchmarkMergeSort/sem_channels_10000_tested_on_array_of_size_1000000-12               2         537632214 ns/op
BenchmarkMergeSort/sem_channels_100000_tested_on_array_of_size_1000000-12              2         888260726 ns/op

BenchmarkMergeSort/sem_waitgroups_2_tested_on_array_of_size_1000000-12                 4         261331455 ns/op
BenchmarkMergeSort/sem_waitgroups_4_tested_on_array_of_size_1000000-12                 5         248030619 ns/op
BenchmarkMergeSort/sem_waitgroups_6_tested_on_array_of_size_1000000-12                 5         227603361 ns/op
BenchmarkMergeSort/sem_waitgroups_8_tested_on_array_of_size_1000000-12                 5         217228761 ns/op
BenchmarkMergeSort/sem_waitgroups_10_tested_on_array_of_size_1000000-12                5         221135752 ns/op
BenchmarkMergeSort/sem_waitgroups_12_tested_on_array_of_size_1000000-12                6         220947591 ns/op
BenchmarkMergeSort/sem_waitgroups_24_tested_on_array_of_size_1000000-12                6         218261241 ns/op
BenchmarkMergeSort/sem_waitgroups_100_tested_on_array_of_size_1000000-12               6         195841480 ns/op
BenchmarkMergeSort/sem_waitgroups_1000_tested_on_array_of_size_1000000-12              6         194972410 ns/op
BenchmarkMergeSort/sem_waitgroups_10000_tested_on_array_of_size_1000000-12             4         334979194 ns/op
BenchmarkMergeSort/sem_waitgroups_100000_tested_on_array_of_size_1000000-12            2         875434906 ns/op
PASS
ok      mergesort/test  49.712s```
```
