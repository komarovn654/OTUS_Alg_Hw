package hw06simplesorts

import (
	"context"
	"time"
)

type Item int64

type SortedArray struct {
	Array []Item
	Time  SortTime
}

type SortTime struct {
	Time    time.Duration
	Timeout bool
}

func swap(a Item, b Item) (newA Item, newB Item) {
	return b, a
}

func SortArray(ctx context.Context, array *[]Item, sortfunc func(array *[]Item) <-chan SortTime) SortedArray {
	st := sortfunc(array)

	select {
	case <-ctx.Done():
		return SortedArray{Array: nil, Time: SortTime{Timeout: true}}
	case done := <-st:
		return SortedArray{Array: *array, Time: done}
	}
}
