package main

import (
	"fmt"
	"log"
	"sync"

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

func worker(wg *sync.WaitGroup, dir string, num int, res <-chan hw07_heapsort.SortTime) {

}

func accumResult(res chan<- hw07_heapsort.SortTime, done chan<- struct{}) {

}

func RunTest(dir string) error {
	var wg sync.WaitGroup
	rc := make(chan hw07_heapsort.SortTime)
	done := make(chan struct{})

	go accumResult(rc, done)
	for i := 0; i < testCount; i++ {
		go worker(&wg, dir, i, rc)
	}

	wg.Wait()
	done <- struct{}{}

	return nil
}

func main() {
	err := RunTest("sorting-tests/3.revers")
	if err != nil {
		log.Fatal(err)
	}
}
