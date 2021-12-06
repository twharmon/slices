# Slice
A simple package that makes working with slices a little bit easier with the help of generics.

## Install
`go get github.com/twharmon/slice`

## Example
```go
package main

import (
	"fmt"
	
	"github.com/twharmon/slice"
)

func main() {
    s := slice.New("foo", "ba")
    s.Push("b")
    s.Sort(func(a, b string) bool {
        // sort by string length, ascending
        return len(a) < len(b)
    })
    fmt.Println(s) // [b ba foo]

    i := slice.New(-3, 4, 2, -8)
    i.Sort(func(a, b int) bool {
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
