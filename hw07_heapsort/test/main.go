package main

import (
	"log"

	"github.com/komarovn654/OTUS_Alg_Hw/hw07_heapsort"
	"github.com/komarovn654/OTUS_Alg_Hw/sortutils"
)

var (
	sortFunctions = sortutils.SortFunctions{"Selection Sort": hw07_heapsort.SelectionSort, "Heap Sort": hw07_heapsort.HeapSort}
	testDirects   = []string{"sorting-tests/0.random", "sorting-tests/1.digits", "sorting-tests/2.sorted", "sorting-tests/3.revers"}
	testsCount    = 8
)

func main() {
	sortutils.WORKERS_COUNT = 1

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
