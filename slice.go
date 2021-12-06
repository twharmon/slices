package slice

import (
	"fmt"
	"strings"
)

type Slice[T any] struct {
	s []T
}

func New[T any](items ...T) *Slice[T] {
	return &Slice[T]{
		s: items,
	}
}

func (s *Slice[T]) Push(item ...T) {
	s.s = append(s.s, item...)
}

func (s *Slice[T]) Reverse() {
	for i := 0; i < len(s.s)/2; i++ {
		s.s[i], s.s[len(s.s)-1-i] = s.s[len(s.s)-1-i], s.s[i]
	}
}

func (s *Slice[T]) Splice(index int, cnt int, item ...T) {
	tail := s.Slice(index, s.Len())
	tail.s = tail.s[cnt:]
	tail.Unshift(item...)
	s.s = s.s[:index]
	s.Concat(tail)
}

func (s *Slice[T]) Concat(other *Slice[T]) {
	s.Push(other.s...)
}

func (s *Slice[T]) Pop() T {
	item := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return item
}

func (s *Slice[T]) Join(sep string) string {
	var b strings.Builder
	for i := 0; i < len(s.s); i++ {
		b.WriteString(fmt.Sprintf("%v", s.s[i]))
		if i < len(s.s)-1 {
			b.WriteString(sep)
		}
	}
	return b.String()
}

func (s *Slice[T]) Shift() T {
	item := s.s[0]
	for i := 0; i < len(s.s)-1; i++ {
		s.s[i] = s.s[i+1]
	}
	s.s = s.s[:len(s.s)-1]
	return item
}

func (s *Slice[T]) Unshift(item ...T) {
	s.s = append(item, s.s...)
}

func (s *Slice[T]) Find(f func(item T) bool) *T {
	for i := range s.s {
		if f(s.s[i]) {
			return &s.s[i]
		}
	}
	return nil
}

func (s *Slice[T]) IndexOf(f func(item T) bool) int {
	for i := range s.s {
		if f(s.s[i]) {
			return i
		}
	}
	return -1
}

func (s *Slice[T]) Some(f func(item T) bool) bool {
	return s.IndexOf(f) >= 0
}

func (s *Slice[T]) Every(f func(item T) bool) bool {
	for i := range s.s {
		if !f(s.s[i]) {
			return false
		}
	}
	return true
}

func (s *Slice[T]) Clear() {
	s.s = nil
}

func (s *Slice[T]) Slice(start int, end int) *Slice[T] {
	n := New[T]()
	n.s = s.s[start:end]
	return n
}

func (s *Slice[T]) Sort(less func(a T, b T) bool) {
	s.quickSort(0, len(s.s)-1, less)
}

func (s *Slice[T]) At(index int) *T {
	if index >= s.Len() {
		return nil
	}
	return &s.s[index]
}

func (s *Slice[T]) Filter(f func(item T) bool) *Slice[T] {
	n := New[T]()
	for i := range s.s {
		if f(s.s[i]) {
			n.Push(s.s[i])
		}
	}
	return n
}

func (s *Slice[T]) Len() int {
	return len(s.s)
}

func (s *Slice[T]) Cap() int {
	return cap(s.s)
}

func (s *Slice[T]) String() string {
	return fmt.Sprintf("%v", s.s)
}
