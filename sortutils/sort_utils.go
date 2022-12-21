package sortutils

import (
	"context"
	"errors"
	"math/rand"
	"time"
)

var (
	ErrArrayIsUnsorted = errors.New("array is not sorted")
)

const (
	SelectionSort = "Selection Sort"
	HeapSort      = "Heap Sort"
	timeout       = time.Second * 120
)

type SortFunc map[string]func(Array) <-chan SortTime

type Item int64

type Array struct {
	Ar []Item
}

type SortTime struct {
	Time    time.Duration
	Timeout bool
}

func (a *Array) SortArray(ctx context.Context, sortMethod func(Array) <-chan SortTime) SortTime {
	newCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	st := sortMethod(*a)

	select {
	case <-newCtx.Done():
		return SortTime{Timeout: true}
	case done := <-st:
		return done
	}
}

func (a *Array) IsArraysEqual(array []Item) bool {
	if len(a.Ar) != len(array) {
		return false
	}

	for i := 0; i < len(a.Ar); i++ {
		if a.Ar[i] != array[i] {
			return false
		}
	}
	return true
}

func (a *Array) InitRandArray(size int64) {
	array := make([]Item, size)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := range array {
		array[i] = Item(r1.Int63n(size))
	}
	a.Ar = array
}

func (a *Array) InitRandDigits(size int64) {

}

func (a *Array) InitSortedArray(size int64) {
	array := make([]Item, size)
	for i := int64(0); i < size-size/100; i++ {
		array[i] = Item(i)
	}
	a.Ar = array
}

func (a *Array) InitReverseArray(size int64) {

}

func (a *Array) IsSorted() bool {
	for i := 1; i < len(a.Ar); i++ {
		if a.Ar[i-1] > a.Ar[i] {
			return false
		}
	}
	return true
}

func (a *Array) FindMax(start int, end int) int {
	max := start
	for i := start + 1; i <= end; i++ {
		if a.Ar[i] > a.Ar[max] {
			max = i
		}
	}

	return max
}

func (a *Array) Swap(aIndex int, bIndex int) {
	tmp := a.Ar[aIndex]
	a.Ar[aIndex] = a.Ar[bIndex]
	a.Ar[bIndex] = tmp
}

func (a *Array) Heapify(rootIndex int, size int) {
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
	a.Heapify(p, size)
}

func (a *Array) binarySearch(key Item, min int, max int) int {
	if max <= min {
		if a.Ar[min] > key {
			return min
		}
		return min + 1
	}

	mid := (max + min) / 2
	if key < a.Ar[mid] {
		return a.binarySearch(key, min, mid-1)
	}
	return a.binarySearch(key, mid+1, max)
}
