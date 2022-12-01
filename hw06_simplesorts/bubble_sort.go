package hw06simplesorts

import (
	"time"
)

func bubbleSort(array *[]Item, itr chan<- struct{}) <-chan sortTime {
	sTime := make(chan sortTime)

	go func() {
		start := time.Now()
		for i := len(*array) - 1; i > 0; i-- {
			for j := 0; j < i; j++ {
				if (*array)[j+1] < (*array)[j] {
					(*array)[j], (*array)[j+1] = swap((*array)[j], (*array)[j+1])
				}
				select {
				case <-time.After(sortTimeout):
					sTime <- sortTime{timeout: true}
					return
				default:
					itr <- struct{}{}
				}
			}
		}
		sTime <- sortTime{time: time.Since(start)}
	}()

	return sTime
}

func BubbleSort(array []Item) sortedArray {
	itr := make(chan struct{})

	sTime := bubbleSort(&array, itr)

	for {
		select {
		case st := <-sTime:
			return sortedArray{array: array, time: st}
		case <-itr:
			// case with each iteration. Do some animation
		}
	}
}
