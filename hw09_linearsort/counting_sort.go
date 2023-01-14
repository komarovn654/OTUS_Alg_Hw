package hw09linearsort

func CountingSort(array []uint16, max uint16) []uint16 {
	indexes := make([]int, max+1)
	for _, v := range array {
		indexes[v]++
	}

	for i := 1; i < len(indexes); i++ {
		indexes[i] += indexes[i-1]
	}

	sorted := make([]uint16, len(array))
	for i := len(array) - 1; i >= 0; i-- {
		indexes[array[i]]--
		sorted[indexes[array[i]]] = array[i]
	}

	return sorted
}
