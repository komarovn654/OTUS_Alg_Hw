package hw07_heapsort

import (
	"context"
	"errors"
	"math/rand"
	"time"
)

var (
	ErrUnknownMethod = errors.New("unknown sort method")
)

const (
	SelectionSort = "Selection Sort"
	HeapSort      = "Heap Sort"
	timeout       = time.Second * 120
)

type Item int64

type Array struct {
	Ar   []Item
	Sort func() <-chan SortTime
}

type SortedTime struct {
	Time    time.Duration
	Timeout bool
}

type SortTime struct {
	Time    time.Duration
	Timeout bool
}

func (s *Array) chooseSortMethod(sm string) func() <-chan SortTime {
	switch sm {
	case SelectionSort:
		return s.SelctionSort
	case HeapSort:
		return s.HeapSort
	default:
		return nil
	}
}

func (s *Array) SortArray(ctx context.Context, sortMethod string) (SortTime, error) {
	newCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	sortFunc := s.chooseSortMethod(sortMethod)
	if sortFunc == nil {
		return SortTime{}, ErrUnknownMethod
	}
	st := sortFunc()

	select {
	case <-newCtx.Done():
		return SortTime{Timeout: true}, nil
	case done := <-st:
		return done, nil
	}
}

func (s *Array) IsArraysEqual(array []Item) bool {
	if len(s.Ar) != len(array) {
		return false
	}

	for i := 0; i < len(s.Ar); i++ {
		if s.Ar[i] != array[i] {
			return false
		}
	}
	return true
}

func (s *Array) InitRandArray(size int64) {
	array := make([]Item, size)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := range array {
		array[i] = Item(r1.Int63n(size))
	}
	s.Ar = array
}

func (s *Array) InitRandDigits(size int64) {

}

func (s *Array) InitSortedArray(size int64) {
	array := make([]Item, size)
	for i := int64(0); i < size-size/100; i++ {
		array[i] = Item(i)
	}
	s.Ar = array
}

func (s *Array) InitReverseArray(size int64) {

}

func (s *Array) IsSorted() bool {
	for i := 1; i < len(s.Ar); i++ {
		if s.Ar[i-1] > s.Ar[i] {
			return false
		}
	}
	return true
}

func (s *Array) findMax(start int, end int) int {
	max := start
	for i := start + 1; i <= end; i++ {
		if s.Ar[i] > s.Ar[max] {
			max = i
		}
	}

	return max
}

func (s *Array) swap(aIndex int, bIndex int) {
	tmp := s.Ar[aIndex]
	s.Ar[aIndex] = s.Ar[bIndex]
	s.Ar[bIndex] = tmp
}

func (s *Array) heapify(rootIndex int, size int) {
	lIndex := 2*rootIndex + 1
	rIndex := 2*rootIndex + 2
	p := rootIndex
	if lIndex < size && s.Ar[lIndex] > s.Ar[p] {
		p = lIndex
	}
	if rIndex < size && s.Ar[rIndex] > s.Ar[p] {
		p = rIndex
	}
	if p == rootIndex {
		return
	}
	s.swap(rootIndex, p)
	s.heapify(p, size)
}
