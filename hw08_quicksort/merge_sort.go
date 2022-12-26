package hw08quicksort

import (
	"time"

	"github.com/komarovn654/OTUS_Alg_Hw/sortutils"
)

func (s *Sort) MergeSort() <-chan sortutils.SortTime {
	sTime := make(chan sortutils.SortTime)

	go func() {
		start := time.Now()

		s.mergeSort(0, len(s.Array.Ar)-1)

		sTime <- sortutils.SortTime{Time: time.Since(start)}
	}()

	return sTime
}

func (s *Sort) mergeSort(lIndex int, rIndex int) {
	if lIndex >= rIndex {
		return
	}

	mIndex := (lIndex + rIndex) / 2
	s.mergeSort(lIndex, mIndex)
	s.mergeSort(mIndex+1, rIndex)
	s.merge(lIndex, rIndex, mIndex)
}

func (s *Sort) merge(lIndex int, rIndex int, mIndex int) {
	tmp := make([]sortutils.Item, rIndex-lIndex+1)
	m := 0
	a := lIndex
	b := mIndex + 1

	for a <= mIndex && b <= rIndex {
		if s.Array.Ar[a] < s.Array.Ar[b] {
			tmp[m] = s.Array.Ar[a]
			a++
		} else {
			tmp[m] = s.Array.Ar[b]
			b++
		}
		m++
	}

	for b <= rIndex {
		tmp[m] = s.Array.Ar[b]
		m++
		b++
	}
	for a <= mIndex {
		tmp[m] = s.Array.Ar[a]
		m++
		a++
	}

	for i := 0; i < len(tmp); i++ {
		s.Array.Ar[lIndex+i] = tmp[i]
	}
}
