package hw09linearsort

func CountingSort(array []int64, max int) []int64 {
	indexes := make([]int, max)
	for _, v := range array {
		indexes[v]++
	}

	for i := 1; i < len(indexes); i++ {
		indexes[i] += indexes[i-1]
	}

	sorted := make([]int64, len(array))
	for i := len(array) - 1; i >= 0; i-- {
		indexes[array[i]]--
		sorted[indexes[array[i]]] = array[i]
	}

	return sorted
}
