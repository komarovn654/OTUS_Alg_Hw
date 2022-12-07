package hw06simplesorts

import "time"

func ShellSort(array *[]Item) <-chan SortTime {
	sTime := make(chan SortTime)

	go func() {
		start := time.Now()
		for gap := len(*array) / 2; gap > 0; gap /= 2 {
			for i := gap; i < len(*array); i++ {
				for j := i; j >= gap && (*array)[j-gap] > (*array)[j]; j -= gap {
					(*array)[j], (*array)[j-gap] = swap((*array)[j], (*array)[j-gap])
				}
			}
		}

		sTime <- SortTime{Time: time.Since(start)}
	}()

	return sTime
}

func ShellSortHibbard(array *[]Item) <-chan SortTime {
	sTime := make(chan SortTime)

	go func() {
		start := time.Now()
		for gap := len(*array) / 2; gap > 0; gap /= 2 {
			for i := gap; i < len(*array); i++ {
				for j := i; j >= gap && (*array)[j-gap] > (*array)[j]; j -= gap {
					(*array)[j], (*array)[j-gap] = swap((*array)[j], (*array)[j-gap])
				}
			}
		}

		sTime <- SortTime{Time: time.Since(start)}
	}()

	return sTime
}
