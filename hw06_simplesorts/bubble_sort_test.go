package hw06simplesorts

import (
	"math/rand"
	"testing"
	"time"
)

func TestBubbleSort(t *testing.T) {
	array := make([]Item, 10)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := range array {
		array[i] = Item(r1.Int63n(100))
	}
}
