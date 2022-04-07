package slices

func Clone[T any](s []T) []T {
	var c []T
	for i := range s {
		c = Append(c, s[i])
	}
	return c
}

func Append[T any](s []T, item ...T) []T {
	return append(s, item...)
}

func Reverse[T any](s []T) []T {
	var res []T
	for i := 0; i < len(s); i++ {
		res = Append(res, s[len(s)-1-i])
	}
	return res
}

func Splice[T any](s []T, index int, cnt int, item ...T) []T {
	tail := Unshift(s[index:][cnt:], item...)
	s = s[:index]
	return Concat(s, tail)
}

func Concat[T any](s []T, other []T) []T {
	return Append(s, other...)
}

func Unshift[T any](s []T, item ...T) []T {
	return append(item, s...)
}

func Find[T any](s []T, f func(item T) bool) T {
	for i := range s {
		if f(s[i]) {
			return s[i]
		}
	}
	var t T
	return t
}

func IndexOf[T any](s []T, f func(item T) bool) int {
	for i := range s {
		if f(s[i]) {
			return i
		}
	}
	return -1
}

func Some[T any](s []T, f func(item T) bool) bool {
	return IndexOf(s, f) >= 0
}

func Contains[T Ordered](s []T, item T) bool {
	for i := range s {
		if s[i] == item {
			return true
		}
	}
	return false
}

func Max[T Ordered](s []T) T {
	if len(s) == 0 {
		var t T
		return t
	}
	max := s[0]
	for i := range s {
		if s[i] > max {
			max = s[i]
		}
	}
	return max
}

func Min[T Ordered](s []T) T {
	if len(s) == 0 {
		var t T
		return t
	}
	min := s[0]
	for i := range s {
		if s[i] < min {
			min = s[i]
		}
	}
	return min
}

func MaxFunc[T Ordered](s []T, less func(T, T) bool) T {
	if len(s) == 0 {
		var t T
		return t
	}
	max := s[0]
	for i := range s {
		if less(max, s[i]) {
			max = s[i]
		}
	}
	return max
}

func MinFunc[T Ordered](s []T, less func(T, T) bool) T {
	if len(s) == 0 {
		var t T
		return t
	}
	min := s[0]
	for i := range s {
		if less(s[i], min) {
			min = s[i]
		}
	}
	return min
}

func Every[T any](s []T, f func(item T) bool) bool {
	for i := range s {
		if !f(s[i]) {
			return false
		}
	}
	return true
}

func SortFunc[T any](s []T, less func(a T, b T) bool) []T {
	c := Clone(s)
	quickSortFunc(c, 0, len(c)-1, less)
	return c
}

type Ordered interface {
	int | int32 | int16 | int8 | int64 | uint | uint32 | uint16 | uint8 | uint64 | float32 | float64 | string
}

func Sort[T Ordered](s []T) []T {
	c := Clone(s)
	quickSort(c, 0, len(s)-1)
	return c
}

func Filter[T any](s []T, f func(item T) bool) []T {
	n := make([]T, 0, len(s))
	for i := range s {
		if f(s[i]) {
			n = Append(n, s[i])
		}
	}
	return n
}

func Map[T any, K any](s []T, f func(item T) K) []K {
	var k []K
	for i := range s {
		k = Append(k, f(s[i]))
	}
	return k
}

func Reduce[T any, K any](s []T, f func(prev K, cur T) K) K {
	var k K
	for i := range s {
		k = f(k, s[i])
	}
	return k
}
