package main

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	hw07 "github.com/komarovn654/OTUS_Alg_Hw/hw07_heapsort"
)

type sortResult struct {
	sc   sortConf
	time hw07.SortTime
	err  error
}

type sortConf struct {
	arrayType  string
	sortMethod string
	n          int
	time       hw07.SortTime
	err        error
}

type Rows map[string][]hw07.SortTime
type resultTable map[string]Rows

func RunTest() (resultTable, error) {
	var wg sync.WaitGroup
	res := make(chan sortResult)
	done := make(chan resultTable)

	go storeResult(res, done)
	for _, dir := range strings.Split(testsDir, ",") {
		for i := 0; i < sizeCount; i++ {
			wg.Add(1)
			go worker(&wg, dir, i, res, "Selection Sort")
		}
	}

	wg.Wait()
	done <- nil
	return <-done, nil
}

func storeResult(sr <-chan sortResult, done chan resultTable) {
	rt := make(resultTable)
	for {
		select {
		case <-done:
			done <- rt
			return
		case r := <-sr:
			if rt[r.sc.arrayType] == nil {
				rt[r.sc.arrayType] = make(Rows)
			}
			if rt[r.sc.arrayType][r.sc.sortMethod] == nil {
				rt[r.sc.arrayType][r.sc.sortMethod] = make([]hw07.SortTime, sizeCount)
			}
			rt[r.sc.arrayType][r.sc.sortMethod][r.sc.n] = r.time
		}
	}
}

func worker(ctx context.Context, wg *sync.WaitGroup, resChan chan<- sortResult, sc sortConf) {
	defer wg.Done()
	t := runSort(ctx, sc)
	resChan <- t
}

func runSort(ctx context.Context, sc sortConf) sortResult {
	tc := testCase{}
	if err := tc.ParseTestCase(sc.arrayType, fmt.Sprintf(sc.arrayType+"/test.%v.in", sc.n),
		fmt.Sprintf(sc.arrayType+"/test.%v.out", sc.n)); err != nil {
		return sortResult{err: err}
	}

	ctx, cancel := context.WithTimeout(ctx, 120*time.Second)
	defer cancel()
	ar := hw07.Array{Ar: tc.Array}
	t, err := ar.SortArray(ctx, sc.sortMethod)
	if err != nil {
		return sortResult{err: err}
	}

	if !ar.IsArraysEqual(tc.Expected) && !t.Timeout {
		return sortResult{err: ErrArrayIsUnsorted}
	}
	return sortResult{sc: sc, time: t}
}
