package hw08quicksort

import (
	"fmt"
	"testing"
)

func TestPartition(t *testing.T) {
	s := Sort{}
	s.Array.InitRandArray(20)

	fmt.Println(s.Array.Ar)
	s.partition(0, 19)
	fmt.Println(s.Array.Ar)
}

func TestQuickSort(t *testing.T) {
	s := Sort{}
	s.Array.InitRandArray(10)
	fmt.Println(s.Array)
	time := s.QuickSort()
	fmt.Println(<-time)
}
