package hw09linearsort

import "fmt"

type Buckets struct {
	buck []List
	len  int
}

func BucketSort(array []int64) []int64 {
	buckets := initBuckets(len(array))

	for _, v := range array {
		buckets.storeValue(v, calculateListIndex(v, len(array)))
	}

	sorted := make([]int64, 0)
	for i := range buckets.buck {
		sorted = append(sorted, buckets.getBucketArray(i)...)
	}
	return sorted
}

func calculateListIndex(value int64, len int) int64 {
	return value * int64(len) / (MAX_VALUE + 1)
}

func initBuckets(len int) Buckets {
	b := Buckets{}
	b.buck = make([]List, len)
	b.len = len
	for i := range b.buck {
		b.buck[i] = NewList()
	}

	return b
}

func (b *Buckets) storeValue(value int64, bucketNum int64) {
	l := b.buck[bucketNum]
	item := l.Front()

	for i := 0; i < l.Len(); i++ {
		fv, ok := item.Value.(int64)
		if !ok {
			return
		}
		if value < fv {
			b.buck[bucketNum].Insert(i, value)
			return
		}
		item = item.Next
	}
	b.buck[bucketNum].PushBack(value)
}

func (b *Buckets) getBucketArray(bucketNum int) []int64 {
	l := b.buck[bucketNum]
	array := make([]int64, l.Len())

	item := l.Front()
	for i := range array {
		v, _ := item.Value.(int64)
		array[i] = v
		item = item.Next
	}
	return array
}

func (b *Buckets) printBucket(index int) {
	l := b.buck[index]
	item := l.Front()
	for i := 0; i < l.Len(); i++ {
		fmt.Println(item.Value)
		item = item.Next
	}
}
