package slice

func Map[T any, K any] (s *Slice[T], f func(item T) K) *Slice[K] {
	k := New[K]()
	for i := range s.s {
		k.Push(f(s.s[i]))
	}
	return k
}

func Reduce[T any, K any] (s *Slice[T], f func(prev K, cur T) K) K {
	var k K
	for i := range s.s {
		k = f(k, s.s[i])
	}
	return k
}
