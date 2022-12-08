package main

import (
	"context"
	"fmt"
	"os"
	"time"

	hw06_sort "github.com/komarovn654/OTUS_Alg_Hw/hw06_simplesorts"
)

const (
	RandArray = iota
	RandDigits
	SortedArray
	ReversArray
)

type TestResult struct {
	SortName string
	N        int64
	SortTime hw06_sort.SortTime
	IsSorted bool
}

func getArrayTypeName(arrayType int) string {
	switch arrayType {
	case RandArray:
		return "Rand Array"
	case RandDigits:
		return "Rand Digits"
	case SortedArray:
		return "Sorted Array"
	case ReversArray:
		return "Revers Array"
	}
	return "undefine"
}

func getArray(size int64, arrayType int) *[]hw06_sort.Item {
	ar := make([]hw06_sort.Item, size)
	switch arrayType {
	case RandArray:
		ar = hw06_sort.RandArray(size, size)
	case RandDigits:
		ar = hw06_sort.RandDigits(size)
	case SortedArray:
		ar = hw06_sort.SortedArray(size)
	case ReversArray:
		ar = hw06_sort.ReversArray(size)
	}
	return &ar
}

func testShellSort(size int64, timeout time.Duration, arrayType int) TestResult {
	ar := getArray(size, arrayType)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, ar, hw06_sort.ShellSort)

	return TestResult{SortName: "Shell Sort", N: size, SortTime: result.Time, IsSorted: hw06_sort.IsSorted(ar)}
}

func testShellSortHibbard(size int64, timeout time.Duration, arrayType int) TestResult {
	ar := getArray(size, arrayType)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, ar, hw06_sort.ShellSortHibbard)

	return TestResult{SortName: "Shell Sort Hibbard", N: size, SortTime: result.Time, IsSorted: hw06_sort.IsSorted(ar)}
}

func testShellSortFrankLazarus(size int64, timeout time.Duration, arrayType int) TestResult {
	ar := getArray(size, arrayType)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, ar, hw06_sort.ShellSortFrankLazarus)

	return TestResult{SortName: "Shell Sort Frank&Lazarus", N: size, SortTime: result.Time, IsSorted: hw06_sort.IsSorted(ar)}
}

func testInsertionSort(size int64, timeout time.Duration, arrayType int) TestResult {
	ar := getArray(size, arrayType)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, ar, hw06_sort.InsertionSort)

	return TestResult{SortName: "Insertion Sort", N: size, SortTime: result.Time, IsSorted: hw06_sort.IsSorted(ar)}
}

func testInsertionSortShift(size int64, timeout time.Duration, arrayType int) TestResult {
	ar := getArray(size, arrayType)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, ar, hw06_sort.InsertionSortShift)

	return TestResult{SortName: "Insertion Sort Shift", N: size, SortTime: result.Time, IsSorted: hw06_sort.IsSorted(ar)}
}

func testInsertionSortBinarySearch(size int64, timeout time.Duration, arrayType int) TestResult {
	ar := getArray(size, arrayType)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, ar, hw06_sort.InsertionSortBinarySearch)

	return TestResult{SortName: "Insertion Binary Search", N: size, SortTime: result.Time, IsSorted: hw06_sort.IsSorted(ar)}
}

func testBubbleSort(size int64, timeout time.Duration, arrayType int) TestResult {
	ar := getArray(size, arrayType)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, ar, hw06_sort.BubbleSort)

	return TestResult{SortName: "Bubble Sort", N: size, SortTime: result.Time, IsSorted: hw06_sort.IsSorted(ar)}
}

func testBubbleSortOpt(size int64, timeout time.Duration, arrayType int) TestResult {
	ar := getArray(size, arrayType)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, ar, hw06_sort.BubbleSortOpt)

	return TestResult{SortName: "Bubble Sort optimized", N: size, SortTime: result.Time, IsSorted: hw06_sort.IsSorted(ar)}
}

func runTests(resultPath string, timeout time.Duration) error {
	f, err := os.OpenFile(resultPath, os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	for size := int64(1); size <= 10000000; size *= 10 {
		for arType := range []int{RandArray, RandDigits, SortedArray, ReversArray} {
			fmt.Println()
			fmt.Println(testBubbleSort(size, timeout, arType))
			fmt.Println(testBubbleSortOpt(size, timeout, arType))
			fmt.Println(testInsertionSort(size, timeout, arType))
			fmt.Println(testInsertionSortShift(size, timeout, arType))
			fmt.Println(testInsertionSortBinarySearch(size, timeout, arType))
			fmt.Println(testShellSort(size, timeout, arType))
			fmt.Println(testShellSortHibbard(size, timeout, arType))
			fmt.Println(testShellSortFrankLazarus(size, timeout, arType))
			fmt.Println()
		}
	}

	return nil
}
