package main

import (
	"log"

	hw08quicksort "github.com/komarovn654/OTUS_Alg_Hw/hw08_quicksort"
	"github.com/komarovn654/OTUS_Alg_Hw/sortutils"
)

var (
	sortFunctions = sortutils.SortFunctions{"Merge Sort": hw08quicksort.MergeSort, "Quick Sort": hw08quicksort.QuickSort}
	testDirects   = []string{"sorting-tests/0.random", "sorting-tests/1.digits", "sorting-tests/2.sorted", "sorting-tests/3.revers"}
	testsCount    = 8
)

func main() {
	rt, err := sortutils.RunTest(sortutils.SortConf{
		SortFuncs: sortFunctions,
		TestsDir:  testDirects,
		SizeCount: testsCount,
	})
	if err != nil {
		log.Fatal(err)
	}

	if err := sortutils.SaveResultTable("README.md", rt); err != nil {
		log.Fatal(err)
	}
}
