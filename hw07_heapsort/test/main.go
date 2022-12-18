package main

import (
	"fmt"
	"log"

	"github.com/komarovn654/OTUS_Alg_Hw/sortutils"
)

var (
	sortFunctions = []string{sortutils.SelectionSort, sortutils.HeapSort}
	testDirects   = []string{"sorting-tests/0.random", "sorting-tests/1.digits", "sorting-tests/2.sorted", "sorting-tests/3.revers"}
	testsCount    = 8
)

func main() {
	rt, err := sortutils.RunTest(sortFunctions, testDirects, testsCount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rt)
}
