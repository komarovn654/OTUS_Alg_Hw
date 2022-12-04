package hw06simplesorts

import (
	"context"
	"time"
)

func bubbleSort(array *[]Item) <-chan sortTime {
	sTime := make(chan sortTime)

	go func() {
		start := time.Now()
		for i := len(*array) - 1; i > 0; i-- {
			for j := 0; j < i; j++ {
				if (*array)[j+1] < (*array)[j] {
					(*array)[j], (*array)[j+1] = swap((*array)[j], (*array)[j+1])
				}
			}
		}
		sTime <- sortTime{time: time.Since(start)}
	}()

	return sTime
}

func bubbleSortAnimate(array *[]Item, animate chan sortAnimation) <-chan sortTime {
	sTime := make(chan sortTime)

	go func() {
		start := time.Now()
		sa := sortAnimation{stepNum: 0, cmpCount: 0, asgCount: 0}
		for i := len(*array) - 1; i > 0; i-- {
			for j := 0; j < i; j++ {
				sa.cmp = []int{j, j + 1}
				sa.transmitCondition(animate, true)
				sa.cmpCount++
				sa.stepNum++

				if (*array)[j+1] < (*array)[j] {
					(*array)[j], (*array)[j+1] = swap((*array)[j], (*array)[j+1])
					sa.swaped = true
					sa.asgCount += 3
					sa.transmitCondition(animate, true)
					continue
				}

				sa.swaped = false
				sa.transmitCondition(animate, true)
			}
		}
		sTime <- sortTime{time: time.Since(start)}
	}()

	return sTime
}

func BubbleSort(ctx context.Context, array []Item, animate bool) sortedArray {
	var st <-chan sortTime
	if animate {
		animateChan := initBubbleAnimation(&array)
		st = bubbleSortAnimate(&array, animateChan)
	} else {
		st = bubbleSort(&array)
	}

	select {
	case <-ctx.Done():
		return sortedArray{array: nil, time: sortTime{timeout: true}}
	case done := <-st:
		return sortedArray{array: array, time: done}
	}
}
