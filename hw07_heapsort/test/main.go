package main

import (
	"log"
	"os"

	"github.com/komarovn654/OTUS_Alg_Hw/hw07_heapsort"
	"github.com/komarovn654/OTUS_Alg_Hw/sortutils"
)

var (
	sortFunctions = []string{sortutils.SelectionSort, sortutils.HeapSort}
)

func main() {
	if len(os.Args) < 4 {
		log.Fatal("exppected four args: ... TODO")
	}
	sf := sortutils.SortFunc{"Selection Sort": hw07_heapsort.SelctionSort, "Heap Sort": hw07_heapsort.SelctionSort}

	sortutils.RunTest()
}
