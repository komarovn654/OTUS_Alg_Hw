package stringsearch

import "fmt"

func FullScan(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		j := 0
		for ; j < len(needle) && (j+i) <= len(haystack)-1; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == len(needle) {
			return i
		}
	}

	return -1
}

func PrefixTable(needle string) []int {
	table := make([]int, 128)

	for i := 1; i < len(needle)-1; i++ {
		fmt.Println(needle[i])
		if table[needle[i]] != 0 {
			continue
		}
		table[needle[i]] = i + 1
	}

	for i := range table {
		if table[i] != 0 {
			continue
		}
		table[i] = 1
	}

	if table[needle[len(needle)-1]] == 0 {
		table[needle[len(needle)-1]] = len(needle) - 1
	}
	table[needle[0]] = 1

	return table
}

func PrefixScan(haystack string, needle string) int {
	shift := PrefixTable(needle)

	for i := 0; i < len(haystack); {
		j := 0
		for ; j < len(needle) && (j+i) <= len(haystack)-1; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == len(needle) {
			return i
		}

		i += shift[needle[j]]
	}

	return -1
}

func Scan(haystack string, needle string) int {
	last := len(needle) - 1

	for i := 0; i < len(haystack)-last; i++ {
		j := 0
		for ; j <= last && haystack[i+j] == needle[j]; j++ {
		}
		if j == len(needle) {
			return i
		}
	}

	return -1
}
