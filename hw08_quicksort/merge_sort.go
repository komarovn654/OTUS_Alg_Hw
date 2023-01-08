package hw08quicksort

import (
	"context"
	"time"

	"github.com/komarovn654/OTUS_Alg_Hw/sortutils"
)

func MergeSort(ctx context.Context, sTime chan<- sortutils.SortTime, ar sortutils.Array) {
	start := time.Now()

	mergeSort(ctx, ar, 0, len(ar.Ar)-1)

	select {
	case <-ctx.Done():
		sTime <- sortutils.SortTime{Timeout: true}
		return
	default:
	}
	sTime <- sortutils.SortTime{Time: time.Since(start)}
}

func mergeSort(ctx context.Context, ar sortutils.Array, lIndex int, rIndex int) {
	select {
	case <-ctx.Done():
		return
	default:
	}

	if lIndex >= rIndex {
		return
	}

	mIndex := (lIndex + rIndex) / 2
	mergeSort(ctx, ar, lIndex, mIndex)
	mergeSort(ctx, ar, mIndex+1, rIndex)
	merge(ar, lIndex, rIndex, mIndex)
}

func merge(ar sortutils.Array, lIndex int, rIndex int, mIndex int) {
	tmp := make([]sortutils.Item, rIndex-lIndex+1)
	m := 0
	a := lIndex
	b := mIndex + 1

	for a <= mIndex && b <= rIndex {
		if ar.Ar[a] < ar.Ar[b] {
			tmp[m] = ar.Ar[a]
			a++
		} else {
			tmp[m] = ar.Ar[b]
			b++
		}
		m++
	}

	for b <= rIndex {
		tmp[m] = ar.Ar[b]
		m++
		b++
	}
	for a <= mIndex {
		tmp[m] = ar.Ar[a]
		m++
		a++
	}

	for i := 0; i < len(tmp); i++ {
		ar.Ar[lIndex+i] = tmp[i]
	}
}
