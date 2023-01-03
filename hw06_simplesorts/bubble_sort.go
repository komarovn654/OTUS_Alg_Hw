package hw06simplesorts

import (
	"context"
	"time"

	"github.com/komarovn654/OTUS_Alg_Hw/sortutils"
)

func BubbleSort(ctx context.Context, sTime chan<- sortutils.SortTime, array sortutils.Array) {
	start := time.Now()
	for i := len(array.Ar) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if array.Ar[j+1] < array.Ar[j] {
				array.Swap(j, j+1)
			}

			select {
			case <-ctx.Done():
				sTime <- sortutils.SortTime{Timeout: true}
				return
			default:
			}
		}
	}
	sTime <- sortutils.SortTime{Time: time.Since(start)}
}

func BubbleSortOpt(ctx context.Context, sTime chan<- sortutils.SortTime, array sortutils.Array) {
	start := time.Now()
	for i := len(array.Ar) - 1; i > 0; i-- {
		sorted := true
		for j := 0; j < i; j++ {
			if array.Ar[j+1] < array.Ar[j] {
				array.Swap(j, j+1)
				sorted = false
			}

			select {
			case <-ctx.Done():
				sTime <- sortutils.SortTime{Timeout: true}
				return
			default:
			}
		}
		if sorted {
			break
		}
	}
	sTime <- sortutils.SortTime{Time: time.Since(start)}
}
