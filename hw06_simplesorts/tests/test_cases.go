/*
* CODE GENERATED AUTOMATICALLY
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/

package main

import (
	"context"
    "log"
	"time"

	hw06_sort "github.com/komarovn654/OTUS_Alg_Hw/hw06_simplesorts"
)

type TestResult struct {
	SortName  string
	ArrayType string
	N         int64
	SortTime  hw06_sort.SortTime
	IsSorted  bool
}

func testShellSortRandArray(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.RandArray(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.ShellSort)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("ShellSort, RandArray, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"ShellSort", "RandArray", size, result.Time, sorted}
}

func testShellSortRandDigits(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.RandDigits(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.ShellSort)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("ShellSort, RandDigits, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"ShellSort", "RandDigits", size, result.Time, sorted}
}

func testShellSortSortedArray(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.SortedArray(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.ShellSort)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("ShellSort, SortedArray, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"ShellSort", "SortedArray", size, result.Time, sorted}
}

func testShellSortReversArray(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.ReversArray(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.ShellSort)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("ShellSort, ReversArray, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"ShellSort", "ReversArray", size, result.Time, sorted}
}

func testShellSortHibbardRandArray(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.RandArray(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.ShellSortHibbard)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("ShellSortHibbard, RandArray, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"ShellSortHibbard", "RandArray", size, result.Time, sorted}
}

func testShellSortHibbardRandDigits(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.RandDigits(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.ShellSortHibbard)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("ShellSortHibbard, RandDigits, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"ShellSortHibbard", "RandDigits", size, result.Time, sorted}
}

func testShellSortHibbardSortedArray(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.SortedArray(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.ShellSortHibbard)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("ShellSortHibbard, SortedArray, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"ShellSortHibbard", "SortedArray", size, result.Time, sorted}
}

func testShellSortHibbardReversArray(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.ReversArray(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.ShellSortHibbard)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("ShellSortHibbard, ReversArray, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"ShellSortHibbard", "ReversArray", size, result.Time, sorted}
}

func testShellSortFrankLazarusRandArray(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.RandArray(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.ShellSortFrankLazarus)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("ShellSortFrankLazarus, RandArray, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"ShellSortFrankLazarus", "RandArray", size, result.Time, sorted}
}

func testShellSortFrankLazarusRandDigits(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.RandDigits(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.ShellSortFrankLazarus)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("ShellSortFrankLazarus, RandDigits, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"ShellSortFrankLazarus", "RandDigits", size, result.Time, sorted}
}

func testShellSortFrankLazarusSortedArray(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.SortedArray(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.ShellSortFrankLazarus)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("ShellSortFrankLazarus, SortedArray, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"ShellSortFrankLazarus", "SortedArray", size, result.Time, sorted}
}

func testShellSortFrankLazarusReversArray(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.ReversArray(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.ShellSortFrankLazarus)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("ShellSortFrankLazarus, ReversArray, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"ShellSortFrankLazarus", "ReversArray", size, result.Time, sorted}
}

func testInsertionSortRandArray(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.RandArray(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.InsertionSort)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("InsertionSort, RandArray, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"InsertionSort", "RandArray", size, result.Time, sorted}
}

func testInsertionSortRandDigits(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.RandDigits(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.InsertionSort)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("InsertionSort, RandDigits, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"InsertionSort", "RandDigits", size, result.Time, sorted}
}

func testInsertionSortSortedArray(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.SortedArray(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.InsertionSort)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("InsertionSort, SortedArray, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"InsertionSort", "SortedArray", size, result.Time, sorted}
}

func testInsertionSortReversArray(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.ReversArray(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.InsertionSort)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("InsertionSort, ReversArray, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"InsertionSort", "ReversArray", size, result.Time, sorted}
}

func testInsertionSortShiftRandArray(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.RandArray(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.InsertionSortShift)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("InsertionSortShift, RandArray, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"InsertionSortShift", "RandArray", size, result.Time, sorted}
}

func testInsertionSortShiftRandDigits(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.RandDigits(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.InsertionSortShift)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("InsertionSortShift, RandDigits, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"InsertionSortShift", "RandDigits", size, result.Time, sorted}
}

func testInsertionSortShiftSortedArray(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.SortedArray(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.InsertionSortShift)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("InsertionSortShift, SortedArray, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"InsertionSortShift", "SortedArray", size, result.Time, sorted}
}

func testInsertionSortShiftReversArray(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.ReversArray(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.InsertionSortShift)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("InsertionSortShift, ReversArray, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"InsertionSortShift", "ReversArray", size, result.Time, sorted}
}

func testInsertionSortBinarySearchRandArray(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.RandArray(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.InsertionSortBinarySearch)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("InsertionSortBinarySearch, RandArray, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"InsertionSortBinarySearch", "RandArray", size, result.Time, sorted}
}

func testInsertionSortBinarySearchRandDigits(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.RandDigits(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.InsertionSortBinarySearch)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("InsertionSortBinarySearch, RandDigits, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"InsertionSortBinarySearch", "RandDigits", size, result.Time, sorted}
}

func testInsertionSortBinarySearchSortedArray(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.SortedArray(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.InsertionSortBinarySearch)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("InsertionSortBinarySearch, SortedArray, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"InsertionSortBinarySearch", "SortedArray", size, result.Time, sorted}
}

func testInsertionSortBinarySearchReversArray(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.ReversArray(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.InsertionSortBinarySearch)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("InsertionSortBinarySearch, ReversArray, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"InsertionSortBinarySearch", "ReversArray", size, result.Time, sorted}
}

func testBubbleSortRandArray(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.RandArray(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.BubbleSort)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("BubbleSort, RandArray, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"BubbleSort", "RandArray", size, result.Time, sorted}
}

func testBubbleSortRandDigits(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.RandDigits(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.BubbleSort)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("BubbleSort, RandDigits, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"BubbleSort", "RandDigits", size, result.Time, sorted}
}

func testBubbleSortSortedArray(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.SortedArray(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.BubbleSort)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("BubbleSort, SortedArray, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"BubbleSort", "SortedArray", size, result.Time, sorted}
}

func testBubbleSortReversArray(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.ReversArray(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.BubbleSort)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("BubbleSort, ReversArray, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"BubbleSort", "ReversArray", size, result.Time, sorted}
}

func testBubbleSortOptRandArray(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.RandArray(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.BubbleSortOpt)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("BubbleSortOpt, RandArray, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"BubbleSortOpt", "RandArray", size, result.Time, sorted}
}

func testBubbleSortOptRandDigits(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.RandDigits(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.BubbleSortOpt)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("BubbleSortOpt, RandDigits, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"BubbleSortOpt", "RandDigits", size, result.Time, sorted}
}

func testBubbleSortOptSortedArray(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.SortedArray(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.BubbleSortOpt)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("BubbleSortOpt, SortedArray, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"BubbleSortOpt", "SortedArray", size, result.Time, sorted}
}

func testBubbleSortOptReversArray(size int64, timeout time.Duration) TestResult {
	ar := hw06_sort.ReversArray(size)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result := hw06_sort.SortArray(ctx, &ar, hw06_sort.BubbleSortOpt)
	sorted := hw06_sort.IsSorted(&ar)

	log.Printf("BubbleSortOpt, ReversArray, Size: %v, Sort Time: %+v, Sorted: %v", size, result.Time, sorted)
	return TestResult{"BubbleSortOpt", "ReversArray", size, result.Time, sorted}
}

func runTestsCases(timeout time.Duration) []TestResult {
	res := make([]TestResult, 0)
    for size := int64(1); size <= 10000000; size *= 10 {
		res = append(res, testShellSortRandArray(size, timeout))
		res = append(res, testShellSortRandDigits(size, timeout))
		res = append(res, testShellSortSortedArray(size, timeout))
		res = append(res, testShellSortReversArray(size, timeout))
		res = append(res, testShellSortHibbardRandArray(size, timeout))
		res = append(res, testShellSortHibbardRandDigits(size, timeout))
		res = append(res, testShellSortHibbardSortedArray(size, timeout))
		res = append(res, testShellSortHibbardReversArray(size, timeout))
		res = append(res, testShellSortFrankLazarusRandArray(size, timeout))
		res = append(res, testShellSortFrankLazarusRandDigits(size, timeout))
		res = append(res, testShellSortFrankLazarusSortedArray(size, timeout))
		res = append(res, testShellSortFrankLazarusReversArray(size, timeout))
		res = append(res, testInsertionSortRandArray(size, timeout))
		res = append(res, testInsertionSortRandDigits(size, timeout))
		res = append(res, testInsertionSortSortedArray(size, timeout))
		res = append(res, testInsertionSortReversArray(size, timeout))
		res = append(res, testInsertionSortShiftRandArray(size, timeout))
		res = append(res, testInsertionSortShiftRandDigits(size, timeout))
		res = append(res, testInsertionSortShiftSortedArray(size, timeout))
		res = append(res, testInsertionSortShiftReversArray(size, timeout))
		res = append(res, testInsertionSortBinarySearchRandArray(size, timeout))
		res = append(res, testInsertionSortBinarySearchRandDigits(size, timeout))
		res = append(res, testInsertionSortBinarySearchSortedArray(size, timeout))
		res = append(res, testInsertionSortBinarySearchReversArray(size, timeout))
		res = append(res, testBubbleSortRandArray(size, timeout))
		res = append(res, testBubbleSortRandDigits(size, timeout))
		res = append(res, testBubbleSortSortedArray(size, timeout))
		res = append(res, testBubbleSortReversArray(size, timeout))
		res = append(res, testBubbleSortOptRandArray(size, timeout))
		res = append(res, testBubbleSortOptRandDigits(size, timeout))
		res = append(res, testBubbleSortOptSortedArray(size, timeout))
		res = append(res, testBubbleSortOptReversArray(size, timeout))
	}
	return res
}
