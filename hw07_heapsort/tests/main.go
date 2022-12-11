package main

import (
	"fmt"
	"log"

	"github.com/komarovn654/OTUS_Alg_Hw/hw07_heapsort"
)

func RunTest(dir string) error {
	tcNum, err := filesCount(dir)
	if err != nil {
		return err
	}
	for i := 0; i < tcNum/2; i++ {
		tc := testCase{}
		if err := tc.ParseTestCase(dir, fmt.Sprintf(dir+"\\test.%v.in", i), fmt.Sprintf(dir+"\\test.%v.out", i)); err != nil {
			return err
		}
		ar := hw07_heapsort.Array{Ar: tc.Array}
		ar.SelctionSort()
		ar.IsArraysEqual(tc.Expected)
	}
	return nil
}

func main() {
	err := RunTest("C:\\OTUS_Alg_Hw\\hw07_heapsort\\tests\\sorting-tests\\0.random")
	if err != nil {
		log.Fatal(err)
	}
}
