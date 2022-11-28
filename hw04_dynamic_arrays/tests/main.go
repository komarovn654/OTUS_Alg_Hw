package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/komarovn654/OTUS_Alg_Hw/hw04arrays"
)

var (
	testcases = [5]int{100, 10_000, 1_000_000, 10_000_000, 100_000_000}
	timeout   = time.Second * 30
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

func runAdd(ar arrayCommon, len int, arrType int) (tr testResult) {
	timer := time.Now()
	for i := 0; i < len; i++ {
		switch arrType {
		case increase:
			ar.Add(i)
		case decrease:
			ar.Add(len - i)
		case random:
			ar.Add(rand.Intn(len))
		}

		if time.Since(timer) >= timeout {
			return testResult{dur: time.Second * 125, res: false}
		}
	}
	tr.dur = time.Since(timer)

	tr.res = true
	for i := 0; i < len; i++ {
		switch arrType {
		case increase:
			if ar.Get(i) != i {
				tr.res = false
			}
		case decrease:
			if ar.Get(i) != (len - i) {
				tr.res = false
			}
		case random:
			return tr
		}
	}

	return tr
}

func runInsert(ar arrayCommon, len int) (tr testResult) {
	timer := time.Now()
	for i := 0; i < len; i++ {
		ar.Insert(i, i)
		if time.Since(timer) >= timeout {
			return testResult{dur: time.Second * 125, res: false}
		}
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
		ar.Remove(0)
		if time.Since(timer) >= timeout {
			return testResult{dur: time.Second * 125, res: false}
		}
	}
	tr.dur = time.Since(timer)

	if ar.Size() != 0 {
		tr.res = false
	}
	return
}

func runTestInc(ar arrayCommon, len int, arrType int, arrName string) {
	res := runAdd(ar, len, arrType)
	fmt.Printf("Add %v Array: %v %v\n", arrName, res.dur, res.res)
	res = runInsert(ar, len)
	fmt.Printf("Insert %v Array: %v %v\n", arrName, res.dur, res.res)
	res = runRemove(ar, 2*len)
	fmt.Printf("Remove %v Array: %v %v\n\n", arrName, res.dur, res.res)
}

func main() {
	for _, tc := range testcases {
		fmt.Printf("Test case: %v\nI", tc)
		fmt.Println("Increasing array")
		sw := hw04arrays.InitSliceWrap()
		runTestInc(&sw, tc, increase, "Slice")
		sa := hw04arrays.InitSingleArray(0)
		runTestInc(&sa, tc, increase, "Single")
		da := hw04arrays.InitDynamicArray(0)
		runTestInc(&da, tc, increase, "Dynamic")
		fa := hw04arrays.InitFactorArray(0)
		runTestInc(&fa, tc, increase, "Factor")
		ma := hw04arrays.InitMatrixArray(0)
		runTestInc(&ma, tc, increase, "Matrix")

		fmt.Println("Decreasing array")
		sw = hw04arrays.InitSliceWrap()
		runTestInc(&sw, tc, decrease, "Slice")
		sa = hw04arrays.InitSingleArray(0)
		runTestInc(&sa, tc, decrease, "Single")
		da = hw04arrays.InitDynamicArray(0)
		runTestInc(&da, tc, decrease, "Dynamic")
		fa = hw04arrays.InitFactorArray(0)
		runTestInc(&fa, tc, decrease, "Factor")
		ma = hw04arrays.InitMatrixArray(0)
		runTestInc(&ma, tc, decrease, "Matrix")

		fmt.Println("Random array")
		sw = hw04arrays.InitSliceWrap()
		runTestInc(&sw, tc, random, "Slice")
		sa = hw04arrays.InitSingleArray(0)
		runTestInc(&sa, tc, random, "Single")
		da = hw04arrays.InitDynamicArray(0)
		runTestInc(&da, tc, random, "Dynamic")
		fa = hw04arrays.InitFactorArray(0)
		runTestInc(&fa, tc, random, "Factor")
		ma = hw04arrays.InitMatrixArray(0)
		runTestInc(&ma, tc, random, "Matrix")
		fmt.Printf("------------------------------\n")
	}
}
