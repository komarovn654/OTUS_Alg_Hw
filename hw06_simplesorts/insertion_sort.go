package hw06simplesorts

func insertionSort(array *[]Item) sortedArray {
	for i := 1; i < len(*array); i++ {
		for j := i; j > 0 && (*array)[j] < (*array)[j-1]; j-- {
			(*array)[j], (*array)[j-1] = swap((*array)[j], (*array)[j-1])
		}
	}
	return sortedArray{array: *array}
}
