package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	hw07 "github.com/komarovn654/OTUS_Alg_Hw/hw07_heapsort"
)

var (
	ErrArrayIsUnsorted = errors.New("array is not sorted")
)

const (
	ArrayTypes = "Random Array, Random Digits, Sorted Array, Reverse Array"
	testCount  = 8
)

type Rows []hw07.SortTime

// type Rows map[string][]hw07_heapsort.SortTime
type ResultTable map[string]Rows

func runTestCase(dir string, num int) (hw07.SortTime, error) {
	tc := testCase{}
	if err := tc.ParseTestCase(dir, fmt.Sprintf(dir+"/test.%v.in", num), fmt.Sprintf(dir+"/test.%v.out", num)); err != nil {
		return hw07.SortTime{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	ar := hw07.Array{Ar: tc.Array}
	t := ar.SortArray(ctx, ar.SelctionSort)

	if !ar.IsArraysEqual(tc.Expected) && !t.Timeout {
		return t, ErrArrayIsUnsorted
	}
	return t, nil
}

func worker(wg *sync.WaitGroup, dir string, num int, res chan<- hw07.SortTime) {
	defer wg.Done()
	t, _ := runTestCase(dir, num)
	res <- t
}

func storeResult(res <-chan hw07.SortTime, done chan Rows) {
	result := make(Rows, 0)
	for {
		select {
		case <-done:
			done <- result
			return
		case r := <-res:
			result = append(result, r)
		}
	}
}

func RunTest(dir string) error {
	var wg sync.WaitGroup
	caseResult := make(chan hw07.SortTime)
	// testResult := make(chan Rows)
	done := make(chan Rows)

	go storeResult(caseResult, done)
	for i := 0; i < testCount; i++ {
		wg.Add(1)
		go worker(&wg, dir, i, caseResult)
	}

	wg.Wait()
	done <- nil
	fmt.Println(<-done)
	return nil
}

func main() {
	fmt.Println(time.Now())
	err := RunTest("sorting-tests/3.revers")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(time.Now())
}
