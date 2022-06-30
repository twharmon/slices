package slices_test

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/twharmon/slices"
)

func TestPush(t *testing.T) {
	s := []int{1}
	got := slices.Append(s, 5)
	want := []int{1, 5}
	assertEqual(t, got, want)
}

func TestContainsTrue(t *testing.T) {
	s := []int{1, 2, 3}
	got := slices.Contains(s, 2)
	want := true
	assertEqual(t, got, want)
}

func TestContainsFalse(t *testing.T) {
	s := []int{1, 2, 3}
	got := slices.Contains(s, 4)
	want := false
	assertEqual(t, got, want)
}

func TestSpliceNoInserts(t *testing.T) {
	s := []string{"foo", "bar", "baz"}
	got := slices.Splice(s, 1, 1)
	want := []string{"foo", "baz"}
	assertEqual(t, got, want)
}

func TestSpliceWithInserts(t *testing.T) {
	s := []string{"foo", "bar", "baz"}
	got := slices.Splice(s, 1, 1, "boo")
	want := []string{"foo", "boo", "baz"}
	assertEqual(t, got, want)
}

func TestReverseEven(t *testing.T) {
	s := []string{"foo", "bar"}
	got := slices.Reverse(s)
	want := []string{"bar", "foo"}
	assertEqual(t, got, want)
}

func TestReverseOdd(t *testing.T) {
	s := []string{"foo", "bar", "baz"}
	got := slices.Reverse(s)
	want := []string{"baz", "bar", "foo"}
	assertEqual(t, got, want)
}

func TestUnshift(t *testing.T) {
	s := []string{"foo"}
	got := slices.Unshift(s, "bar")
	want := []string{"bar", "foo"}
	assertEqual(t, got, want)
}

func TestFindFound(t *testing.T) {
	s := []string{"foo"}
	got := slices.Find(s, func(item string) bool {
		return item == "foo"
	})
	want := "foo"
	assertEqual(t, got, want)
}

func TestFindNotFound(t *testing.T) {
	s := []string{"foo"}
	got := slices.Find(s, func(item string) bool {
		return item == "bar"
	})
	want := ""
	assertEqual(t, got, want)
}

func TestFilter(t *testing.T) {
	s := []string{"foo", "bar"}
	got := slices.Filter(s, func(item string) bool {
		return strings.HasPrefix(item, "f")
	})
	assertEqual(t, len(got), 1)
}

func TestMaxZeroValue(t *testing.T) {
	s := []string{}
	got := slices.Max(s)
	want := ""
	assertEqual(t, got, want)
}
func TestMinZeroValue(t *testing.T) {
	s := []string{}
	got := slices.Min(s)
	want := ""
	assertEqual(t, got, want)
}

func TestMax(t *testing.T) {
	s := []int{2, 6, 1, 4, 3}
	got := slices.Max(s)
	want := 6
	assertEqual(t, got, want)
}

func TestMin(t *testing.T) {
	s := []int{2, 6, 1, 4, 3}
	got := slices.Min(s)
	want := 1
	assertEqual(t, got, want)
}

func TestMaxFuncZeroValue(t *testing.T) {
	s := []int{}
	got := slices.MaxFunc(s, func(a, b int) bool { return a < b })
	want := 0
	assertEqual(t, got, want)
}

func TestMinFuncZeroValue(t *testing.T) {
	s := []int{}
	got := slices.MinFunc(s, func(a, b int) bool { return a < b })
	want := 0
	assertEqual(t, got, want)
}

func TestMaxFunc(t *testing.T) {
	s := []int{2, 6, 1, 4, 3}
	got := slices.MaxFunc(s, func(a, b int) bool { return a < b })
	want := 6
	assertEqual(t, got, want)
}

func TestMinFunc(t *testing.T) {
	s := []int{2, 6, 1, 4, 3}
	got := slices.MinFunc(s, func(a, b int) bool { return a < b })
	want := 1
	assertEqual(t, got, want)
}

func TestSomeTrue(t *testing.T) {
	s := []string{"foo", "bar", "baz"}
	if !slices.Some(s, func(s string) bool { return s == "bar" }) {
		t.Fatalf("some was false; expected true")
	}
}

func TestSomeFalse(t *testing.T) {
	s := []string{"foo", "bar", "baz"}
	if slices.Some(s, func(s string) bool { return s == "x" }) {
		t.Fatalf("some was true; expected false")
	}
}

