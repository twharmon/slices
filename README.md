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

BenchmarkStdLibSortFunc-8   	  106189	      9576 ns/op	    4152 B/op	       3 allocs/op
BenchmarkSortFuncShort-8    	   96883	     11896 ns/op	    8192 B/op	       2 allocs/op
BenchmarkStdLibSort-8       	  125802	      9386 ns/op	    4120 B/op	       2 allocs/op
BenchmarkSortShort-8        	  120718	      9696 ns/op	    8192 B/op	       2 allocs/op

BenchmarkReverse/strings_237_pure-8         	 2007199	       613.9 ns/op	    4096 B/op	       1 allocs/op
BenchmarkReverse/strings_237-8              	 1950008	       616.5 ns/op	    4096 B/op	       1 allocs/op
BenchmarkReverse/strings_7584_pure-8        	   68170	     17541 ns/op	  122880 B/op	       1 allocs/op
BenchmarkReverse/strings_7584-8             	   68252	     17510 ns/op	  122880 B/op	       1 allocs/op
BenchmarkReverse/int64_70_pure-8            	14714168	        80.83 ns/op	     576 B/op	       1 allocs/op
BenchmarkReverse/int64_70-8                 	14874696	        81.01 ns/op	     576 B/op	       1 allocs/op

BenchmarkConcat/strings_6x237_pure-8        	  442657	      2474 ns/op	   24576 B/op	       1 allocs/op
BenchmarkConcat/strings_6x237-8             	  486512	      2468 ns/op	   24576 B/op	       1 allocs/op
BenchmarkConcat/int64_6x70_pure-8           	 4153170	       292.9 ns/op	    3456 B/op	       1 allocs/op
BenchmarkConcat/int64_6x70-8                	 4092565	       288.2 ns/op	    3456 B/op	       1 allocs/op

BenchmarkFilter/strings_237_(len>5)_pure-8  	 1285369	       918.2 ns/op	    4096 B/op	       1 allocs/op
BenchmarkFilter/strings_237_(len>5)-8       	 1277200	       933.4 ns/op	    4096 B/op	       1 allocs/op
BenchmarkFilter/int64_70_(val>40)_pure-8    	 6271178	       194.5 ns/op	     576 B/op	       1 allocs/op
BenchmarkFilter/int64_70_(val>40)-8         	 6183364	       194.4 ns/op	     576 B/op	       1 allocs/op

BenchmarkMap/strings_237_replace_a_b_pure-8 	  221564	      5347 ns/op	    4768 B/op	      99 allocs/op
BenchmarkMap/strings_237_replace_a_b-8      	  219157	      5379 ns/op	    4768 B/op	      99 allocs/op
BenchmarkMap/int64_237_*5_pure-8            	 9189534	       135.1 ns/op	     576 B/op	       1 allocs/op
BenchmarkMap/int64_237_*5-8                 	 8258142	       139.0 ns/op	     576 B/op	       1 allocs/op

BenchmarkIntersect/int64_70-46-47_pure-8    	  224929	      5325 ns/op	     576 B/op	       1 allocs/op
BenchmarkIntersect/int64_70-46-47-8         	  220717	      5423 ns/op	     576 B/op	       1 allocs/op
BenchmarkIntersect/int64_70-46-47_hash-8    	  208978	      5653 ns/op	    2417 B/op	       4 allocs/op
```

## Contribute
Make a pull request.
