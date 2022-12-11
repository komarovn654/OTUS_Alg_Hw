package hw07_heapsort

func (s *Array) SelctionSort() {
	for i := len(s.Ar) - 1; i > 1; i-- {
		s.swap(s.findMax(0, i), i)
	}
}
