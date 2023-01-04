package sortutils

import (
	"context"
	"fmt"
	"log"
	"sync"
)

var (
	WORKERS_COUNT = 8
)

type SortConf struct {
	SortFuncs SortFunctions
	TestsDir  []string
	SizeCount int
}

type SortFunctions map[string]func(context.Context, chan<- SortTime, Array)

type taskConf struct {
	sortFunc  func(context.Context, chan<- SortTime, Array)
	num       int
	sortName  string
	arrayType string
	inFile    string
	outFile   string
}

type caseResult struct {
	task taskConf
	time SortTime
	err  error
}

type rows map[string][]SortTime
type resultTable map[string]rows
type testResult struct {
	rt  resultTable
	err error
}

func RunTest(conf SortConf) (resultTable, error) {
	log.Printf("RunTest\n")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup

	resCase := worker(ctx, &wg, scheduler(ctx, &wg, conf))

	res := <-storeResult(ctx, &wg, resCase, conf.SizeCount, len(conf.SortFuncs)*len(conf.TestsDir)*conf.SizeCount)
	cancel()
	wg.Wait()
	return res.rt, res.err
}

func scheduler(ctx context.Context, wg *sync.WaitGroup, conf SortConf) chan taskConf {
	cntrl := make(chan taskConf)
	wg.Add(1)

	go func() {
		defer wg.Done()
		sendTask(ctx, cntrl, conf)
	}()

	return cntrl
}

func sendTask(ctx context.Context, cntrl chan<- taskConf, conf SortConf) {
	for sName, sFunc := range conf.SortFuncs {
		for _, dir := range conf.TestsDir {
			for n := 0; n < conf.SizeCount; n++ {
				select {
				case <-ctx.Done():
					log.Printf("Send task. Context done\n")
					return
				case cntrl <- taskConf{
					sortFunc:  sFunc,
					sortName:  sName,
					arrayType: dir,
					inFile:    fmt.Sprintf(dir+"/test.%v.in", n),
					outFile:   fmt.Sprintf(dir+"/test.%v.out", n),
					num:       n,
				}:
					log.Printf("Send task. Send: Array: %v, Sort: %v, n = %v\n", dir, sName, n)
				}
			}
		}
	}
	log.Printf("Send task. All tasks send\n")
}

func worker(ctx context.Context, wg *sync.WaitGroup, in <-chan taskConf) chan caseResult {
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
				case conf := <-in:
					log.Printf("Worker. Get task: Array: %v, Sort: %v, n = %v\n", conf.arrayType, conf.sortName, conf.num)
					tc, err := ParseTestCase(conf.arrayType, conf.inFile, conf.outFile)
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
					case resCase <- runSort(ctx, conf, tc):
					}
				}
			}
		}()
	}

	return resCase
}

func runSort(ctx context.Context, conf taskConf, tc testCase) caseResult {
	log.Printf("Start sort: Array: %v, Sort: %v, n = %v\n", conf.arrayType, conf.sortName, conf.num)
	ar := Array{Ar: tc.Array}
	t := ar.SortArray(ctx, conf.sortFunc)

	if !ar.IsArraysEqual(tc.Expected) && !t.Timeout {
		log.Printf("Unsorted error: Array: %v, Sort: %v, n = %v\n", conf.arrayType, conf.sortName, conf.num)
		return caseResult{err: ErrArrayIsUnsorted}
	}

	log.Printf("Sorted: Array: %v, Sort: %v, n = %v; Result: %+v\n", conf.arrayType, conf.sortName, conf.num, t)
	return caseResult{task: conf, time: t}
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
			log.Printf("Store result. Get: Array: %v, Sort: %v, n = %v\n", rc.task.arrayType, rc.task.sortName, rc.task.num)
			if rc.err != nil {
				resTest <- testResult{err: rc.err}
				log.Printf("Store result. Send error\n")
				return
			}

			if table[rc.task.arrayType] == nil {
				table[rc.task.arrayType] = make(rows)
			}
			if table[rc.task.arrayType][rc.task.sortName] == nil {
				table[rc.task.arrayType][rc.task.sortName] = make([]SortTime, sizeCount)
			}
			table[rc.task.arrayType][rc.task.sortName][rc.task.num] = rc.time
		}
		resTest <- testResult{rt: table}
		log.Printf("Store result. Send result\n")
	}()

	return resTest
}
