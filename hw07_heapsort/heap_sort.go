package hw07_heapsort

import "time"

func (s *Array) HeapSort() <-chan SortTime {
	sTime := make(chan SortTime)

	go func() {
		start := time.Now()

		for h := len(s.Ar) / 2; h >= 0; h-- {
			s.heapify(h, len(s.Ar))
		}
		for i := len(s.Ar) - 1; i > 0; i-- {
			s.swap(0, i)
			s.heapify(0, i)
		}

		sTime <- SortTime{Time: time.Since(start)}
	}()

	return sTime
}

func (s *Array) heapify(rootIndex int, size int) {
	lIndex := 2*rootIndex + 1
	rIndex := 2*rootIndex + 2
	p := rootIndex
	if lIndex < size && s.Ar[lIndex] > s.Ar[p] {
		p = lIndex
	}
	if rIndex < size && s.Ar[rIndex] > s.Ar[p] {
		p = rIndex
	}
	if p == rootIndex {
		return
	}
	s.swap(rootIndex, p)
	s.heapify(p, size)
}
