# Test out goroutines

## Testing semaphores

```
$ go test -bench ^BenchmarkMergeSort$ mergesort/test
goos: linux
goarch: amd64
pkg: mergesort/test
cpu: Intel(R) Core(TM) i7-10850H CPU @ 2.70GHz
BenchmarkMergeSort/synchronous_tested_on_array_of_size_10000-12                      447           2686703 ns/op
BenchmarkMergeSort/channel_tested_on_array_of_size_10000-12                          312           3674670 ns/op
BenchmarkMergeSort/waitgroups_tested_on_array_of_size_10000-12                       350           3051504 ns/op
BenchmarkMergeSort/mix_channel_tested_on_array_of_size_10000-12                      196           5721401 ns/op
BenchmarkMergeSort/mix_waitgroups_tested_on_array_of_size_10000-12                   327           3395364 ns/op
PASS
ok      mergesort/test  7.662s
```

## Vary max goroutines

Efficient but ugly with waitgroups

```
$ go run main.go 
Testing on array of size 100000
MaxGorouties, elapsed time to sort with mix_waitgroups:  0 25.280125ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  1 38.475792ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  2 85.713414ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  3 73.578046ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  4 63.897672ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  5 73.5212ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  6 73.696789ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  7 66.126773ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  8 70.999692ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  9 69.88496ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  10 75.923891ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  11 71.100494ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  12 73.687804ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  13 74.102677ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  14 65.191949ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  15 75.284638ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  16 73.124474ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  17 72.856956ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  18 70.798082ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  19 77.577048ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  20 71.713677ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  21 78.799624ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  22 74.069334ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  23 77.115241ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  24 69.72653ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  25 79.344264ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  26 79.985829ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  27 78.782757ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  28 78.390523ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  29 74.578587ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  30 78.891325ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  31 80.628108ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  32 78.682167ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  33 80.218254ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  34 76.729696ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  35 76.643802ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  36 75.331885ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  37 83.35433ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  38 75.994047ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  39 78.087902ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  40 78.4177ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  41 83.001192ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  42 78.350368ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  43 77.838968ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  44 79.921509ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  45 82.971925ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  46 82.329917ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  47 83.268216ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  48 77.978733ms
MaxGorouties, elapsed time to sort with mix_waitgroups:  49 77.251178ms
```

Well-written but not efficient

```
$ go run main.go 
Testing on array of size 1000
MaxGorouties, elapsed time to sort with mix_channel:  0 386.534µs
MaxGorouties, elapsed time to sort with mix_channel:  1 4.40095249s
MaxGorouties, elapsed time to sort with mix_channel:  2 4.414204741s
MaxGorouties, elapsed time to sort with mix_channel:  3 13.432042033s
MaxGorouties, elapsed time to sort with mix_channel:  4 4.080943329s
MaxGorouties, elapsed time to sort with mix_channel:  5 4.368644696s
MaxGorouties, elapsed time to sort with mix_channel:  6 15.214650104s
MaxGorouties, elapsed time to sort with mix_channel:  7 4.238464805s
MaxGorouties, elapsed time to sort with mix_channel:  8 7.639577747s
MaxGorouties, elapsed time to sort with mix_channel:  9 4.371116306s
MaxGorouties, elapsed time to sort with mix_channel:  10 7.921132652s
MaxGorouties, elapsed time to sort with mix_channel:  11 9.373492325s
MaxGorouties, elapsed time to sort with mix_channel:  12 4.489271123s
MaxGorouties, elapsed time to sort with mix_channel:  13 4.246922374s
MaxGorouties, elapsed time to sort with mix_channel:  14 6.018175241s
MaxGorouties, elapsed time to sort with mix_channel:  15 4.438091346s
MaxGorouties, elapsed time to sort with mix_channel:  16 7.789133261s
MaxGorouties, elapsed time to sort with mix_channel:  17 7.849180748s
MaxGorouties, elapsed time to sort with mix_channel:  18 4.438035418s
MaxGorouties, elapsed time to sort with mix_channel:  19 9.378744721s
MaxGorouties, elapsed time to sort with mix_channel:  20 4.367376192s
MaxGorouties, elapsed time to sort with mix_channel:  21 4.207179635s
MaxGorouties, elapsed time to sort with mix_channel:  22 4.257137243s
MaxGorouties, elapsed time to sort with mix_channel:  23 4.227572223s
MaxGorouties, elapsed time to sort with mix_channel:  24 4.514966327s
MaxGorouties, elapsed time to sort with mix_channel:  25 7.754495657s
MaxGorouties, elapsed time to sort with mix_channel:  26 4.46826782s
MaxGorouties, elapsed time to sort with mix_channel:  27 4.156374304s
MaxGorouties, elapsed time to sort with mix_channel:  28 4.086168052s
MaxGorouties, elapsed time to sort with mix_channel:  29 4.437829735s
MaxGorouties, elapsed time to sort with mix_channel:  30 4.397753537s
MaxGorouties, elapsed time to sort with mix_channel:  31 7.831230336s
MaxGorouties, elapsed time to sort with mix_channel:  32 4.232700805s
MaxGorouties, elapsed time to sort with mix_channel:  33 13.433844919s
MaxGorouties, elapsed time to sort with mix_channel:  34 4.423723164s
MaxGorouties, elapsed time to sort with mix_channel:  35 14.701112041s
MaxGorouties, elapsed time to sort with mix_channel:  36 7.851772011s
MaxGorouties, elapsed time to sort with mix_channel:  37 15.073910424s
MaxGorouties, elapsed time to sort with mix_channel:  38 4.2745187s
MaxGorouties, elapsed time to sort with mix_channel:  39 5.743801555s
MaxGorouties, elapsed time to sort with mix_channel:  40 4.219069289s
MaxGorouties, elapsed time to sort with mix_channel:  41 15.315638649s
MaxGorouties, elapsed time to sort with mix_channel:  42 4.196930325s
MaxGorouties, elapsed time to sort with mix_channel:  43 4.197485694s
MaxGorouties, elapsed time to sort with mix_channel:  44 7.821202594s
MaxGorouties, elapsed time to sort with mix_channel:  45 15.345928227s
MaxGorouties, elapsed time to sort with mix_channel:  46 4.19694878s
MaxGorouties, elapsed time to sort with mix_channel:  47 7.838420194s
MaxGorouties, elapsed time to sort with mix_channel:  48 4.226745357s
MaxGorouties, elapsed time to sort with mix_channel:  49 12.186180411s
```
