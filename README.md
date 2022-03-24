# Slices
Pure functions for slices.

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
BenchmarkStdLibSortFunc-10    	  743384	      1434 ns/op	    1064 B/op	       8 allocs/op
BenchmarkSortFuncShort-10     	  612254	      1946 ns/op	    2016 B/op	      12 allocs/op
BenchmarkStdLibSort-10        	  785678	      1477 ns/op	    1032 B/op	       7 allocs/op
BenchmarkSortShort-10         	  692116	      1672 ns/op	    2016 B/op	      12 allocs/op
```

## Contribute
Make a pull request.
