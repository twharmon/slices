# Slices
Utility functions for immutable slices.

![](https://github.com/twharmon/slices/workflows/Test/badge.svg) [![](https://goreportcard.com/badge/github.com/twharmon/slices)](https://goreportcard.com/report/github.com/twharmon/slices) [![](https://gocover.io/_badge/github.com/twharmon/slices)](https://gocover.io/github.com/twharmon/slices)

## Documentation
For full documentation see [pkg.go.dev](https://pkg.go.dev/github.com/twharmon/slices).

## Install
`go get github.com/twharmon/slices`

## Example
```go
package main

import (
	"fmt"
	
	"github.com/twharmon/slices"
)

func main() {
    s := []string{"foo", "ba"}
    s = slices.Push(s, "b")
    s = slices.SortFunc(s, func(a, b string) bool {
        // sort by string length, ascending
        return len(a) < len(b)
    })
    fmt.Println(s) // [b ba foo]

    i := []int{-3, 4, 2, -8}
    i = slices.SortFunc(i, func(a, b int) bool {
        // sort by asbolute value, ascending
        return a*a < b*b
    })
    fmt.Println(i) // [2, -3, 4, -8]
}
```

## Benchmarks
```
BenchmarkStdLibShortSort-10    	 1831299	       653.5 ns/op	     112 B/op	       4 allocs/op
BenchmarkSortShort-10       	 1634320	       732.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkStdLibSortLong-10     	  366229	      3261 ns/op	     112 B/op	       4 allocs/op
BenchmarkSortLong-10        	  292084	      4097 ns/op	       0 B/op	       0 allocs/op
```

## Contribute
Make a pull request.
