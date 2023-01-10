package hw09linearsort

func RadixSort(array []int64, max int) []int64 {

	return nil
}

func radixSort(array []int64, digits int, max int) []int64 {
	if digits == 0 {
		return array
	}

	digit := make([]int64, len(array))
	for i, v := range array {
		digit[i] = v % 10
	}

	digit = CountingSort(digit, 9)

	return nil
}