func TestEveryFalse(t *testing.T) {
	s := []string{"foo", "bar", "baz"}
	if slices.Every(s, func(s string) bool { return len(s) < 2 }) {
		t.Fatalf("every was true; expected false")
	}
}

func TestEveryTrue(t *testing.T) {
	s := []string{"foo", "bar", "baz"}
	if !slices.Every(s, func(s string) bool { return len(s) == 3 }) {
		t.Fatalf("every was true; expected false")
	}
}

func TestMap(t *testing.T) {
	s := []string{"f", "ba", "baz"}
	got := slices.Map(s, func(i string) int {
		return len(i)
	})
	want := []int{1, 2, 3}
	assertEqual(t, got, want)
}

func TestReduce(t *testing.T) {
	s := []string{"f", "ba", "baz"}
	got := slices.Reduce(s, func(cnt int, i string) int {
		return cnt + len(i)
	})
	want := 6
	assertEqual(t, got, want)
}

var unsortedString = []string{"Mauris", "in", "luctus", "mi", "Suspendisse", "enim", "mi", "volutpat", "at", "urna", "ut", "condimentum", "feugiat", "lectus", "Nunc", "pulvinar", "arcu", "eget", "quam", "facilisis", "nec", "lobortis", "urna", "aliquet", "Vestibulum", "feugiat", "enim", "nec", "justo", "aliquam", "fringilla", "Fusce", "vitae", "ultrices", "orci", "Pellentesque", "consectetur", "ex", "quis", "fringilla", "finibus", "Nullam", "suscipit", "arcu", "suscipit", "vestibulum", "eros", "nec", "aliquet", "erat", "Pellentesque", "finibus", "sollicitudin", "libero", "sed", "lobortis", "Sed", "ut", "diam", "venenatis", "cursus", "velit", "non", "interdum", "lacus", "Mauris", "malesuada", "eros", "in", "dictum", "pellentesque", "ante", "tellus", "faucibus", "purus", "a", "ultricies", "sapien", "turpis", "nec", "sem", "Fusce", "sit", "amet", "aliquam", "turpis", "Phasellus", "id", "magna", "magna", "Morbi", "at", "erat", "est", "Pellentesque", "et", "porta", "loremSed", "fermentum", "metus", "at", "enim", "tempor", "auctor", "Phasellus", "hendrerit", "nunc", "sed", "cursus", "mattis", "sem", "mauris", "aliquet", "turpis", "ut", "gravida", "eros", "lacus", "non", "enim", "Quisque", "auctor", "turpis", "et", "nulla", "cursus", "sit", "amet", "blandit", "magna", "blandit", "In", "quis", "sodales", "ligula", "at", "aliquam", "lacus", "Duis", "dictum", "dapibus", "efficitur", "Sed", "sem", "tortor", "tincidunt", "vel", "sem", "ut", "suscipit", "porta", "mauris", "Mauris", "at", "nisl", "at", "odio", "cursus", "sollicitudin", "Etiam", "sit", "amet", "blandit", "ipsum", "Nullam", "et", "vehicula", "felis", "Integer", "sed", "elit", "ut", "ligula", "dignissim", "commodo", "Lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit", "Sed", "venenatis", "metus", "varius", "tristique", "auctor", "urna", "arcu", "pharetra", "quam", "ut", "consequat", "urna", "felis", "ut", "arcu", "Phasellus", "iaculis", "facilisis", "odio", "non", "finibus", "odio", "finibus", "vel", "Morbi", "risus", "quam", "laoreet", "sit", "amet", "maximus", "eu", "laoreet", "sit", "amet", "magna", "Interdum", "et", "malesuada", "fames", "ac", "ante", "ipsum", "primis", "in", "faucibus", "Etiam", "quis", "lacinia", "purus", "ut", "interdum", "diam"}
var unsortedInt64 = []int64{53, 88, 23, 25, 64, 71, 7, 83, 17, 33, 12, 31, 69, 14, 90, 77, 22, 2, 96, 10, 45, 47, 35, 89, 49, 42, 76, 32, 15, 75, 62, 79, 72, 27, 57, 5, 59, 30, 61, 60, 9, 67, 40, 85, 46, 73, 34, 65, 36, 82, 20, 4, 3, 13, 58, 99, 24, 1, 51, 78, 100, 86, 28, 26, 68, 41, 43, 91, 18, 55}

func TestSortFunc(t *testing.T) {
	s := slices.Clone(unsortedString)
	got := slices.SortFunc(s, func(a string, b string) bool {
		return a < b
	})
	want := slices.Clone(s)
	sort.Slice(want, func(i, j int) bool {
		return want[i] < want[j]
	})
	assertEqual(t, got, want)
}

