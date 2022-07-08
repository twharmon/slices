package slices

type Ordered interface {
	int | int32 | int16 | int8 | int64 | uint | uint32 | uint16 | uint8 | uint64 | float32 | float64 | string
}

// Clone creates a clone slice and returns it.
func Clone[T any](s []T) []T {
	c := make([]T, len(s))
	for i := range s {
		c[i] = s[i]
	}
	return c
}

// Append appends an item to the slice and returns the new slice. The
// given slice is not changed.
func Append[T any](s []T, item ...T) []T {
	return append(Clone(s), item...)
}

// Reverse creates a slice that is the reverse of the provided slice
// and returns it. The given slice is not changed.
func Reverse[T any](s []T) []T {
	res := make([]T, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = s[len(s)-1-i]
	}
	return res
}

// Splice creates a new slice that is spliced by removing or
// replacing existing elements and/or adding new elements in place.
// The given slice is not changed.
func Splice[T any](s []T, index int, cnt int, item ...T) []T {
	tail := Unshift(s[index:][cnt:], item...)
	s = s[:index]
	return Concat(s, tail)
}

// Concat combines the contents of all the given slices. The given
// slices are not changed.
func Concat[T any](s ...[]T) []T {
	totalLen := 0
	for i := range s {
		totalLen += len(s[i])
	}

	output := make([]T, 0, totalLen)
	for i := range s {
		output = append(output, s[i]...)
	}

	return output
}

// Unshift creates a new slice and prepends all the given items and
// returns it. The given slice is not changed.
func Unshift[T any](s []T, item ...T) []T {
	return append(item, s...)
}

// Find finds an item in the given slice that satisfies the given
// test function.
func Find[T any](s []T, test func(item T) bool) T {
	for i := range s {
		if test(s[i]) {
			return s[i]
		}
	}
	var t T
	return t
}

// IndexOf finds the index of the first item in the given slice that
// satisfies the given test function.
func IndexOf[T Ordered](s []T, item T) int {
	for i := range s {
		if s[i] == item {
			return i
		}
	}
	return -1
}

// IndexOfFunc finds the index of the first item in the given slice
// that satisfies the given test function.
func IndexOfFunc[T any](s []T, test func(item T) bool) int {
	for i := range s {
		if test(s[i]) {
			return i
		}
	}
	return -1
}

// Some checks is any of the items in the given slice satisfies the
// given test function.
func Some[T any](s []T, test func(item T) bool) bool {
	return IndexOfFunc(s, test) >= 0
}

// Contains checks if any of the items in the given slice are equal
// to the given item.
func Contains[T Ordered](s []T, item T) bool {
	for i := range s {
		if s[i] == item {
			return true
		}
	}
	return false
}

// Max returns the max item in the given slice.
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

// Min returns the min item in the given slice.
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

// MaxFunc returns the max item in the given slice according to the
// given less func.
func MaxFunc[T any](s []T, less func(T, T) bool) T {
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

// MinFunc returns the min item in the given slice according to the
// given less func.
func MinFunc[T any](s []T, less func(T, T) bool) T {
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

// Every checks is every item in the given slice satisfies the
// given test function.
func Every[T any](s []T, test func(item T) bool) bool {
	for i := range s {
		if !test(s[i]) {
			return false
		}
	}
	return true
}

// SortFunc creates a new slice that is sorted in ascending order
// according the the given less func and returns it. The given slice
// is not changed.
func SortFunc[T any](s []T, less func(a T, b T) bool) []T {
	c := Clone(s)
	quickSortFunc(c, 0, len(c)-1, less)
	return c
}

// Sort creates a new slice that is sorted in ascending order. The
// given slice is not changed.
func Sort[T Ordered](s []T) []T {
	c := Clone(s)
	quickSort(c, 0, len(s)-1)
	return c
}

// Filter creates a new slice that contains items from the given
// slice that satisfy the given test function and returns it. The
// given slice is not changed.
func Filter[T any](s []T, test func(item T) bool) []T {
	n := make([]T, 0, len(s))
	for i := range s {
		if test(s[i]) {
			n = append(n, s[i])
		}
	}
	return n
}

// Map creates a new slice with items that are mapped to new values
// according to the given m function. The given slice is not
// changed.
func Map[T any, K any](s []T, m func(item T) K) []K {
	k := make([]K, len(s))
	for i := range s {
		k[i] = m(s[i])
	}
	return k
}

// Reduce iterates through the given slice, reducing the items to a
// value according to the given reducer function and returns the
// reduced value. The given slice is not changed.
func Reduce[T any, K any](s []T, f func(prev K, cur T) K) K {
	var k K
	for i := range s {
		k = f(k, s[i])
	}
	return k
}

// Intersection creates a new slice that contains the intersection of
// all the given slices. The given slices are not changed.
func Intersection[T Ordered](s ...[]T) []T {
	if len(s) == 0 {
		return []T{}
	}
	hash := make(map[T]int)
	for i := range s {
		for j := range s[i] {
			if hash[s[i][j]] == i {
				hash[s[i][j]]++
			}
		}
	}
	result := make([]T, 0, len(hash))
	for k := range hash {
		if hash[k] == len(s) {
			result = append(result, k)
		}
	}
	return result
}

// Union creates a new slice that contains the union of all the given
// slices. The given slices are not changed.
func Union[T Ordered](s ...[]T) []T {
	if len(s) == 0 {
		return []T{}
	}
	hash := make(map[T]struct{})
	for i := range s {
		for j := range s[i] {
			hash[s[i][j]] = struct{}{}
		}
	}
	output := make([]T, len(hash))
	i := 0
	for k := range hash {
		output[i] = k
		i++
	}
	return output
}
