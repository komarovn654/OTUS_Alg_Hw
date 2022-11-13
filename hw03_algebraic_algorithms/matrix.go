package hw03algebraicalgorithms

type Matrix struct {
	x11 int
	x12 int
	x21 int
	x22 int
}

// Init identity matrix.
func IdentityMatrix() Matrix {
	return Matrix{1, 0, 0, 1}
}

// Square matrix multiplication 2x2.
func (m Matrix) Multiply(v Matrix) Matrix {
	res := Matrix{}
	res.x11 = m.x11*v.x11 + m.x12*v.x21
	res.x12 = m.x11*v.x12 + m.x12*v.x22
	res.x21 = m.x21*v.x11 + m.x22*v.x21
	res.x22 = m.x21*v.x12 + m.x22*v.x22
	return res
}

// Binary degree decomposition algorithm. O(2LogN) = O(LogN).
func (m Matrix) Pwr(pwr int) Matrix {
	res := IdentityMatrix()
	d := IdentityMatrix()
	for ; pwr >= 1; pwr /= 2 {
		m = m.Multiply(d)
		d = m
		if pwr%2 == 1 {
			res = res.Multiply(d)
		}
	}
	return res
}
