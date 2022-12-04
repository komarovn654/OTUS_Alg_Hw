package hw06simplesorts

import "time"

type Item int64

type sortedArray struct {
	array []Item
	time  sortTime
}

type sortTime struct {
	time    time.Duration
	timeout bool
}

func swap(a Item, b Item) (newA Item, newB Item) {
	return b, a
}

func isSliceContain(sl []int, item int) bool {
	for _, s := range sl {
		if s == item {
			return true
		}
	}
	return false
}
