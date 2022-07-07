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
Benchmarks are compared with the same methods written using pure go.  
The results for MacBook Air M1, MacOS 12.0.1, Go 1.18.1 are:
```
goos: darwin
goarch: arm64
pkg: github.com/twharmon/slices
BenchmarkStdLibSortFunc-10    	  119398	      9323 ns/op	    4152 B/op	       3 allocs/op
BenchmarkSortFuncShort-10     	   94508	     12153 ns/op	    8192 B/op	       2 allocs/op
BenchmarkStdLibSort-10        	  126046	      9241 ns/op	    4120 B/op	       2 allocs/op
BenchmarkSortShort-10         	  115418	     10242 ns/op	    8192 B/op	       2 allocs/op

BenchmarkReverse/strings_237_pure-10         	 1947543	       650.8 ns/op	    4096 B/op	       1 allocs/op
BenchmarkReverse/strings_237-10              	 1653981	       740.0 ns/op	    4096 B/op	       1 allocs/op
BenchmarkReverse/strings_7584_pure-10        	   66444	     18537 ns/op	  122880 B/op	       1 allocs/op
BenchmarkReverse/strings_7584-10             	   44816	     26791 ns/op	  122880 B/op	       1 allocs/op
BenchmarkReverse/int64_488_pure-10           	 1934253	       621.5 ns/op	    4096 B/op	       1 allocs/op
BenchmarkReverse/int64_488-10                	 1994956	       590.0 ns/op	    4096 B/op	       1 allocs/op

BenchmarkConcat/strings_6x237_pure-10        	  428758	      2576 ns/op	   24576 B/op	       1 allocs/op
BenchmarkConcat/strings_6x237-10             	  489048	      2506 ns/op	   24576 B/op	       1 allocs/op
BenchmarkConcat/int64_6x488_pure-10          	  696577	      1754 ns/op	   24576 B/op	       1 allocs/op
BenchmarkConcat/int64_6x488-10               	  678524	      1750 ns/op	   24576 B/op	       1 allocs/op

BenchmarkFilter/strings_237_(len>5)_pure-10  	 1283282	       931.6 ns/op	    4096 B/op	       1 allocs/op
BenchmarkFilter/strings_237_(len>5)-10       	 1274114	       987.7 ns/op	    4096 B/op	       1 allocs/op
BenchmarkFilter/int64_488_(val>40)_pure-10   	  867286	      1447 ns/op	    4096 B/op	       1 allocs/op
BenchmarkFilter/int64_488_(val>40)-10        	  783907	      1519 ns/op	    4096 B/op	       1 allocs/op

BenchmarkMap/strings_237_replace_a_b_pure-10 	  225678	      5187 ns/op	    4768 B/op	      99 allocs/op
BenchmarkMap/strings_237_replace_a_b-10      	  225409	      5269 ns/op	    4768 B/op	      99 allocs/op
BenchmarkMap/int64_237_*5_pure-10            	 1443272	       841.6 ns/op	    4096 B/op	       1 allocs/op
BenchmarkMap/int64_237_*5-10                 	 1390299	       850.0 ns/op	    4096 B/op	       1 allocs/op

BenchmarkIntersect/int_10x10_pure-sm-result-10         	  692601	      1709 ns/op	      80 B/op	       1 allocs/op
BenchmarkIntersect/int_10x10_pkg-sm-result-10          	 1338831	       894.7 ns/op	     112 B/op	       3 allocs/op
BenchmarkIntersect/int_10x10_pure-lg-result-10         	 8502908	       141.3 ns/op	      80 B/op	       1 allocs/op
BenchmarkIntersect/int_10x10_pkg-lg-result-10          	 2866723	       412.4 ns/op	      56 B/op	       3 allocs/op
BenchmarkIntersect/int_10x100_pure-sm-result-10        	   83098	     14213 ns/op	      80 B/op	       1 allocs/op
BenchmarkIntersect/int_10x100_pkg-sm-result-10         	  191071	      5585 ns/op	      96 B/op	       3 allocs/op
BenchmarkIntersect/int_10x100_pure-lg-result-10        	 1000000	      1181 ns/op	      80 B/op	       1 allocs/op
BenchmarkIntersect/int_10x100_pkg-lg-result-10         	  367261	      3262 ns/op	      56 B/op	       3 allocs/op
BenchmarkIntersect/int_100x10000_pure-sm-result-10     	      16	  67781260 ns/op	     896 B/op	       1 allocs/op
BenchmarkIntersect/int_100x10000_pkg-sm-result-10      	      76	  15645106 ns/op	    5719 B/op	      13 allocs/op
BenchmarkIntersect/int_100x10000_pure-lg-result-10     	       1	3796862833 ns/op	     896 B/op	       1 allocs/op
BenchmarkIntersect/int_100x10000_pkg-lg-result-10      	      58	  19518427 ns/op	     416 B/op	       4 allocs/op
BenchmarkIntersect/int_10000x100_pure-sm-result-10     	       1	5512020125 ns/op	   81920 B/op	       1 allocs/op
BenchmarkIntersect/int_10000x100_pkg-sm-result-10      	      67	  17902264 ns/op	  413049 B/op	     203 allocs/op
BenchmarkIntersect/int_10000x100_pure-lg-result-10     	       4	 282252500 ns/op	   81920 B/op	       1 allocs/op
BenchmarkIntersect/int_10000x100_pkg-lg-result-10      	      67	  17784601 ns/op	   91609 B/op	      37 allocs/op
```

## Contribute
Make a pull request.
