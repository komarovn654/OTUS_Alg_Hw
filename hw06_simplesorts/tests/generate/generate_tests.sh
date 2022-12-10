#!/bin/bash

start_size=1
end_size=10000000
sort_names=(ShellSort ShellSortHibbard ShellSortFrankLazarus InsertionSort InsertionSortShift InsertionSortBinarySearch BubbleSort BubbleSortOpt)
array_types=(RandArray RandDigits SortedArray ReversArray)

touch tmp.txt
> test_cases.go

cat << EOF > test_cases.go
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

EOF

declare -a test_names
for sn in ${sort_names[*]}
do
    for at in ${array_types[*]}
    do
        > tmp.txt
        sed "s/%SortName%/$sn/g" $1 >> tmp.txt
        sed "s/%ArrayType%/$at/g" tmp.txt >> test_cases.go
        echo >> test_cases.go
        test_names+=( "test$sn$at" )
    done
done

echo "func runTestsCases(timeout time.Duration) []TestResult {
	res := make([]TestResult, 0)
    for size := int64($start_size); size <= $end_size; size *= 10 {"  >> test_cases.go
for tm in ${test_names[*]}
do
    echo -e "\t\tres = append(res, $tm(size, timeout))" >> test_cases.go
done
echo -e "\t}\n\treturn res\n}" >> test_cases.go

rm tmp.txt