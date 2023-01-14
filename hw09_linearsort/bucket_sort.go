package hw09linearsort

type Buckets struct {
	buck []bucket
}

func BucketSort(array []uint16, max uint16) []uint16 {
	buckets := initBuckets(len(array))

	for i, v := range array {
		buckets.buck[getBuckNum(v, len(array), max)].insert(array[i])
	}

	return buckets.getAll(len(array))
}

func getBuckNum(value uint16, len int, max uint16) int {
	return int(value) * len / (int(max) + 1)
}

func initBuckets(len int) Buckets {
	b := Buckets{}
	b.buck = make([]bucket, len)
	return b
}

type bucketItem struct {
	value    uint16
	nextItem *bucketItem
}

type bucket struct {
	frontItem *bucketItem
}

func (b *bucket) insert(item uint16) {
	if b.frontItem == nil {
		b.frontItem = &bucketItem{value: item, nextItem: nil}
		return
	}

	if item <= b.frontItem.value {
		b.frontItem = &bucketItem{value: item, nextItem: b.frontItem}
		return
	}

	for current := b.frontItem; ; current = current.nextItem {
		if current.nextItem == nil || item <= current.nextItem.value {
			current.nextItem = &bucketItem{value: item, nextItem: current.nextItem}
			return
		}
	}
}

func (b *Buckets) getAll(resultLen int) []uint16 {
	array := make([]uint16, resultLen)
	ptr := 0
	for _, b := range b.buck {
		current := b.frontItem
		buckLen := 0
		for ; current != nil; buckLen++ {
			array[buckLen+ptr] = current.value
			current = current.nextItem
		}
		ptr += buckLen
	}
	return array
}
