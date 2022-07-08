package slices_test

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"testing"

	"github.com/twharmon/slices"
)

var unsortedStringSlice = []string{"Mauris", "in", "luctus", "mi", "Suspendisse", "enim", "mi", "volutpat", "at", "urna", "ut", "condimentum", "feugiat", "lectus", "Nunc", "pulvinar", "arcu", "eget", "quam", "facilisis", "nec", "lobortis", "urna", "aliquet", "Vestibulum", "feugiat", "enim", "nec", "justo", "aliquam", "fringilla", "Fusce", "vitae", "ultrices", "orci", "Pellentesque", "consectetur", "ex", "quis", "fringilla", "finibus", "Nullam", "suscipit", "arcu", "suscipit", "vestibulum", "eros", "nec", "aliquet", "erat", "Pellentesque", "finibus", "sollicitudin", "libero", "sed", "lobortis", "Sed", "ut", "diam", "venenatis", "cursus", "velit", "non", "interdum", "lacus", "Mauris", "malesuada", "eros", "in", "dictum", "pellentesque", "ante", "tellus", "faucibus", "purus", "a", "ultricies", "sapien", "turpis", "nec", "sem", "Fusce", "sit", "amet", "aliquam", "turpis", "Phasellus", "id", "magna", "magna", "Morbi", "at", "erat", "est", "Pellentesque", "et", "porta", "loremSed", "fermentum", "metus", "at", "enim", "tempor", "auctor", "Phasellus", "hendrerit", "nunc", "sed", "cursus", "mattis", "sem", "mauris", "aliquet", "turpis", "ut", "gravida", "eros", "lacus", "non", "enim", "Quisque", "auctor", "turpis", "et", "nulla", "cursus", "sit", "amet", "blandit", "magna", "blandit", "In", "quis", "sodales", "ligula", "at", "aliquam", "lacus", "Duis", "dictum", "dapibus", "efficitur", "Sed", "sem", "tortor", "tincidunt", "vel", "sem", "ut", "suscipit", "porta", "mauris", "Mauris", "at", "nisl", "at", "odio", "cursus", "sollicitudin", "Etiam", "sit", "amet", "blandit", "ipsum", "Nullam", "et", "vehicula", "felis", "Integer", "sed", "elit", "ut", "ligula", "dignissim", "commodo", "Lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit", "Sed", "venenatis", "metus", "varius", "tristique", "auctor", "urna", "arcu", "pharetra", "quam", "ut", "consequat", "urna", "felis", "ut", "arcu", "Phasellus", "iaculis", "facilisis", "odio", "non", "finibus", "odio", "finibus", "vel", "Morbi", "risus", "quam", "laoreet", "sit", "amet", "maximus", "eu", "laoreet", "sit", "amet", "magna", "Interdum", "et", "malesuada", "fames", "ac", "ante", "ipsum", "primis", "in", "faucibus", "Etiam", "quis", "lacinia", "purus", "ut", "interdum", "diam"}

func BenchmarkSortFunc(b *testing.B) {
	b.Run("std lib", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := slices.Clone(unsortedStringSlice)
			sort.Slice(s, func(i, j int) bool {
				return s[i] < s[j]
			})
		}
	})
	b.Run("slices", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = slices.SortFunc(unsortedStringSlice, func(a string, b string) bool {
				return a < b
			})
		}
	})
}

func BenchmarkSort(b *testing.B) {
	b.Run("std lib", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := slices.Clone(unsortedStringSlice)
			sort.Strings(s)
		}
	})
	b.Run("slices", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = slices.Sort(unsortedStringSlice)
		}
	})
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = slices.Reverse(unsortedStringSlice)
	}
}

func BenchmarkConcat(b *testing.B) {
	s := make([][]string, 10)
	for i := 0; i < len(s); i++ {
		s[i] = slices.Clone(unsortedStringSlice)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = slices.Concat(s...)
	}
}

func BenchmarkFilter(b *testing.B) {
	filterFunc := func(v string) bool {
		return len(v) > 5
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = slices.Filter(unsortedStringSlice, filterFunc)
	}
}

func BenchmarkDistinct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = slices.Distinct(unsortedStringSlice)
	}
}

func BenchmarkUnion(b *testing.B) {
	makeStringSlice := func(size int, options int) []string {
		s := make([]string, size)
		for i := range s {
			s[i] = strconv.Itoa(rand.Intn(options + 1))
		}
		return s
	}
	run := func(sliceLen int, sliceCnt int) {
		s := make([][]string, sliceCnt)
		for i := range s {
			s[i] = makeStringSlice(sliceLen, sliceLen)
		}
		b.Run(fmt.Sprintf("%dx%d", sliceLen, sliceCnt), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = slices.Intersection(s...)
			}
		})
	}
	run(2, 20)
	run(20, 2)
	run(20, 2000)
	run(2000, 20)
}

func BenchmarkIntersection(b *testing.B) {
	makeStringSlice := func(size int, options int) []string {
		s := make([]string, size)
		for i := range s {
			s[i] = strconv.Itoa(rand.Intn(options + 1))
		}
		return s
	}
	run := func(sliceLen int, sliceCnt int) {
		s := make([][]string, sliceCnt)
		for i := range s {
			s[i] = makeStringSlice(sliceLen, sliceLen)
		}
		b.Run(fmt.Sprintf("%dx%d", sliceLen, sliceCnt), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = slices.Union(s...)
			}
		})
	}
	run(2, 20)
	run(20, 2)
	run(20, 2000)
	run(2000, 20)
}
