# X-Team Go Bootcamp

https://github.com/julianosouza/go-crash-course/tree/master/07-testing

## Interfaces

https://gobyexample.com/testing

#### Notes

**Basic Unit Test**
`go test -v -cover ./sort` Shows how we cannot just believe that the tests have passed or that the coverage is 100%

**Basic Integration Test**
`go test -v ./services`

**Basic DB Mock**
`go test -v ./calculator`

**Benchmark**
with 10 elements
BubbleSort x Sort = BubbleSort is faster
```
sort>go test -bench=.
goos: windows
goarch: amd64
pkg: github.com/gugazimmermann/xteam-bootcamp-go/sort
BenchmarkBubbleSortAscending-8          177773378                6.68 ns/op
BenchmarkSort-8                         13953894               121 ns/op
PASS
ok      github.com/gugazimmermann/xteam-bootcamp-go/sort  3.830s
```

with 100.000 elements
BubbleSort x Sort = Sort is way faster
```
sort>go test -bench=.
goos: windows
goarch: amd64
pkg: github.com/gugazimmermann/xteam-bootcamp-go/sort
BenchmarkBubbleSortAscending2-8                1        8012594600 ns/op
BenchmarkSort2-8                             217           5396318 ns/op
PASS
ok      github.com/gugazimmermann/xteam-bootcamp-go/sort  9.896s
```

**timeout**
`./sort/sort_test.go` /  `TestBubbleSortAscendingTimeOut`

**Testify**

https://github.com/stretchr/testify

`go test -v ./calc_testify`

