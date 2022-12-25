package hw08quicksort

import (
	"fmt"
	"testing"
)

func TestPartition(t *testing.T) {
	s := Sort{}
	s.Ar.InitRandArray(20)

	fmt.Println(s.Ar.Ar)
	s.partition(0, 19)
	fmt.Println(s.Ar.Ar)
}

func TestQuickSort(t *testing.T) {
	s := Sort{}
	s.Ar.InitRandArray(10)
	fmt.Println(s.Ar)
	time := s.QuickSort()
	fmt.Println(<-time)
}
