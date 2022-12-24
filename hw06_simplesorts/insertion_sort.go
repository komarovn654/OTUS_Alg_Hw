package hw06simplesorts

import (
	"time"

	"github.com/komarovn654/OTUS_Alg_Hw/sortutils"
)

func InsertionSort(array sortutils.Array) <-chan sortutils.SortTime {
	sTime := make(chan sortutils.SortTime)

	go func() {
		start := time.Now()
		for i := 1; i < len(array.Ar); i++ {
			for j := i; j > 0 && (array.Ar)[j] < (array.Ar)[j-1]; j-- {
				array.Swap(j, j-1)
			}
		}

		sTime <- sortutils.SortTime{Time: time.Since(start)}
	}()

	return sTime
}

func InsertionSortShift(array sortutils.Array) <-chan sortutils.SortTime {
	sTime := make(chan sortutils.SortTime)

	go func() {
		start := time.Now()
		for i := 1; i < len(array.Ar); i++ {
			j := i
			cache := array.Ar[j]
			for ; j > 0 && cache < array.Ar[j-1]; j-- {
				array.Ar[j] = array.Ar[j-1]
			}
			array.Ar[j] = cache
		}

		sTime <- sortutils.SortTime{Time: time.Since(start)}
	}()

	return sTime
}

func InsertionSortBinarySearch(array sortutils.Array) <-chan sortutils.SortTime {
	sTime := make(chan sortutils.SortTime)

	go func() {
		start := time.Now()
		for i := 1; i < len(array.Ar); i++ {
			j := i
			cache := array.Ar[j]
			pos := array.BinarySearch(cache, 0, j)
			for ; j > pos && cache < array.Ar[j-1]; j-- {
				array.Ar[j] = array.Ar[j-1]
			}
			array.Ar[j] = cache
		}

		sTime <- sortutils.SortTime{Time: time.Since(start)}
	}()

	return sTime
}
