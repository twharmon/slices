package slices_test

import (
	"math/rand"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/twharmon/slices"
)

func assertEqual(t *testing.T, got, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestPush(t *testing.T) {
	s := []int{1}
	got := slices.Append(s, 5)
	want := []int{1, 5}
	assertEqual(t, got, want)
}

func TestContains(t *testing.T) {
	s := []int{1, 2, 3}
	t.Run("true", func(t *testing.T) {
		got := slices.Contains(s, 2)
		want := true
		assertEqual(t, got, want)
	})
	t.Run("false", func(t *testing.T) {
		got := slices.Contains(s, 4)
		want := false
		assertEqual(t, got, want)
	})
}

func TestSplice(t *testing.T) {
	s := []string{"foo", "bar", "baz"}
	t.Run("with inserts", func(t *testing.T) {
		got := slices.Splice(s, 1, 1, "boo")
		want := []string{"foo", "boo", "baz"}
		assertEqual(t, got, want)
	})
	t.Run("no inserts", func(t *testing.T) {
		got := slices.Splice(s, 1, 1)
		want := []string{"foo", "baz"}
		assertEqual(t, got, want)
	})
}

func TestReverse(t *testing.T) {
	t.Run("even", func(t *testing.T) {
		s := []string{"foo", "bar"}
		got := slices.Reverse(s)
		want := []string{"bar", "foo"}
		assertEqual(t, got, want)
	})
	t.Run("odd", func(t *testing.T) {
		s := []string{"foo", "bar", "baz"}
		got := slices.Reverse(s)
		want := []string{"baz", "bar", "foo"}
		assertEqual(t, got, want)
	})
}

func TestUnshift(t *testing.T) {
	s := []string{"foo"}
	got := slices.Unshift(s, "bar")
	want := []string{"bar", "foo"}
	assertEqual(t, got, want)
}

func TestFind(t *testing.T) {
	s := []string{"foo"}
	t.Run("found", func(t *testing.T) {
		got := slices.Find(s, func(item string) bool {
			return item == "foo"
		})
		want := "foo"
		assertEqual(t, got, want)
	})
	t.Run("not found", func(t *testing.T) {
		got := slices.Find(s, func(item string) bool {
			return item == "bar"
		})
		want := ""
		assertEqual(t, got, want)
	})
}

func TestIndexOf(t *testing.T) {
	s := []string{"foo"}
	t.Run("found", func(t *testing.T) {
		got := slices.IndexOf(s, "foo")
		want := 0
		assertEqual(t, got, want)
	})
	t.Run("not found", func(t *testing.T) {
		got := slices.IndexOf(s, "bar")
		want := -1
		assertEqual(t, got, want)
	})
}

func TestIndexOfFunc(t *testing.T) {
	s := []string{"foo"}
	t.Run("found", func(t *testing.T) {
		got := slices.IndexOfFunc(s, func(item string) bool { return item == "foo" })
		want := 0
		assertEqual(t, got, want)
	})
	t.Run("not found", func(t *testing.T) {
		got := slices.IndexOfFunc(s, func(item string) bool { return item == "bar" })
		want := -1
		assertEqual(t, got, want)
	})
}

func TestFilter(t *testing.T) {
	s := []string{"foo", "bar"}
	got := slices.Filter(s, func(item string) bool {
		return strings.HasPrefix(item, "f")
	})
	assertEqual(t, len(got), 1)
}

func TestMax(t *testing.T) {
	t.Run("common", func(t *testing.T) {
		s := []int{2, 6, 1, 4, 3}
		got := slices.Max(s)
		want := 6
		assertEqual(t, got, want)
	})
	t.Run("zero value", func(t *testing.T) {
		s := []string{}
		got := slices.Max(s)
		want := ""
		assertEqual(t, got, want)
	})
}

func TestMin(t *testing.T) {
	t.Run("common", func(t *testing.T) {
		s := []int{2, 6, 1, 4, 3}
		got := slices.Min(s)
		want := 1
		assertEqual(t, got, want)
	})
	t.Run("zero value", func(t *testing.T) {
		s := []string{}
		got := slices.Min(s)
		want := ""
		assertEqual(t, got, want)
	})
}

func TestMaxFunc(t *testing.T) {
	t.Run("common", func(t *testing.T) {
		s := []int{2, 6, 1, 4, 3}
		got := slices.MaxFunc(s, func(a, b int) bool { return a < b })
		want := 6
		assertEqual(t, got, want)
	})
	t.Run("zero value", func(t *testing.T) {
		s := []int{}
		got := slices.MaxFunc(s, func(a, b int) bool { return a < b })
		want := 0
		assertEqual(t, got, want)
	})
}

func TestMinFunc(t *testing.T) {
	t.Run("common", func(t *testing.T) {
		s := []int{2, 6, 1, 4, 3}
		got := slices.MinFunc(s, func(a, b int) bool { return a < b })
		want := 1
		assertEqual(t, got, want)
	})
	t.Run("zero value", func(t *testing.T) {
		s := []int{}
		got := slices.MinFunc(s, func(a, b int) bool { return a < b })
		want := 0
		assertEqual(t, got, want)
	})
}

func TestSome(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		s := []string{"foo", "bar", "baz"}
		if !slices.Some(s, func(s string) bool { return s == "bar" }) {
			t.Fatalf("some was false; expected true")
		}
	})
	t.Run("false", func(t *testing.T) {
		s := []string{"foo", "bar", "baz"}
		if slices.Some(s, func(s string) bool { return s == "x" }) {
			t.Fatalf("some was true; expected false")
		}
	})
}

