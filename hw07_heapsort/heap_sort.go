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
