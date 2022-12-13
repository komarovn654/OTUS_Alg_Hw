package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	hw07 "github.com/komarovn654/OTUS_Alg_Hw/hw07_heapsort"
)

var (
	ErrArrayIsUnsorted = errors.New("array is not sorted")
)

const (
	testsDir  = "sorting-tests/0.random,sorting-tests/1.digits,sorting-tests/2.sorted,sorting-tests/3.revers"
	sizeCount = 8
)

type caseResult struct {
	arrayType  string
	sortMethod string
	n          int
	time       hw07.SortTime
	err        error
}

type Rows map[string][]hw07.SortTime
type resultTable map[string]Rows

func runSort(dir string, num int, sm string) caseResult {
	tc := testCase{}
	if err := tc.ParseTestCase(dir, fmt.Sprintf(dir+"/test.%v.in", num), fmt.Sprintf(dir+"/test.%v.out", num)); err != nil {
		return caseResult{err: err}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	ar := hw07.Array{Ar: tc.Array}
	t, err := ar.SortArray(ctx, sm)
	if err != nil {
		return caseResult{err: err}
	}

	if !ar.IsArraysEqual(tc.Expected) && !t.Timeout {
		return caseResult{err: ErrArrayIsUnsorted}
	}
	return caseResult{arrayType: dir, time: t, n: num, sortMethod: "Selection sort"}
}

func worker(wg *sync.WaitGroup, dir string, num int, res chan<- caseResult, sm string) {
	defer wg.Done()
	t := runSort(dir, num, sm)
	res <- t
}

func storeResult(cr <-chan caseResult, done chan resultTable) {
	rt := make(resultTable)
	for {
		select {
		case <-done:
			done <- rt
			return
		case r := <-cr:
			if rt[r.arrayType] == nil {
				rt[r.arrayType] = make(Rows)
			}
			if rt[r.arrayType][r.sortMethod] == nil {
				rt[r.arrayType][r.sortMethod] = make([]hw07.SortTime, sizeCount)
			}
			rt[r.arrayType][r.sortMethod][r.n] = r.time
		}
	}
}

func RunTest() (resultTable, error) {
	var wg sync.WaitGroup
	cr := make(chan caseResult)
	done := make(chan resultTable)

	go storeResult(cr, done)
	for _, d := range strings.Split(testsDir, ",") {
		for i := 0; i < sizeCount; i++ {
			wg.Add(1)
			go worker(&wg, d, i, cr, "Selection Sort")
		}
	}

	wg.Wait()
	done <- nil
	return <-done, nil
}

func main() {

	tr, err := RunTest()
	if err != nil {
		log.Fatal(err)
	}

	// err := RunTestHeapSort(result)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println(tr)

}
