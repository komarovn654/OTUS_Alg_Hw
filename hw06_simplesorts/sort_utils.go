package hw06simplesorts

import "time"

var sortTimeout = time.Second * 120

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
