# Benchmark Test

Bubble Sort is the worst possible type of Sort Algorithm.
Compare to [Native Sort](api/services/sort_service_test.go#L60)

Using the following Benchmark Tests

```sh
$ go test -v -run="none" -bench=. -benchtime="3s"
goos: darwin
goarch: amd64
pkg: gotrain/GoTestingInteg/mygolang-testing/api/services
BenchmarkBubbleSort-8           339131985               10.6 ns/op
BenchmarkSort-8                 35946357                97.6 ns/op
BenchmarkBubbleSortS-8            568296              6406 ns/op
BenchmarkBSortLimit-8              46938             76386 ns/op
BenchmarkSortS-8                   46779             77380 ns/op
PASS
```