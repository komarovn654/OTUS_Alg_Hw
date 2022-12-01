package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/komarovn654/OTUS_Alg_Hw/hw04arrays"
)

var (
	testcases  = [5]int{1_000, 10_000, 100_000, 1_000_000, 10_000_000}
	arrayOrder = [3]string{"increase", "decrease", "random"}
	timeout    = time.Second * 30
)

const (
	increase = iota
	decrease
	random
)

type arrayCommon interface {
	Get(index int) hw04arrays.Item
	Add(t hw04arrays.Item)
	Insert(t hw04arrays.Item, index int)
	Remove(index int) hw04arrays.Item
	Size() int
}

type testResult struct {
	dur time.Duration
	res bool
}

func runAdd(ar arrayCommon, len int, arrType string) (tr testResult) {
	timer := time.Now()
	for i := 0; i < len; i++ {
		if time.Since(timer) >= timeout {
			return testResult{dur: timeout, res: false}
		}
		switch arrType {
		case "increase":
			ar.Add(i)
		case "decrease":
			ar.Add(len - i)
		case "random":
			ar.Add(rand.Intn(len))
		}
	}
	tr.dur = time.Since(timer)

	tr.res = true
	for i := 0; i < len; i++ {
		switch arrType {
		case "increase":
			if ar.Get(i) != i {
				tr.res = false
			}
		case "decrease":
			if ar.Get(i) != (len - i) {
				tr.res = false
			}
		case "random":
			return tr
		}
	}

	return tr
}

func runInsert(ar arrayCommon, len int) (tr testResult) {
	timer := time.Now()
	for i := 0; i < len; i++ {
		if time.Since(timer) >= timeout {
			return testResult{dur: timeout, res: false}
		}
		ar.Insert(i, i)
	}
	tr.dur = time.Since(timer)

	tr.res = true
	for j := 0; j < 2; j++ {
		for i := 0; i < len; i++ {
			if ar.Get(i) != i {
				tr.res = false
			}
		}
	}

	return tr
}

func runRemove(ar arrayCommon, len int) (tr testResult) {
	tr.res = true
	timer := time.Now()
	for i := 0; i < len; i++ {
		if time.Since(timer) >= timeout {
			return testResult{dur: timeout, res: false}
		}
		ar.Remove(0)
	}
	tr.dur = time.Since(timer)

	if ar.Size() != 0 {
		tr.res = false
	}
	return
}

func runTest(ar arrayCommon, len int, arrType string, arrName string) {
	res := runAdd(ar, len, arrType)
	fmt.Printf("Add %v Array: %v %v\n", arrName, res.dur, res.res)
	res = runInsert(ar, len)
	fmt.Printf("Insert %v Array: %v %v\n", arrName, res.dur, res.res)
	res = runRemove(ar, 2*len)
	fmt.Printf("Remove %v Array: %v %v\n\n", arrName, res.dur, res.res)
}

func main() {
	for _, tc := range testcases {
		for _, order := range arrayOrder {
			fmt.Printf("Test case: %v\n", tc)
			fmt.Printf("%v array\n", order)
			sw := hw04arrays.InitSliceWrap()
			sa := hw04arrays.InitSingleArray(0)
			da := hw04arrays.InitDynamicArray(0)
			fa := hw04arrays.InitFactorArray(0)
			ma := hw04arrays.InitMatrixArray(0)

			runTest(&sw, tc, order, "Slice")
			runTest(&sa, tc, order, "Single")
			runTest(&da, tc, order, "Dynamic")
			runTest(&fa, tc, order, "Factor")
			runTest(&ma, tc, order, "Matrix")
			fmt.Printf("------------------------------\n")
		}
	}
}
