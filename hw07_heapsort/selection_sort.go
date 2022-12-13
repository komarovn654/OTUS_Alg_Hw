package hw07_heapsort

import "time"

func (s *Array) SelctionSort() <-chan SortTime {
	sTime := make(chan SortTime)

	go func() {
		start := time.Now()
		for i := len(s.Ar) - 1; i > 0; i-- {
			s.swap(s.findMax(0, i), i)
		}
		sTime <- SortTime{Time: time.Since(start)}
	}()

	return sTime
}