func TestSort(t *testing.T) {
	s := slices.Clone(unsortedString)
	got := slices.Sort(s)
	want := slices.Clone(s)
	sort.Strings(want)
	assertEqual(t, got, want)
}

func TestIntersection(t *testing.T) {
	t.Run("common", func(t *testing.T) {
		a := []string{"foo", "bar"}
		b := []string{"bar", "baz"}
		want := []string{"bar"}
		got := slices.Intersection(a, b)
		assertEqual(t, got, want)
	})

	t.Run("empty slice", func(t *testing.T) {
		var s [][]string
		got := slices.Intersection(s...)
		assertEqual(t, got, []string{})
	})
}

func TestIntersectionHash(t *testing.T) {
	t.Run("common", func(t *testing.T) {
		a := []string{"foo", "bar", "baz"}
		b := []string{"bar", "baz"}
		want := []string{"bar", "baz"}
		got := slices.Sort(slices.IntersectionHash(a, b))
		assertEqual(t, got, want)
	})

	t.Run("empty slice", func(t *testing.T) {
		var s [][]string
		got := slices.IntersectionHash(s...)
		assertEqual(t, got, []string{})
	})
}

func TestUnion(t *testing.T) {
	a := []string{"foo", "bar"}
	b := []string{"bar", "baz"}
	want := slices.Sort([]string{"foo", "bar", "baz"})
	got := slices.Sort(slices.Union(a, b))
	assertEqual(t, got, want)
}

func BenchmarkStdLibSortFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := slices.Clone(unsortedString)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
	}
}

func BenchmarkSortFuncShort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := slices.Clone(unsortedString)
		_ = slices.SortFunc(s, func(a string, b string) bool {
			return a < b
		})
	}
}

func BenchmarkStdLibSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := slices.Clone(unsortedString)
		sort.Strings(s)
	}
}

func BenchmarkSortShort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := slices.Clone(unsortedString)
		_ = slices.Sort(s)
	}
}

func BenchmarkReverse(b *testing.B) {
	s0 := slices.Clone(unsortedString)
	b.Run(fmt.Sprintf("strings_%d_pure", len(s0)), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = reverseStringsPure(s0)
		}
	})

	s1 := slices.Clone(unsortedString)
	b.Run(fmt.Sprintf("strings_%d", len(s1)), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = slices.Reverse(s1)
		}
	})

	sl0 := slices.Clone(unsortedString)
	for i := 0; i < 5; i++ {
		sl0 = append(sl0, sl0...)
	}
	b.Run(fmt.Sprintf("strings_%d_pure", len(sl0)), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = reverseStringsPure(sl0)
		}
	})

	sl1 := slices.Clone(unsortedString)
	for i := 0; i < 5; i++ {
		sl1 = append(sl1, sl1...)
	}

	b.Run(fmt.Sprintf("strings_%d", len(sl1)), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = slices.Reverse(sl1)
		}
	})

	i0 := slices.Clone(unsortedInt64)
	b.Run(fmt.Sprintf("int64_%d_pure", len(i0)), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = reverseInt64sPure(i0)
		}
	})

	i1 := slices.Clone(unsortedInt64)
	b.Run(fmt.Sprintf("int64_%d", len(i1)), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = slices.Reverse(i1)
		}
	})
}

func reverseStringsPure(v []string) []string {
	reversed := make([]string, len(v))
	for j := range v {
		reversed[j] = v[len(v)-j-1]
	}
	return reversed
}

func reverseInt64sPure(v []int64) []int64 {
	reversed := make([]int64, len(v))
	for j := range v {
		reversed[j] = v[len(v)-j-1]
	}

	return reversed
}

func BenchmarkConcat(b *testing.B) {
	n := 6
	s0 := make([][]string, n)
	for i := 0; i < n; i++ {
		s0[i] = slices.Clone(unsortedString)
	}

	b.Run(fmt.Sprintf("strings_%dx%d_pure", n, len(unsortedString)), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = concanStringsPure(s0[0], s0[1], s0[2], s0[3], s0[4], s0[5])
		}
	})

	s1 := make([][]string, n)
	for i := 0; i < n; i++ {
		s1[i] = slices.Clone(unsortedString)
	}

	b.Run(fmt.Sprintf("strings_%dx%d", n, len(unsortedString)), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = slices.Concat(s1[0], s1[1], s1[2], s1[3], s1[4], s1[5])
		}
	})

	i0 := make([][]int64, n)
	for i := 0; i < n; i++ {
		i0[i] = slices.Clone(unsortedInt64)
	}
	b.Run(fmt.Sprintf("int64_%dx%d_pure", n, len(unsortedInt64)), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = concanInt64Pure(i0[0], i0[1], i0[2], i0[3], i0[4], i0[5])
		}
	})

	i1 := make([][]int64, n)
	for i := 0; i < n; i++ {
		i1[i] = slices.Clone(unsortedInt64)
	}
	b.Run(fmt.Sprintf("int64_%dx%d", n, len(unsortedInt64)), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = slices.Concat(i1[0], i1[1], i1[2], i1[3], i1[4], i1[5])
		}
	})
}

