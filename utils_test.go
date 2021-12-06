package slice_test

import (
	"testing"

	"github.com/twharmon/slice"
)

func TestMap(t *testing.T) {
	s := slice.New("f", "ba", "baz")
	n := slice.Map(s, func(i string) int {
		return len(i)
	})
	if n.Len() != 3 {
		t.Fatalf("%d != %d", n.Len(), 3)
	}
	if *n.At(0) != 1 {
		t.Fatalf("%d != %d", *n.At(0), 1)
	}
	if *n.At(1) != 2 {
		t.Fatalf("%d != %d", *n.At(1), 2)
	}
	if *n.At(2) != 3 {
		t.Fatalf("%d != %d", *n.At(2), 3)
	}
}

func TestReduce(t *testing.T) {
	s := slice.New("f", "ba", "baz")
	cnt := slice.Reduce(s, func(cnt int, i string) int {
		return cnt + len(i)
	})
	if cnt != 6 {
		t.Fatalf("%d != %d", cnt, 6)
	}
}
