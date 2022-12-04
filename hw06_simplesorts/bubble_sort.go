package hw06simplesorts

import (
	"time"
)

func BubbleSort(array *[]Item) <-chan SortTime {
	sTime := make(chan SortTime)

	go func() {
		start := time.Now()
		for i := len(*array) - 1; i > 0; i-- {
			for j := 0; j < i; j++ {
				if (*array)[j+1] < (*array)[j] {
					(*array)[j], (*array)[j+1] = swap((*array)[j], (*array)[j+1])
				}
			}
		}
		sTime <- SortTime{Time: time.Since(start)}
	}()

	return sTime
}
