package main

import (
	hw08quicksort "github.com/komarovn654/OTUS_Alg_Hw/hw08_quicksort"
	"github.com/komarovn654/OTUS_Alg_Hw/sortutils"
)

var (
	testDirects = []string{"sorting-tests/0.random", "sorting-tests/1.digits", "sorting-tests/2.sorted", "sorting-tests/3.revers"}
	testsCount  = 8
)

func myFunc(s sortutils.Sort) {
	s.BubbleSort()
}

func main() {
	sort := hw08quicksort.Sort{}
	myFunc(&sort)
	// rt, err := sortutils.RunTest(sortFunctions, testDirects, testsCount)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err := sortutils.SaveResultTable("README.md", rt); err != nil {
	// 	log.Fatal(err)
	// }
}
