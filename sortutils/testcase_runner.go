package sortutils

import (
	"context"
	"fmt"
	"log"
	"sync"
)

const (
	WORKERS_COUNT = 10
)

type sortConf struct {
	arrayType  string
	sortName   string
	sortMethod func(Array) <-chan SortTime
	n          int
	inDir      string
	resDir     string
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

func RunTest(sort SortFunc, testsDir []string, sizeCount int) (resultTable, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup

	caseCh := make(chan caseResult)
	testRes := scheduler(ctx, sizeCount, caseCh)

	worker(ctx, wg, testRes)

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

func sendTask(cntrl chan<- sortConf, sort SortFunc, testsDir []string, sizeCount int) {
	for sName, sFunc := range sort {
		for _, dir := range testsDir {
			for n := 0; n < sizeCount; n++ {
				cntrl <- sortConf{
					sortName:   sName,
					sortMethod: sFunc,
					arrayType:  dir,
					n:          n,
					inDir:      fmt.Sprintf(dir+"/test.%v.in", n),
					resDir:     fmt.Sprintf(dir+"/test.%v.out", n)}
			}
		}
	}
}

func scheduler(ctx context.Context, sort SortFunc, testsDir []string, sizeCount int) (chan testResult, chan sortConf) {
	// testCh := make(chan testResult)
	// testRes := make(resultTable)

	// go

	// return testCh
	cntrl := make(chan sortConf)
	res := make(chan caseResult)

	sendTask(cntrl, sort, testsDir, sizeCount)
	storeResult(ctx, res, sizeCount)

	return nil, cntrl
}

func worker(ctx context.Context, wg *sync.WaitGroup, resChan chan<- caseResult, in <-chan sortConf) {
	for i := 0; i < WORKERS_COUNT; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sc := <-in

			tc, err := ParseTestCase(sc.arrayType, sc.inDir, sc.resDir) //
			if err != nil {
				resChan <- caseResult{err: err}
				return
			}

			resChan <- runSort(ctx, sc, tc)
		}()
	}
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
			if res[r.sc.arrayType][r.sc.sortName] == nil {
				res[r.sc.arrayType][r.sc.sortName] = make([]SortTime, sizeCount)
			}
			res[r.sc.arrayType][r.sc.sortName][r.sc.n] = r.time
		}
	}
}

func runSort(ctx context.Context, sc sortConf, tc testCase) caseResult {
	log.Printf("Start sort: Array: %v, Sort: %v, N = %v\n", sc.arrayType, sc.sortName, tc.N)
	ar := Array{Ar: tc.Array}
	t := ar.SortArray(ctx, sc.sortMethod)

	if !ar.IsArraysEqual(tc.Expected) && !t.Timeout {
		log.Printf("Unsorted error: Array: %v, Sort: %v, N = %v\n", sc.arrayType, sc.sortName, tc.N)
		return caseResult{err: ErrArrayIsUnsorted}
	}

	log.Printf("Sorted: Array: %v, Sort: %v, N = %v; Result: %+v\n", sc.arrayType, sc.sortName, tc.N, t)
	return caseResult{sc: sc, time: t}
}
