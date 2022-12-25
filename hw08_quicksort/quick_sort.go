package hw08quicksort

import (
	"time"

	"github.com/komarovn654/OTUS_Alg_Hw/sortutils"
)

type Sort struct {
	Ar sortutils.Array
}

func (s *Sort) QuickSort() <-chan sortutils.SortTime {
	sTime := make(chan sortutils.SortTime)

	go func() {
		start := time.Now()

		s.quickSort(0, len(s.Ar.Ar)-1)

		sTime <- sortutils.SortTime{Time: time.Since(start)}
	}()

	return sTime
}

func (s *Sort) quickSort(lIndex int, rIndex int) {
	if lIndex >= rIndex {
		return
	}

	mIndex := s.partition(lIndex, rIndex)
	s.quickSort(lIndex, mIndex-1)
	s.quickSort(mIndex+1, rIndex)
}

func (s *Sort) partition(lIndex int, rIndex int) int {
	sPtr := lIndex - 1
	bPtr := 0
	for i := lIndex; i <= rIndex; i++ {
		if s.Ar.Ar[i] <= s.Ar.Ar[rIndex] {
			sPtr++
			s.Ar.Swap(sPtr, i)
		}
		bPtr++
	}
	return sPtr
}
