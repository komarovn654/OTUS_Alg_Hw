package hw06simplesorts

import (
	"context"
	"math/rand"
	"time"
)

type Item int64

type SortedArrayRes struct {
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

func SortArray(ctx context.Context, array *[]Item, sortfunc func(array *[]Item) <-chan SortTime) SortedArrayRes {
	st := sortfunc(array)

	select {
	case <-ctx.Done():
		return SortedArrayRes{Array: nil, Time: SortTime{Timeout: true}}
	case done := <-st:
		return SortedArrayRes{Array: *array, Time: done}
	}
}

func RandArray(size int64) []Item {
	array := make([]Item, size)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := range array {
		array[i] = Item(r1.Int63n(size))
	}
	return array
}

func RandDigits(size int64) []Item {
	array := make([]Item, size)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := range array {
		array[i] = Item(r1.Int63n(9))
	}
	return array
}

func SortedArray(size int64) []Item {
	ar := make([]Item, size)
	for i := int64(0); i < size-size/100; i++ {
		ar[i] = Item(i)
	}
	return ar
}

func ReversArray(size int64) []Item {
	ar := make([]Item, size)
	for i := size - 1; i > 0; i-- {
		ar[i] = Item(i)
	}
	return ar
}

func IsSorted(array *[]Item) bool {
	prev := (*array)[0]
	for i := 1; i < len(*array); i++ {
		if prev > (*array)[i] {
			return false
		}
		prev = (*array)[i]
	}
	return true
}
