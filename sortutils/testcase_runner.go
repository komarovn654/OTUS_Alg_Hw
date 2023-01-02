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
	log.Printf("RunTest\n")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup

	resCase := worker(ctx, &wg, scheduler(ctx, &wg, sort, testsDir, sizeCount))

	res := <-storeResult(ctx, &wg, resCase, sizeCount, len(sort)*len(testsDir)*sizeCount)
	cancel()
	wg.Wait()
	return res.rt, res.err
}

func scheduler(ctx context.Context, wg *sync.WaitGroup, sort SortFunc, testsDir []string, sizeCount int) chan sortConf {
	cntrl := make(chan sortConf)
	wg.Add(1)

	go func() {
		defer wg.Done()
		sendTask(ctx, cntrl, sort, testsDir, sizeCount)
	}()

	return cntrl
}

func sendTask(ctx context.Context, cntrl chan<- sortConf, sort SortFunc, testsDir []string, sizeCount int) {
	for sName, sFunc := range sort {
		for _, dir := range testsDir {
			for n := 0; n < sizeCount; n++ {
				select {
				case <-ctx.Done():
					log.Printf("Send task. Context done\n")
					return
				case cntrl <- sortConf{
					sortName:   sName,
					sortMethod: sFunc,
					arrayType:  dir,
					n:          n,
					inDir:      fmt.Sprintf(dir+"/test.%v.in", n),
					resDir:     fmt.Sprintf(dir+"/test.%v.out", n)}:
					log.Printf("Send task. Send: Array: %v, Sort: %v, n = %v\n", dir, sName, n)
				}
			}
		}
	}
	log.Printf("Send task. All tasks send\n")
}

func worker(ctx context.Context, wg *sync.WaitGroup, in <-chan sortConf) chan caseResult {
	resCase := make(chan caseResult)

	for i := 0; i < WORKERS_COUNT; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					log.Printf("Worker. Context done before get task")
					return
				case sc := <-in:
					log.Printf("Worker. Get task: Array: %v, Sort: %v, n = %v\n", sc.arrayType, sc.sortName, sc.n)
					tc, err := ParseTestCase(sc.arrayType, sc.inDir, sc.resDir)
					if err != nil {
						select {
						case <-ctx.Done():
							log.Printf("Worker. Context done parse error")
						case resCase <- caseResult{err: err}:
							log.Printf("Worker. Parse error")
						}
						return
					}

					select {
					case <-ctx.Done():
						log.Printf("Worker. Context done after parse")
						return
					case resCase <- runSort(ctx, sc, tc):
					}
				}
			}
		}()
	}

	return resCase
}

func runSort(ctx context.Context, sc sortConf, tc testCase) caseResult {
	log.Printf("Start sort: Array: %v, Sort: %v, n = %v\n", sc.arrayType, sc.sortName, sc.n)
	ar := Array{Ar: tc.Array}
	t := ar.SortArray(ctx, sc.sortMethod)

	if !ar.IsArraysEqual(tc.Expected) && !t.Timeout {
		log.Printf("Unsorted error: Array: %v, Sort: %v, n = %v\n", sc.arrayType, sc.sortName, sc.n)
		return caseResult{err: ErrArrayIsUnsorted}
	}

	log.Printf("Sorted: Array: %v, Sort: %v, n = %v; Result: %+v\n", sc.arrayType, sc.sortName, sc.n, t)
	return caseResult{sc: sc, time: t}
}

func storeResult(ctx context.Context, wg *sync.WaitGroup, resCase <-chan caseResult, sizeCount int, testsCount int) chan testResult {
	log.Printf("Store result. Start: tests: %v\n", testsCount)
	resTest := make(chan testResult)
	table := resultTable{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < testsCount; i++ {
			rc := <-resCase
			log.Printf("Store result. Get: Array: %v, Sort: %v, n = %v\n", rc.sc.arrayType, rc.sc.sortName, rc.sc.n)
			if rc.err != nil {
				resTest <- testResult{err: rc.err}
				log.Printf("Store result. Send error\n")
				return
			}

			if table[rc.sc.arrayType] == nil {
				table[rc.sc.arrayType] = make(rows)
			}
			if table[rc.sc.arrayType][rc.sc.sortName] == nil {
				table[rc.sc.arrayType][rc.sc.sortName] = make([]SortTime, sizeCount)
			}
			table[rc.sc.arrayType][rc.sc.sortName][rc.sc.n] = rc.time
		}
		resTest <- testResult{rt: table}
		log.Printf("Store result. Send result\n")
	}()

	return resTest
}
