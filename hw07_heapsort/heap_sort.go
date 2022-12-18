package hw07_heapsort

import (
	"time"

	"github.com/komarovn654/OTUS_Alg_Hw/sortutils"
)

func HeapSort(array sortutils.Array) <-chan sortutils.SortTime {
	sTime := make(chan sortutils.SortTime)

	go func() {
		start := time.Now()

		for h := len(array.Ar) / 2; h >= 0; h-- {
			array.Heapify(h, len(array.Ar))
		}
		for i := len(array.Ar) - 1; i > 0; i-- {
			array.Swap(0, i)
			array.Heapify(0, i)
		}

		sTime <- sortutils.SortTime{Time: time.Since(start)}
	}()

	return sTime
}