func concanInt64Pure(v ...[]int64) []int64 {
	totalLen := 0
	for j := range v {
		totalLen += len(v[j])
	}

	result := make([]int64, 0, totalLen)
	for j := range v {
		result = append(result, v[j]...)
	}

	return result
}

func concanStringsPure(v ...[]string) []string {
	totalLen := 0
	for j := range v {
		totalLen += len(v[j])
	}

	result := make([]string, 0, totalLen)
	for j := range v {
		result = append(result, v[j]...)
	}

	return result
}

func BenchmarkFilter(b *testing.B) {
	sn := 5
	sFunc := func(v string) bool {
		return len(v) > sn
	}
	s0 := slices.Clone(unsortedString)
	b.Run(fmt.Sprintf("strings_%d_(len>%d)_pure", len(s0), sn), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = filterStringsPure(s0, sFunc)
		}
	})

	s1 := slices.Clone(unsortedString)
	b.Run(fmt.Sprintf("strings_%d_(len>%d)", len(s1), sn), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = slices.Filter(s0, sFunc)
		}
	})

	in := int64(40)
	iFunc := func(v int64) bool {
		return v > in
	}
	i0 := slices.Clone(unsortedInt64)
	b.Run(fmt.Sprintf("int64_%d_(val>%d)_pure", len(i0), in), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = filterInt64Pure(i0, iFunc)
		}
	})

	i1 := slices.Clone(unsortedInt64)
	b.Run(fmt.Sprintf("int64_%d_(val>%d)", len(i1), in), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = slices.Filter(i1, iFunc)
		}
	})
}

func filterStringsPure(v []string, f func(string) bool) []string {
	result := make([]string, 0, len(v))
	for i := range v {
		if f(v[i]) {
			result = append(result, v[i])
		}
	}

	return result
}

func filterInt64Pure(v []int64, f func(int64) bool) []int64 {
	result := make([]int64, 0, len(v))
	for i := range v {
		if f(v[i]) {
			result = append(result, v[i])
		}
	}

	return result
}

func BenchmarkMap(b *testing.B) {
	sFunc := func(s string) string {
		return strings.ReplaceAll(s, "a", "b")
	}
	s0 := slices.Clone(unsortedString)
	b.Run(fmt.Sprintf("strings_%d_replace_a_b_pure", len(s0)), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = mapStringsPure(s0, sFunc)
		}
	})

	s1 := slices.Clone(unsortedString)
	b.Run(fmt.Sprintf("strings_%d_replace_a_b", len(s1)), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = slices.Map(s1, sFunc)
		}
	})

	in := int64(5)
	iFunc := func(i int64) int64 {
		return i * in
	}
	i0 := slices.Clone(unsortedInt64)
	b.Run(fmt.Sprintf("int64_%d_*%d_pure", len(s0), in), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = mapInt64Pure(i0, iFunc)
		}
	})

	i1 := slices.Clone(unsortedInt64)
	b.Run(fmt.Sprintf("int64_%d_*%d", len(s1), in), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = slices.Map(i1, iFunc)
		}
	})
}

func mapStringsPure(v []string, f func(string) string) []string {
	result := make([]string, len(v))
	for j := range v {
		result[j] = f(v[j])
	}

	return result
}

func mapInt64Pure(v []int64, f func(int64) int64) []int64 {
	result := make([]int64, len(v))
	for j := range v {
		result[j] = f(v[j])
	}

	return result
}

