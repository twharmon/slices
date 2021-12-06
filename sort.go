package slice

func (s *Slice[T]) partition(low, high int, less func(a T, b T) bool) int {
	for j := low; j < high; j++ {
		if less(s.s[j], s.s[high]) {
			s.s[low], s.s[j] = s.s[j], s.s[low]
			low++
		}
	}
	s.s[low], s.s[high] = s.s[high], s.s[low]
	return low
}

func (s *Slice[T]) quickSort(low, high int, less func(a T, b T) bool) {
	if low < high {
		if high-low < 12 {
			s.insertionSort(low, high, less)
		} else {
			p := s.partition(low, high, less)
			s.quickSort(low, p-1, less)
			s.quickSort(p+1, high, less)
		}
	}
}

func (s *Slice[T]) insertionSort(a, b int, less func(a T, b T) bool) {
	for i := 1; i < b-a+1; i++ {
		j := i
		for j > 0 {
			if less(s.s[a+j], s.s[a+j-1]) {
				s.s[a+j-1], s.s[a+j] = s.s[a+j], s.s[a+j-1]
			}
			j--
		}
	}
}
