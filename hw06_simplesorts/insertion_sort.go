package hw06simplesorts

import (
	"fmt"
	"time"
)

func InsertionSort(array *[]Item) <-chan SortTime {
	sTime := make(chan SortTime)

	go func() {
		start := time.Now()
		for i := 1; i < len(*array); i++ {
			for j := i; j > 0 && (*array)[j] < (*array)[j-1]; j-- {
				(*array)[j], (*array)[j-1] = swap((*array)[j], (*array)[j-1])
			}
		}

		sTime <- SortTime{Time: time.Since(start)}
	}()

	return sTime
}

func InsertionSortShift(array *[]Item) <-chan SortTime {
	sTime := make(chan SortTime)

	go func() {
		start := time.Now()
		for i := 1; i < len(*array); i++ {
			j := i
			cache := (*array)[j]
			for ; j > 0 && cache < (*array)[j-1]; j-- {
				(*array)[j] = (*array)[j-1]
			}
			(*array)[j] = cache
		}

		sTime <- SortTime{Time: time.Since(start)}
	}()

	return sTime
}

func InsertionSortBinarySearch(array *[]Item) <-chan SortTime {
	sTime := make(chan SortTime)

	go func() {
		start := time.Now()
		for i := 1; i < len(*array); i++ {
			j := i
			cache := (*array)[j]
			pos := binarySearch(array, 0, 0, 0)
			for ; j > pos && cache < (*array)[j-1]; j-- {
				(*array)[j] = (*array)[j-1]
			}
			(*array)[j] = cache
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
	fmt.Println(mid)
	if (*array)[mid] < key {
		return binarySearch(array, key, mid, max)
	}
	return binarySearch(array, key, min, mid)
}
