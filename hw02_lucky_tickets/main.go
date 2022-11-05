package main

import (
	"fmt"
)

const (
	lastDigit = 9
)

func magicFunc(refSums []int, n int) ([]int, int) {
	refSums = append(refSums, make([]int, 9)...)
	sums := make([]int, 10+(n-1)*9)
	sum := 0
	sums[0] = 1
	for i := 1; i < len(sums); i++ {
		subSliceIndex := i - 9
		if i-9 < 0 {
			subSliceIndex = 0
		}
		for _, v := range refSums[subSliceIndex : i+1] {
			sums[i] += v
		}
		sum += sums[i] * sums[i]
	}

	return sums, sum + 1
}

func main() {
	firstSlice := make([]int, 10)
	for i := 0; i < 10; i++ {
		firstSlice[i] = 1
	}

	second, ss := magicFunc(firstSlice, 2)
	fmt.Printf("n = 2\nslice: %v\nsum: %v\n", second, ss)

	third, ts := magicFunc(second, 3)
	fourth, fs := magicFunc(third, 4)

	fmt.Printf("n = 3\nslice: %v\nsum: %v\n", third, ts)
	fmt.Printf("n = 4\nslice: %v\nsum: %v\n", fourth, fs)
	//fmt.Println(magicFunc(firstSlice, 5))
	/*
		secondSlice = append(secondSlice, make([]int, 9)...)
		thirdSlice := make([]int, 28)
		thirdSlice[0] = 1
		for i := 1; i < len(thirdSlice); i++ {
			subSliceIndex := i - 9
			if i-9 < 0 {
				subSliceIndex = 0
			}
			for _, v := range secondSlice[subSliceIndex : i+1] {
				thirdSlice[i] += v
			}

			fmt.Printf("subslice: %v, ssi: %v, i:%v \n", secondSlice[subSliceIndex:i+1], subSliceIndex, i+1)
			fmt.Printf("slcie: %v\n\n", thirdSlice)
		}
	*/
}
