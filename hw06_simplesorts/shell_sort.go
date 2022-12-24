package hw06simplesorts

import (
	"math"
	"time"

	"github.com/komarovn654/OTUS_Alg_Hw/sortutils"
)

func ShellSort(array sortutils.Array) <-chan sortutils.SortTime {
	sTime := make(chan sortutils.SortTime)

	go func() {
		start := time.Now()
		for gap := len(array.Ar) / 2; gap > 0; gap /= 2 {
			for i := gap; i < len(array.Ar); i++ {
				for j := i; j >= gap && (array.Ar)[j-gap] > (array.Ar)[j]; j -= gap {
					array.Swap(j, j-gap)
				}
			}
		}

		sTime <- sortutils.SortTime{Time: time.Since(start)}
	}()

	return sTime
}

func ShellSortFrankLazarus(array sortutils.Array) <-chan sortutils.SortTime {
	sTime := make(chan sortutils.SortTime)

	go func() {
		start := time.Now()
		gap := 2*(len(array.Ar)/4) + 1
		for k := 4; gap > 1; k *= 2 {
			gap = (2 * (len(array.Ar) / k)) + 1
			for i := gap; i < len(array.Ar); i++ {
				for j := i; j >= gap && (array.Ar)[j-gap] > (array.Ar)[j]; j -= gap {
					array.Swap(j, j-gap)
				}
			}
		}

		sTime <- sortutils.SortTime{Time: time.Since(start)}
	}()

	return sTime
}

func ShellSortHibbard(array sortutils.Array) <-chan sortutils.SortTime {
	sTime := make(chan sortutils.SortTime)

	go func() {
		start := time.Now()
		stepsNum := getSteps(len(array.Ar))
		for gap := int(math.Pow(2, float64(stepsNum))) - 1; gap > 0; gap /= 2 {
			for i := gap; i < len(array.Ar); i++ {
				for j := i; j >= gap && (array.Ar)[j-gap] > (array.Ar)[j]; j -= gap {
					array.Swap(j, j-gap)
				}
			}
		}

		sTime <- sortutils.SortTime{Time: time.Since(start)}
	}()

	return sTime
}

func getSteps(N int) (count int) {
	for i := 2; N > i; count++ {
		i = i * 2
	}
	return count
}
