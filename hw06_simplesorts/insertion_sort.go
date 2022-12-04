package hw06simplesorts

import "time"

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
