package hw09linearsort

func RadixSort(array []uint16, max uint16) []uint16 {
	digits := countDigits(max)

	for digit := 0; digit < digits; digit++ {
		sorted := countingSortRadix(array, digit)
		copy(array, sorted)
	}

	return array
}

func radixSort(array []uint16, digit int, lastDigit uint16) []uint16 {
	if digit == int(lastDigit) {
		return array
	}

	sorted := countingSortRadix(array, digit)
	return radixSort(sorted, digit+1, lastDigit)
}

func countingSortRadix(array []uint16, digit int) []uint16 {
	digits := make([]int, len(array))
	for i, v := range array {
		digits[i] = (int(v) % intPow(10, digit+1)) / intPow(10, digit)
	}

	indexes := make([]int, 10)
	for _, v := range digits {
		indexes[v]++
	}

	for i := 1; i < len(indexes); i++ {
		indexes[i] += indexes[i-1]
	}

	sorted := make([]uint16, len(array))
	for i := len(digits) - 1; i >= 0; i-- {
		indexes[digits[i]]--
		sorted[indexes[digits[i]]] = array[i]
	}

	return sorted
}

func countDigits(max uint16) int {
	cnt := 0
	for ; max > 0; max /= 10 {
		cnt++
	}

	return cnt
}

func intPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}
