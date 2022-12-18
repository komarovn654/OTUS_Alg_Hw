package hw07_heapsort

import (
	"time"

	"github.com/komarovn654/OTUS_Alg_Hw/sortutils"
)

func SelctionSort(array sortutils.Array) <-chan sortutils.SortTime {
	sTime := make(chan sortutils.SortTime)

	go func() {
		start := time.Now()
		for i := len(array.Ar) - 1; i > 0; i-- {
			array.Swap(array.FindMax(0, i), i)
		}
		sTime <- sortutils.SortTime{Time: time.Since(start)}
	}()

	return sTime
}
