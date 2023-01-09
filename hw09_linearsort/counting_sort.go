package hw09linearsort

func CountingSort(array []int64) []int64 {
	indexes := make([]int, MAX_VALUE)
	for _, v := range array {
		indexes[v]++
	}

	for i := 1; i < len(indexes); i++ {
		indexes[i] += indexes[i-1]
	}

	for i := len(indexes) - 1; i > 0; i-- {
		indexes[i] = indexes[i-1]
	}

	sorted := make([]int64, len(array))
	for _, v := range array {
		sorted[indexes[v]] = v
		indexes[v]++
	}

	return sorted
}
