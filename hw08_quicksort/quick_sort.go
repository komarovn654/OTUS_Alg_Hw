package hw08quicksort

import (
	"context"
	"time"

	"github.com/komarovn654/OTUS_Alg_Hw/sortutils"
)

func QuickSort(ctx context.Context, sTime chan<- sortutils.SortTime, ar sortutils.Array) {
	start := time.Now()

	quickSort(ctx, ar, 0, len(ar.Ar)-1)

	select {
	case <-ctx.Done():
		sTime <- sortutils.SortTime{Timeout: true}
		return
	default:
	}
	sTime <- sortutils.SortTime{Time: time.Since(start)}
}

func quickSort(ctx context.Context, ar sortutils.Array, lIndex int, rIndex int) {
	select {
	case <-ctx.Done():
		return
	default:
	}

	if lIndex >= rIndex {
		return
	}

	mIndex := partition(ar, lIndex, rIndex)
	quickSort(ctx, ar, lIndex, mIndex-1)
	quickSort(ctx, ar, mIndex+1, rIndex)
}

func partition(ar sortutils.Array, lIndex int, rIndex int) int {
	sPtr := lIndex - 1
	bPtr := 0
	for i := lIndex; i <= rIndex; i++ {
		if ar.Ar[i] <= ar.Ar[rIndex] {
			sPtr++
			ar.Swap(sPtr, i)
		}
		bPtr++
	}
	return sPtr
}
