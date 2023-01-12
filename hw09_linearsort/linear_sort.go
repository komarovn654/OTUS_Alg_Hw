package hw09linearsort

import (
	"math/rand"
	"time"
)

func GenerateArray(size int, max uint16) []uint16 {
	array := make([]uint16, size)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := range array {
		array[i] = uint16(r1.Int31n(int32(max)))
	}
	return array
}

func IsSorted(array []uint16) bool {
	for i := 1; i < len(array); i++ {
		if array[i-1] > array[i] {
			return false
		}
	}
	return true
}
