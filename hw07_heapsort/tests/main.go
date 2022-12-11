package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/komarovn654/OTUS_Alg_Hw/hw07_heapsort"
)

const (
	ArrayTypes = "Random Array, Random Digits, Sorted Array, Reverse Array"
	testCount  = 8
)

type Rows map[string][]hw07_heapsort.SortTime
type ResultTable map[string]Rows

func idkfunc(dir string, num int) error {
	tc := testCase{}
	if err := tc.ParseTestCase(dir, fmt.Sprintf(dir+"/test.%v.in", num), fmt.Sprintf(dir+"/test.%v.out", num)); err != nil {
		return err
	}
	ar := hw07_heapsort.Array{Ar: tc.Array}
	ar.SelctionSort()
	ar.IsArraysEqual(tc.Expected)
	return nil
}

func RunTest(dir string) error {
	// rt := make(ResultTable)
	// r := make(chan Rows)

	for _, at := range strings.Split(ArrayTypes, ", ") {
		for i := 0; i < testCount; i++ {
			fmt.Println(at)
		}
	}
	return nil
}

func main() {
	err := RunTest("sorting-tests/3.revers")
	if err != nil {
		log.Fatal(err)
	}
}