func BenchmarkIntersect(b *testing.B) {
	i0 := slices.Clone(unsortedInt64)
	i1 := slices.Clone(unsortedInt64[0:int(float64(len(unsortedInt64))/1.5)])
	i2 := slices.Clone(unsortedInt64[len(unsortedInt64)/3:])

	b.Run(fmt.Sprintf("int64_%d-%d-%d_pure", len(i0), len(i1), len(i2)), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = intersectInt64Pure(i0, i1, i2)
		}
	})

	b.Run(fmt.Sprintf("int64_%d-%d-%d", len(i0), len(i1), len(i2)), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = slices.Intersection(i0, i1, i2)
		}
	})

	b.Run(fmt.Sprintf("int64_%d-%d-%d_hash", len(i0), len(i1), len(i2)), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = slices.IntersectionHash(i0, i1, i2)
		}
	})
}

func intersectInt64Pure(v ...[]int64) []int64 {
	result := make([]int64, 0, len(v[0]))

	isInSlices := func(t int64, v ...[]int64) bool {
		n := 0
		for i := range v {
			for j := range v[i] {
				if v[i][j] == t {
					n++
					break
				}
			}

			if n == i {
				return false
			}
		}

		return n == len(v)
	}
	isInSlice := func(t int64, v []int64) bool {
		for i := range v {
			if v[i] == t {
				return true
			}
		}

		return false
	}

	for i := range v {
		for j := range v[i] {
			if !isInSlice(v[i][j], result) && isInSlices(v[i][j], v...) {
				result = append(result, v[i][j])
			}
		}
	}

	return result
}

func TestBenchmarkPureFuncs(t *testing.T) {
	t.Run("reverse", func(t *testing.T) {
		s := []string{"a", "b", "c"}
		i := []int64{1, 2, 3}

		gotStringsPure := reverseStringsPure(s)
		gotInt64Pure := reverseInt64sPure(i)

		gotStringsLib := slices.Reverse(s)
		gotInt64Lib := slices.Reverse(i)

		assertEqual(t, gotStringsLib, gotStringsPure)
		assertEqual(t, gotInt64Lib, gotInt64Pure)
	})

	t.Run("concat", func(t *testing.T) {
		s0 := []string{"a", "b", "c"}
		s1 := []string{"d", "e", "f"}
		s2 := []string{"g", "h", "i"}

		i0 := []int64{1, 2, 3}
		i1 := []int64{4, 5, 6}
		i2 := []int64{7, 8, 9}

		gotStringsPure := concanStringsPure(s0, s1, s2)
		gotInt64Pure := concanInt64Pure(i0, i1, i2)

		gotStringsLib := slices.Concat(s0, s1, s2)
		gotInt64Lib := slices.Concat(i0, i1, i2)

		assertEqual(t, gotStringsLib, gotStringsPure)
		assertEqual(t, gotInt64Lib, gotInt64Pure)
	})

	t.Run("filter", func(t *testing.T) {
		s := []string{"a", "b", "c"}
		sFunc := func(v string) bool {
			return v == "a"
		}
		i := []int64{1, 2, 3}
		iFunc := func(v int64) bool {
			return v == 2
		}

		gotStringsPure := filterStringsPure(s, sFunc)
		gotInt64Pure := filterInt64Pure(i, iFunc)

		gotStringsLib := slices.Filter(s, sFunc)
		gotInt64Lib := slices.Filter(i, iFunc)

		assertEqual(t, gotStringsLib, gotStringsPure)
		assertEqual(t, gotInt64Lib, gotInt64Pure)
	})

	t.Run("map", func(t *testing.T) {
		s := []string{"a", "b", "c"}
		sFunc := func(v string) string {
			return v + v
		}
		i := []int64{1, 2, 3}
		iFunc := func(v int64) int64 {
			return v * v
		}

		gotStringsPure := mapStringsPure(s, sFunc)
		gotInt64Pure := mapInt64Pure(i, iFunc)

		gotStringsLib := slices.Map(s, sFunc)
		gotInt64Lib := slices.Map(i, iFunc)

		assertEqual(t, gotStringsLib, gotStringsPure)
		assertEqual(t, gotInt64Lib, gotInt64Pure)
	})

	t.Run("intersect", func(t *testing.T) {
		i0 := slices.Clone(unsortedInt64)
		i1 := []int64{1, 3, 79, 72, 27, 57, 5, 4, 30, 61, 60, 9, 67, 40}
		i2 := []int64{12, 75, 74, 54, 34, 32, 23, 43, 68}

		gotInt64Pure := slices.Sort(intersectInt64Pure(i0, i1, i2))

		gotInt64Lib := slices.Sort(slices.Intersection(i0, i1, i2))

		assertEqual(t, gotInt64Lib, gotInt64Pure)
	})

}

func assertEqual(t *testing.T, got, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v; got %v", want, got)
	}
}
