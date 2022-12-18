package sortutils

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
)

type sortConf struct {
	arrayType  string
	sortMethod string
	n          int
}

type caseResult struct {
	sc   sortConf
	time SortTime
	err  error
}

type testResult struct {
	rt  resultTable
	err error
}

type rows map[string][]SortTime
type resultTable map[string]rows

func RunTest(sortTypes string, testsDir string, sizeCount int) (resultTable, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup

	caseCh := make(chan caseResult)
	testRes := scheduler(ctx, sizeCount, caseCh)

	for _, sort := range strings.Split(sortTypes, ",") {
		for _, dir := range strings.Split(testsDir, ",") {
			for n := 0; n < sizeCount; n++ {
				wg.Add(1)
				go worker(ctx, &wg, caseCh, sortConf{sortMethod: sort, arrayType: dir, n: n})
			}
		}
	}

	go func() {
		wg.Wait()
		cancel()
	}()

	select {
	case <-ctx.Done():
		res := <-testRes
		return res.rt, res.err
	case res := <-testRes:
		return nil, res.err
	}
}

func scheduler(ctx context.Context, sizeCount int, caseResCh <-chan caseResult) chan testResult {
	testCh := make(chan testResult)
	testRes := make(resultTable)

	go storeResult(ctx, sizeCount, testRes, caseResCh, testCh)

	return testCh
}

func storeResult(ctx context.Context, sizeCount int, res resultTable, caseCh <-chan caseResult, testCh chan<- testResult) {
	for {
		select {
		case <-ctx.Done():
			testCh <- testResult{rt: res}
			return
		case r := <-caseCh:
			if r.err != nil {
				testCh <- testResult{err: r.err}
				return
			}
			if res[r.sc.arrayType] == nil {
				res[r.sc.arrayType] = make(rows)
			}
			if res[r.sc.arrayType][r.sc.sortMethod] == nil {
				res[r.sc.arrayType][r.sc.sortMethod] = make([]SortTime, sizeCount)
			}
			res[r.sc.arrayType][r.sc.sortMethod][r.sc.n] = r.time
		}
	}
}

func worker(ctx context.Context, wg *sync.WaitGroup, resChan chan<- caseResult, sc sortConf) {
	defer wg.Done()
	resChan <- runSort(ctx, sc)
}

func runSort(ctx context.Context, sc sortConf) caseResult {
	tc := testCase{}
	if err := tc.ParseTestCase(sc.arrayType, fmt.Sprintf(sc.arrayType+"/test.%v.in", sc.n),
		fmt.Sprintf(sc.arrayType+"/test.%v.out", sc.n)); err != nil {
		return caseResult{err: err}
	}

	ar := Array{Ar: tc.Array}
	log.Printf("Start sort: %+v\n", sc)
	t, err := ar.SortArray(ctx, sc.sortMethod)
	if err != nil {
		return caseResult{err: err}
	}

	if !ar.IsArraysEqual(tc.Expected) && !t.Timeout {
		log.Printf("Unsorted error: %+v\n", sc)
		return caseResult{err: ErrArrayIsUnsorted}
	}

	log.Printf("Sorted: %+v; Result: %+v\n", sc, t)
	return caseResult{sc: sc, time: t}
}
