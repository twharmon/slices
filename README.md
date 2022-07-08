# Slices
Pure functions for slices. Slices are never operated on "in place". Instead, new ones are returned.

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

	// append new item to end of slice
	s = slices.Append(s, "b")
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
goos: darwin
goarch: arm64
pkg: github.com/twharmon/slices

BenchmarkSortFunc/std_lib-10         	  116497	      9424 ns/op	    4152 B/op	       3 allocs/op
BenchmarkSortFunc/slices-10          	  101791	     11479 ns/op	    4096 B/op	       1 allocs/op
BenchmarkSort/std_lib-10             	  125790	      9441 ns/op	    4120 B/op	       2 allocs/op
BenchmarkSort/slices-10              	  113931	     10162 ns/op	    4096 B/op	       1 allocs/op

BenchmarkReverse-10                  	 1631348	       749.4 ns/op	    4096 B/op	       1 allocs/op
BenchmarkConcat-10                   	  275463	      4259 ns/op	   40960 B/op	       1 allocs/op
BenchmarkFilter-10                   	 1000000	      1018 ns/op	    4096 B/op	       1 allocs/op

BenchmarkUnion/2x20-10               	 3258876	       361.2 ns/op	      32 B/op	       1 allocs/op
BenchmarkUnion/20x2-10               	  502176	      2329 ns/op	    1592 B/op	       3 allocs/op
BenchmarkUnion/20x2000-10            	    1140	   1054108 ns/op	    1590 B/op	       3 allocs/op
BenchmarkUnion/2000x20-10            	     714	   1697117 ns/op	  142481 B/op	      39 allocs/op

BenchmarkIntersection/2x20-10        	 2380771	       498.0 ns/op	      48 B/op	       1 allocs/op
BenchmarkIntersection/20x2-10        	  821226	      1427 ns/op	    1196 B/op	       3 allocs/op
BenchmarkIntersection/20x2000-10     	    1178	   1020620 ns/op	    1282 B/op	       3 allocs/op
BenchmarkIntersection/2000x20-10     	     938	   1276258 ns/op	  203294 B/op	      58 allocs/op
```

## Contribute
Make a pull request.
