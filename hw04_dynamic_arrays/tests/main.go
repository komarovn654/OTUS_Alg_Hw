package main

import (
	"fmt"
	"time"

	"github.com/komarovn654/OTUS_Alg_Hw/hw04arrays"
)

var (
	test100       = 100
	test10000     = 10_000
	test1000000   = 1_000_000
	test100000000 = 100_000_000
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

func runAdd(ar arrayCommon, len int) (tr testResult) {
	timer := time.Now()
	for i := 0; i < len; i++ {
		ar.Add(i)
	}
	tr.dur = time.Since(timer)

	tr.res = true
	for i := 0; i < len; i++ {
		if ar.Get(i) != i {
			tr.res = false
		}
	}

	return tr
}

func runInsert(ar arrayCommon, len int) (tr testResult) {
	timer := time.Now()
	for i := 0; i < len; i++ {
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
		ar.Remove(0)
	}
	tr.dur = time.Since(timer)

	if ar.Size() != 0 {
		tr.res = false
	}
	return
}

func main() {
	// sa := hw04arrays.InitSingleArray(0)
	// res := runAdd(&sa, test100)
	// fmt.Printf("%v %v\n", res.dur, res.res)
	// res = runInsert(&sa, test100)
	// fmt.Printf("%v %v\n", res.dur, res.res)
	// res = runRemove(&sa, 2*test100)
	// fmt.Printf("%v %v\n", res.dur, res.res)

	da := hw04arrays.InitDynamicArray(0)
	res := runAdd(&da, test10000)
	fmt.Printf("%v %v\n", res.dur, res.res)

	fa := hw04arrays.InitFactorArray(0)
	res = runAdd(&fa, test100000000)
	fmt.Printf("%v %v\n", res.dur, res.res)

	ma := hw04arrays.InitMatrixArray(0)
	res = runAdd(&ma, test100000000)
	fmt.Printf("%v %v\n", res.dur, res.res)
}
