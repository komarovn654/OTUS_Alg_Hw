package main

import (
	"log"

	hw06simplesorts "github.com/komarovn654/OTUS_Alg_Hw/hw06_simplesorts"
	"github.com/komarovn654/OTUS_Alg_Hw/sortutils"
)

var (
	sortFunctions = sortutils.SortFunc{
		"Bubble Sort": hw06simplesorts.BubbleSort, "Bubble Sort Opt": hw06simplesorts.BubbleSortOpt,
		"Insertion Sort": hw06simplesorts.InsertionSort, "Insertion Sort Shift": hw06simplesorts.InsertionSortShift,
		"Insertion Sort BinarySearch": hw06simplesorts.InsertionSortBinarySearch, "Shell Sort": hw06simplesorts.ShellSort,
		"Shell Sort Frank&Lazarus": hw06simplesorts.ShellSortFrankLazarus, "Shell Sort Hibbard": hw06simplesorts.ShellSortHibbard,
	}
	testDirects = []string{"sorting-tests/0.random", "sorting-tests/1.digits", "sorting-tests/2.sorted", "sorting-tests/3.revers"}
	testsCount  = 8
)

func main() {
	rt, err := sortutils.RunTest(sortFunctions, testDirects, testsCount)
	if err != nil {
		log.Fatal(err)
	}

	if err := sortutils.SaveResultTable("README.md", rt); err != nil {
		log.Fatal(err)
	}
}
