# Slice
A simple package that makes working with slices a little bit easier with the help of generics.


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

    i := slice.New(-33, 4, 2, -8)
    i.Sort(func(a, b int) bool {
        // sort by asbolute value, ascending
        return a*a < b*b
    })
    fmt.Println(i) // [2, -33, 4, -8]
}
```

## Benchmarks
```
BenchmarkStdShortSort-4   	 5534931	       209.7 ns/op	     104 B/op	       3 allocs/op
BenchmarkSortShort-4      	34339386	        36.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkStdSort-4        	 2434693	       524.5 ns/op	     264 B/op	       3 allocs/op
BenchmarkSort-4           	 2980101	       402.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkStdSortLong-4    	  343713	      3362 ns/op	     952 B/op	       3 allocs/op
BenchmarkSortLong-4       	  398869	      3259 ns/op	       0 B/op	       0 allocs/op
```

## Contribute
Make a pull request.
