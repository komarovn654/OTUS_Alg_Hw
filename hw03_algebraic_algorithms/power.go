package hw03alg

// Iterative exponentiation algorithm. O(N).
func PwrIterative(base float64, degree int64) float32 {
	res := 1.0
	for i := int64(0); i < degree; i++ {
		res = res * base
	}

	return float32(res)
}

// Multiplication of squares algorithm. O(N/2+LogN) = O(N).
func PwrSqrMultiply(base float64, degree int64) float32 {
	if degree == 0 {
		return 1
	}
	res := 1.0
	for ; !isDegreeTwo(degree); degree -= 1 {
		res = res * base
	}
	vSqr := base * base
	for ; degree > 0; degree -= 2 {
		res = res * vSqr
	}
	return float32(res)
}

func isDegreeTwo(v int64) bool {
	if v == 0 || v&(v-1) != 0 {
		return false
	}
	return true
}

// Binary degree decomposition algorithm. O(2LogN) = O(LogN).
func PwrBinary(base float64, degree int64) float32 {
	res := 1.0
	d := 1.0
	for ; degree >= 1; degree /= 2 {
		base = d * base
		d = base
		if degree%2 == 1 {
			res = res * d
		}
	}
	return float32(res)
}
