package hw06simplesorts

import (
	"time"

	"github.com/komarovn654/OTUS_Alg_Hw/sortutils"
)

func BubbleSort(array sortutils.Array) <-chan sortutils.SortTime {
	sTime := make(chan sortutils.SortTime)

	go func() {
		start := time.Now()
		for i := len(array.Ar) - 1; i > 0; i-- {
			for j := 0; j < i; j++ {
				if array.Ar[j+1] < array.Ar[j] {
					array.Swap(j, j+1)
				}
			}
		}
		sTime <- sortutils.SortTime{Time: time.Since(start)}
	}()

	return sTime
}

func BubbleSortOpt(array sortutils.Array) <-chan sortutils.SortTime {
	sTime := make(chan sortutils.SortTime)

	go func() {
		start := time.Now()
		for i := len(array.Ar) - 1; i > 0; i-- {
			sorted := true
			for j := 0; j < i; j++ {
				if array.Ar[j+1] < array.Ar[j] {
					array.Swap(j, j+1)
					sorted = false
				}
			}
			if sorted {
				break
			}
		}
		sTime <- sortutils.SortTime{Time: time.Since(start)}
	}()

	return sTime
}
