# Slices
Pure functions for slices. Slices are never operated on "in place" but new ones are always returned.

![](https://github.com/twharmon/slices/workflows/Test/badge.svg) [![](https://goreportcard.com/badge/github.com/twharmon/slices)](https://goreportcard.com/report/github.com/twharmon/slices) [![codecov](https://codecov.io/gh/twharmon/slices/branch/main/graph/badge.svg?token=K0P59TPRAL)](https://codecov.io/gh/twharmon/slices)

## Documentation
For full documentation see [pkg.go.dev](https://pkg.go.dev/github.com/twharmon/slices).

## Install
`go get github.com/twharmon/slices`

## Usage
```go
package main

import (
	"fmt"

	"github.com/twharmon/slices"
)

func main() {
	// use plain go slices
	s := []string{"foo", "ba"}

	// push new item to end of slice
	s = slices.Push(s, "b")
	fmt.Println(s) // [foo ba b]

	
	// sort by string length, ascending
	sorted := slices.SortFunc(s, func(a, b string) bool {
		return len(a) < len(b)
	})
    // original slice is not chaged
	fmt.Println(s, sorted) // [foo ba b] [b ba foo]

	// sum the lengths of all the strings    
	totalLen := slices.Reduce(s, func(cnt int, i string) int {
		return cnt + len(i)
	})
	fmt.Println(totalLen) // 6

	// find the first item with length 2
	str := slices.Find(s, func(item string) bool { return len(item) == 2 })    
	fmt.Println(str) // ba


	// map slice to new slice of different type
    ints := slices.Map(s, func(item string) int { return len(s) })    
	fmt.Println(ints) // [3 2 1]
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
