package hw03alg

// Iterative exponentiation algorithm. O(N).
func PwrIterative(v float64, pwr int) float64 {
	res := 1.0
	for i := 0; i < pwr; i++ {
		res = res * v
	}
	return res
}

// Multiplication of squares algorithm. O(N/2+LogN) = O(N).
func PwrSqrMultiply(v float64, pwr int) float64 {
	if pwr == 0 {
		return 1
	}
	res := 1.0
	for ; isDegreeTwo(pwr) == false; pwr -= 1 {
		res = res * v
	}
	vSqr := v * v
	for ; pwr > 0; pwr -= 2 {
		res = res * vSqr
	}
	return v
}

func isDegreeTwo(v int) bool {
	if v == 0 || v&(v-1) != 0 {
		return false
	}
	return true
}

// Binary degree decomposition algorithm. O(2LogN) = O(LogN).
func PwrBinary(v float64, pwr int) float64 {
	res := 1.0
	d := 1.0
	for ; pwr >= 1; pwr /= 2 {
		v = d * v
		d = v
		if pwr%2 == 1 {
			res = res * d
		}
	}
	return res
}
