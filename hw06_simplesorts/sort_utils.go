package hw06simplesorts

import (
	"context"
	"math/rand"
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

func RandArray(size int64, rndRange int64) []Item {
	array := make([]Item, size)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := range array {
		array[i] = Item(r1.Int63n(rndRange))
	}
	return array
}
