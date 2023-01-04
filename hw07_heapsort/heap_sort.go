package hw07_heapsort

import (
	"context"
	"time"

	"github.com/komarovn654/OTUS_Alg_Hw/sortutils"
)

func HeapSort(ctx context.Context, sTime chan<- sortutils.SortTime, array sortutils.Array) {
	start := time.Now()

	for h := len(array.Ar) / 2; h >= 0; h-- {
		heapify(array, h, len(array.Ar))
	}
	for i := len(array.Ar) - 1; i > 0; i-- {
		array.Swap(0, i)
		heapify(array, 0, i)

		select {
		case <-ctx.Done():
			sTime <- sortutils.SortTime{Timeout: true}
			return
		default:
		}
	}

	sTime <- sortutils.SortTime{Time: time.Since(start)}
}

func heapify(a sortutils.Array, rootIndex int, size int) {
	lIndex := 2*rootIndex + 1
	rIndex := 2*rootIndex + 2
	p := rootIndex
	if lIndex < size && a.Ar[lIndex] > a.Ar[p] {
		p = lIndex
	}
	if rIndex < size && a.Ar[rIndex] > a.Ar[p] {
		p = rIndex
	}
	if p == rootIndex {
		return
	}
	a.Swap(rootIndex, p)
	heapify(a, p, size)
}
