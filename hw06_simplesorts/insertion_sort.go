package hw06simplesorts

import (
	"time"

	"github.com/komarovn654/OTUS_Alg_Hw/sortutils"
)

func InsertionSort(array sortutils.Array) <-chan SortTime {
	sTime := make(chan SortTime)

	go func() {
		start := time.Now()
		for i := 1; i < len(array.Ar); i++ {
			for j := i; j > 0 && (array.Ar)[j] < (array.Ar)[j-1]; j-- {
				array.Swap(j, j-1)
			}
		}

		sTime <- SortTime{Time: time.Since(start)}
	}()

	return sTime
}

func InsertionSortShift(array sortutils.Array) <-chan SortTime {
	sTime := make(chan SortTime)

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

		sTime <- SortTime{Time: time.Since(start)}
	}()

	return sTime
}

func InsertionSortBinarySearch(array sortutils.Array) <-chan SortTime {
	sTime := make(chan SortTime)

	go func() {
		start := time.Now()
		for i := 1; i < len(array.Ar); i++ {
			j := i
			cache := array.Ar[j]
			pos := binarySearch(array, 0, 0, 0)
			for ; j > pos && cache < array.Ar[j-1]; j-- {
				array.Ar[j] = array.Ar[j-1]
			}
			array.Ar[j] = cache
		}

		sTime <- SortTime{Time: time.Since(start)}
	}()

	return sTime
}

func binarySearch(array *[]Item, key Item, min int, max int) int {
	if max <= min {
		if (*array)[min] > key {
			return min
		}
		return min + 1
	}

	mid := (max + min) / 2
	if key < (*array)[mid] {
		return binarySearch(array, key, min, mid-1)
	}
	return binarySearch(array, key, mid+1, max)
}
