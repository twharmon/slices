package slice_test

import (
	"sort"
	"strings"
	"testing"

	"github.com/twharmon/slice"
)

func TestLen(t *testing.T) {
	s := slice.New("foo")
	if s.Len() != 1 {
		t.Fatalf("wrong len %d", s.Len())
	}
}

func TestClear(t *testing.T) {
	s := slice.New("foo")
	s.Clear()
	if s.Len() != 0 {
		t.Fatalf("wrong len %d", s.Len())
	}
}

func TestPop(t *testing.T) {
	s := slice.New("foo")
	i := s.Pop()
	if s.Len() != 0 {
		t.Fatalf("wrong len %d", s.Len())
	}
	if i != "foo" {
		t.Fatalf("wrong item %s", i)
	}
}

func TestJoin(t *testing.T) {
	s := slice.New("foo", "bar", "baz")
	if s.Join(", ") != "foo, bar, baz" {
		t.Fatalf("wrong join %s", s.Join(", "))
	}
}

func TestShift(t *testing.T) {
	s := slice.New("foo", "bar", "baz")
	i := s.Shift()
	if s.Len() != 2 {
		t.Fatalf("wrong len %d", s.Len())
	}
	if i != "foo" {
		t.Fatalf("wrong item %s", i)
	}
	if *s.At(0) != "bar" {
		t.Fatalf("wrong val %s", *s.At(0))
	}
	if *s.At(1) != "baz" {
		t.Fatalf("wrong val %s", *s.At(1))
	}
}

func TestSpliceNoInserts(t *testing.T) {
	s := slice.New("foo", "bar", "baz")
	s.Splice(1, 1)
	if s.Len() != 2 {
		t.Fatalf("wrong len %d", s.Len())
	}
	if *s.At(0) != "foo" {
		t.Fatalf("wrong val %s", *s.At(0))
	}
	if *s.At(1) != "baz" {
		t.Fatalf("wrong val %s", *s.At(1))
	}
}

func TestSpliceWithInserts(t *testing.T) {
	s := slice.New("foo", "bar", "baz")
	s.Splice(1, 1, "boo")
	if s.Len() != 3 {
		t.Fatalf("wrong len %d", s.Len())
	}
	if *s.At(0) != "foo" {
		t.Fatalf("wrong val %s", *s.At(0))
	}
	if *s.At(1) != "boo" {
		t.Fatalf("wrong val %s", *s.At(1))
	}
	if *s.At(2) != "baz" {
		t.Fatalf("wrong val %s", *s.At(2))
	}
}

func TestReverseEven(t *testing.T) {
	s := slice.New("foo", "bar")
	s.Reverse()
	if *s.At(0) != "bar" {
		t.Fatalf("wrong val %s", *s.At(0))
	}
	if *s.At(1) != "foo" {
		t.Fatalf("wrong val %s", *s.At(1))
	}
}

func TestReverseOdd(t *testing.T) {
	s := slice.New("foo", "bar", "baz")
	s.Reverse()
	if *s.At(0) != "baz" {
		t.Fatalf("wrong val %s", *s.At(0))
	}
	if *s.At(1) != "bar" {
		t.Fatalf("wrong val %s", *s.At(1))
	}
	if *s.At(2) != "foo" {
		t.Fatalf("wrong val %s", *s.At(1))
	}
}

func TestUnshift(t *testing.T) {
	s := slice.New("foo")
	s.Unshift("bar")
	if s.Len() != 2 {
		t.Fatalf("wrong len %d", s.Len())
	}
	if *s.At(0) != "bar" {
		t.Fatalf("wrong val %s", *s.At(0))
	}
	if *s.At(1) != "foo" {
		t.Fatalf("wrong val %s", *s.At(1))
	}
}

func TestFindFound(t *testing.T) {
	s := slice.New("foo")
	found := s.Find(func(item string) bool {
		return item == "foo"
	})
	if *found != "foo" {
		t.Fatalf("wrong find %s", *found)
	}
}

func TestFindNotFound(t *testing.T) {
	s := slice.New("foo")
	found := s.Find(func(item string) bool {
		return item == "bar"
	})
	if found != nil {
		t.Fatalf("expected nil; got %s", *found)
	}
}

func TestFilter(t *testing.T) {
	s := slice.New("foo", "bar")
	ns := s.Filter(func(item string) bool {
		return strings.HasPrefix(item, "f")
	})
	if ns.Len() != 1 {
		t.Fatalf("wrong len %d", ns.Len())
	}
}

func TestSlice(t *testing.T) {
	s := slice.New("foo", "bar", "baz")
	n := s.Slice(0, 1)
	if n.String() != "[foo]" {
		t.Fatalf("wrong slice %s", n)
	}
}

