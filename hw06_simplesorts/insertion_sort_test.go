package hw06simplesorts

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestInsertionSort(t *testing.T) {
	array := make([]Item, 20)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := range array {
		array[i] = Item(r1.Int63n(100))
	}
	fmt.Println(InsertionSort(&array))
}
