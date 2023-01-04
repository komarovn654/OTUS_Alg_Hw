package hw07_heapsort

import (
	"context"
	"time"

	"github.com/komarovn654/OTUS_Alg_Hw/sortutils"
)

func SelectionSort(ctx context.Context, sTime chan<- sortutils.SortTime, array sortutils.Array) {
	start := time.Now()
	for i := len(array.Ar) - 1; i > 0; i-- {
		array.Swap(array.FindMax(0, i), i)

		select {
		case <-ctx.Done():
			sTime <- sortutils.SortTime{Timeout: true}
			return
		default:
		}
	}

	sTime <- sortutils.SortTime{Time: time.Since(start)}
}