func TestSomeTrue(t *testing.T) {
	s := slice.New("foo", "bar", "baz")
	if !s.Some(func(s string) bool { return s == "bar" }) {
		t.Fatalf("some was false; expected true")
	}
}

func TestSomeFalse(t *testing.T) {
	s := slice.New("foo", "bar", "baz")
	if s.Some(func(s string) bool { return s == "x" }) {
		t.Fatalf("some was true; expected false")
	}
}

func TestEveryFalse(t *testing.T) {
	s := slice.New("foo", "bar", "baz")
	if s.Every(func(s string) bool { return len(s) < 2 }) {
		t.Fatalf("every was true; expected false")
	}
}

func TestEveryTrue(t *testing.T) {
	s := slice.New("foo", "bar", "baz")
	if !s.Every(func(s string) bool { return len(s) == 3 }) {
		t.Fatalf("every was true; expected false")
	}
}

func TestAtNil(t *testing.T) {
	s := slice.New("foo", "bar", "baz")
	if s.At(4) != nil {
		t.Fatalf("at not nil; found %s", *s.At(4))
	}
}

func TestCap(t *testing.T) {
	s := slice.New("foo", "bar", "baz")
	if s.Cap() != 3 {
		t.Fatalf("cap not 3; found %d", s.Cap())
	}
}

func TestSort(t *testing.T) {
	s := slice.New("foo", "bar", "baz", "lorem", "ipsum", "donor", "sit", "alpha", "beta", "delta", "gamma", "omega", "epsilon", "mu", "nu", "lambda", "upsilon", "zeta", "eta", "rho", "psi", "iota", "apple", "banana", "pomegranite", "orange", "kiwi", "carrot", "brocoli")
	s.Sort(func(a string, b string) bool {
		return a < b
	})
	s2 := []string{"foo", "bar", "baz", "lorem", "ipsum", "donor", "sit", "alpha", "beta", "delta", "gamma", "omega", "epsilon", "mu", "nu", "lambda", "upsilon", "zeta", "eta", "rho", "psi", "iota", "apple", "banana", "pomegranite", "orange", "kiwi", "carrot", "brocoli"}
	sort.Slice(s2, func(i, j int) bool {
		return s2[i] < s2[j]
	})
	for i := 0; i < len(s2); i++ {
		if *s.At(i) != s2[i] {
			t.Fatalf("wrong order %s", s)
		}
	}
}

func TestFmt(t *testing.T) {
	s := slice.New("foo", "bar")
	if s.String() != "[foo bar]" {
		t.Fatalf("wrong fmt %s", s.String())
	}
}

func shortSlice() []string {
	return []string{"foo", "bar", "baz", "lorem", "ipsum", "donor", "sit", "alpha", "beta", "delta", "gamma", "omega", "epsilon", "mu", "nu"}
}

func BenchmarkStdLibSort(b *testing.B) {
	s := shortSlice()
	for i := 0; i < b.N; i++ {
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sort.Slice(s, func(i, j int) bool {
			return len(s[i]) < len(s[j])
		})
	}
}

func BenchmarkSortShort(b *testing.B) {
	s := slice.New(shortSlice()...)
	for i := 0; i < b.N; i++ {
		s.Sort(func(a string, b string) bool {
			return a < b
		})
		s.Sort(func(a string, b string) bool {
			return len(a) < len(b)
		})
	}
}

func longSlice() []string {
	return []string{"foo", "bar", "baz", "lorem", "ipsum", "donor", "sit", "alpha", "beta", "delta", "gamma", "omega", "epsilon", "mu", "nu", "lambda", "upsilon", "zeta", "eta", "rho", "psi", "iota", "apple", "banana", "pomegranite", "orange", "kiwi", "carrot", "brocoli", "dog", "cat", "leopard", "bull", "pig", "zebra", "hippo", "rhinocerous", "deer", "elk", "moose", "duck", "rabbit", "snake", "sloth", "aardvark", "monkey", "armadillo", "gorilla", "chimapnzee", "ape", "bird"}
}

func BenchmarkStdLibSortLong(b *testing.B) {
	s := longSlice()
	for i := 0; i < b.N; i++ {
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sort.Slice(s, func(i, j int) bool {
			return len(s[i]) < len(s[j])
		})
	}
}

func BenchmarkSortLong(b *testing.B) {
	s := slice.New(longSlice()...)
	for i := 0; i < b.N; i++ {
		s.Sort(func(a string, b string) bool {
			return a < b
		})
		s.Sort(func(a string, b string) bool {
			return len(a) < len(b)
		})
	}
}
