package main

import (
	"errors"
	"log"

	"github.com/komarovn654/OTUS_Alg_Hw/hw07_heapsort"
)

var (
	ErrArrayIsUnsorted = errors.New("array is not sorted")
)

const (
	sortTypes = hw07_heapsort.SelectionSort + "," + hw07_heapsort.HeapSort
	testsDir  = "sorting-tests/0.random,sorting-tests/1.digits,sorting-tests/2.sorted,sorting-tests/3.revers"
	sizeCount = 8
)

func main() {
	tr, err := RunTest()
	if err != nil {
		log.Fatal(err)
	}

	if err := saveResultTable("README.md", tr); err != nil {
		log.Fatal(err)
	}
}
