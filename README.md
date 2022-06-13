# Test out goroutines

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
