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

func TestBinarySearch(t *testing.T) {
	// array := RandArray(10, 100)
	// w := BubbleSort(&array)
	// <-w
	// fmt.Println(array)
	// time.Sleep(time.Second)
	array := []Item{3, 3, 5, 15, 18, 46, 50, 50, 64, 93}
	array = []Item{13, 18, 21, 40, 46, 52, 68, 75, 85, 92}
	array = []Item{27, 28, 31, 44, 48, 53, 59, 63, 87, 96}
	array = []Item{6, 9, 40, 55, 74, 74, 78, 81, 86, 98}

	fmt.Println(binarySearch(&array, 63, 0, 9))
}
