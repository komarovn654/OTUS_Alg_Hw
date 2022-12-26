package hw08quicksort

import (
	"fmt"
	"testing"

	"github.com/komarovn654/OTUS_Alg_Hw/sortutils"
)

func TestMerge(t *testing.T) {
	s := Sort{}
	s.Array.Ar = []sortutils.Item{3, 5, 7, 9, 2, 4, 6, 8}

	s.merge(0, 7, 3)
	fmt.Println(s.Array.Ar)
}

func TestMergeSort(t *testing.T) {
	s := Sort{}
	s.Array.InitRandArray(50)
	fmt.Println(s.Array)
	time := s.MergeSort()
	fmt.Println(<-time)
	fmt.Println(s.Array)
}
