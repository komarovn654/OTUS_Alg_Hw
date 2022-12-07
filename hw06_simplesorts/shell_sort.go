package hw06simplesorts

import (
	"math"
	"time"
)

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

func ShellSortFrankLazarus(array *[]Item) <-chan SortTime {
	sTime := make(chan SortTime)

	go func() {
		start := time.Now()
		gap := 2*(len(*array)/4) + 1
		for k := 4; gap > 1; k *= 2 {
			gap = (2 * (len(*array) / k)) + 1
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
		stepsNum := getSteps(len(*array))
		for gap := int(math.Pow(2, float64(stepsNum))) - 1; gap > 0; gap /= 2 {
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

func getSteps(N int) (count int) {
	for i := 2; N > i; count++ {
		i = i * 2
	}
	return count
}
