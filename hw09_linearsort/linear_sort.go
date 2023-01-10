package hw09linearsort

import (
	"math/rand"
	"time"
)

const (
	MAX_VALUE = 10
)

func GenerateArray(size int) []int64 {
	array := make([]int64, size)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := range array {
		array[i] = r1.Int63n(MAX_VALUE)
	}
	return array
}

func IsSorted(array []int64) bool {
	for i := 1; i < len(array); i++ {
		if array[i-1] > array[i] {
			return false
		}
	}
	return true
}