func TestEvery(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		s := []string{"foo", "bar", "baz"}
		if !slices.Every(s, func(s string) bool { return len(s) == 3 }) {
			t.Fatalf("every was true; expected false")
		}
	})
	t.Run("false", func(t *testing.T) {
		s := []string{"foo", "bar", "baz"}
		if slices.Every(s, func(s string) bool { return len(s) < 2 }) {
			t.Fatalf("every was true; expected false")
		}
	})
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

func TestSortFunc(t *testing.T) {
	var s []int
	for i := 0; i < 1000; i++ {
		s = append(s, rand.Int())
	}
	got := slices.SortFunc(s, func(a int, b int) bool {
		return a < b
	})
	want := slices.Clone(s)
	sort.Slice(want, func(i, j int) bool {
		return want[i] < want[j]
	})
	assertEqual(t, got, want)
}

func TestSort(t *testing.T) {
	var s []int
	for i := 0; i < 1000; i++ {
		s = append(s, rand.Int())
	}
	got := slices.Sort(s)
	want := slices.Clone(s)
	sort.Ints(want)
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
	t.Run("empty result", func(t *testing.T) {
		a := []string{"foo", "bar"}
		b := []string{"baz"}
		want := []string{}
		got := slices.Intersection(a, b)
		assertEqual(t, got, want)
	})
	t.Run("dupes in slice", func(t *testing.T) {
		a := []string{"foo", "bar", "baz", "foo"}
		b := []string{"baz", "baz"}
		c := []string{"foo", "foo", "foo"}
		want := []string{}
		got := slices.Sort(slices.Intersection(a, b, c))
		assertEqual(t, got, want)
	})
}

func TestUnion(t *testing.T) {
	t.Run("common", func(t *testing.T) {
		a := []string{"foo", "bar"}
		b := []string{"bar", "baz"}
		want := slices.Sort([]string{"foo", "bar", "baz"})
		got := slices.Sort(slices.Union(a, b))
		assertEqual(t, got, want)
	})
	t.Run("empty slice", func(t *testing.T) {
		var s [][]string
		got := slices.Union(s...)
		assertEqual(t, got, []string{})
	})
}
