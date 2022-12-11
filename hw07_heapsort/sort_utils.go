package hw07_heapsort

import (
	"fmt"
	"math/rand"
	"time"
)

// type Sort interface {
// 	InitRandArray(int64)
// 	InitRandDigits(int64)
// 	InitSortedArray(int64)
// 	InitReverseArray(int64)
// }

type Item int64

type Array struct {
	Ar []Item
}

type SortTime struct {
}

func (s *Array) IsArraysEqual(array []Item) bool {
	if len(s.Ar) != len(array) {
		return false
	}

	for i := 0; i < len(s.Ar); i++ {
		fmt.Printf("%v == %v\n", s.Ar[i], array[i])
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
