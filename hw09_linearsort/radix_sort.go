package hw09linearsort

func RadixSort(array []int64, max int) []int64 {
	digitCnt := countDigits(max)
	sorted := make([]int64, len(array))

	for j := 0; j < digitCnt; j++ {
		digits := make([]int64, len(array))
		for i, v := range array {
			digits[i] = (v % intPow(10, j+1)) / intPow(10, j)
		}

		indexes := make([]int, 10)
		for _, v := range digits {
			indexes[v]++
		}

		for i := 1; i < len(indexes); i++ {
			indexes[i] += indexes[i-1]
		}

		for i := len(digits) - 1; i >= 0; i-- {
			indexes[digits[i]]--
			sorted[indexes[digits[i]]] = array[i]
		}
	}

	return sorted
}

func countingSortRadix(array []int64, digits int, max int) []int64 {
	if digits == 0 {
		return array
	}

	return nil
}

func countDigits(max int) int {
	cnt := 0
	for ; max > 0; max /= 10 {
		cnt++
	}

	return cnt
}

func intPow(n, m int) int64 {
	if m == 0 {
		return 1
	}
	result := int64(n)
	for i := 2; i <= m; i++ {
		result *= int64(n)
	}
	return result
}
