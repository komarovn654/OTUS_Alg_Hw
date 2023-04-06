package stringsearch

func StringSearchFullScan(source string, mask string) int {
	for i := 0; i < len(source); {
		j := 0
		for ; j < len(mask); j++ {
			if source[i+j] != mask[j] {
				break
			}
		}
		if j == len(mask) {
			return i
		}

		i++
	}

	return -1
}
