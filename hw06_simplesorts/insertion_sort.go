package hw06simplesorts

import (
<<<<<<< HEAD
	"context"
=======
>>>>>>> master
	"time"

	"github.com/komarovn654/OTUS_Alg_Hw/sortutils"
)

func InsertionSort(ctx context.Context, sTime chan<- sortutils.SortTime, array sortutils.Array) {
	start := time.Now()
	for i := 1; i < len(array.Ar); i++ {
		for j := i; j > 0 && (array.Ar)[j] < (array.Ar)[j-1]; j-- {
			array.Swap(j, j-1)

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

func InsertionSortShift(ctx context.Context, sTime chan<- sortutils.SortTime, array sortutils.Array) {
	start := time.Now()
	for i := 1; i < len(array.Ar); i++ {
		j := i
		cache := array.Ar[j]
		for ; j > 0 && cache < array.Ar[j-1]; j-- {
			array.Ar[j] = array.Ar[j-1]
		}
		array.Ar[j] = cache

		select {
		case <-ctx.Done():
			sTime <- sortutils.SortTime{Timeout: true}
			return
		default:
		}
	}
	sTime <- sortutils.SortTime{Time: time.Since(start)}
}

func InsertionSortBinarySearch(ctx context.Context, sTime chan<- sortutils.SortTime, array sortutils.Array) {
	start := time.Now()
	for i := 1; i < len(array.Ar); i++ {
		j := i
		cache := array.Ar[j]
		pos := binarySearch(array, cache, 0, j)
		for ; j > pos && cache < array.Ar[j-1]; j-- {
			array.Ar[j] = array.Ar[j-1]
		}
		array.Ar[j] = cache

		select {
		case <-ctx.Done():
			sTime <- sortutils.SortTime{Timeout: true}
			return
		default:
		}
	}
	sTime <- sortutils.SortTime{Time: time.Since(start)}
}

func binarySearch(a sortutils.Array, key sortutils.Item, min int, max int) int {
	if max <= min {
		if a.Ar[min] > key {
			return min
		}
		return min + 1
	}

	mid := (max + min) / 2
	if key < a.Ar[mid] {
		return binarySearch(a, key, min, mid-1)

	}
	return binarySearch(a, key, mid+1, max)
}
