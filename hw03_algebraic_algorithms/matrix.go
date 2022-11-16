package hw03alg

import (
	"math/big"
)

type Matrix struct {
	x11 *big.Int
	x12 *big.Int
	x21 *big.Int
	x22 *big.Int
}

// Binary degree decomposition algorithm. O(2LogN) = O(LogN).
func (m *Matrix) Pwr(pwr int) Matrix {
	res := IdentityMatrix()
	d := IdentityMatrix()

	for ; pwr >= 1; pwr /= 2 {
		m.Set(m.Multiply(d))
		d.Set(*m)
		if pwr%2 == 1 {
			res.Set(res.Multiply(d))
		}
	}

	return res
}

// Set new value.
func (m *Matrix) Set(v Matrix) Matrix {
	m.x11.Set(v.x11)
	m.x12.Set(v.x12)
	m.x21.Set(v.x21)
	m.x22.Set(v.x22)

	return *m
}

// Matrix multiplication.
func (m *Matrix) Multiply(v Matrix) Matrix {
	res := NullMatrix()
	tmp := NullMatrix()
	tmp.Set(*m)

	res.x11.Add(m.x11.Mul(m.x11, v.x11), m.x12.Mul(m.x12, v.x21))
	m.x11 = tmp.x11
	m.x12 = tmp.x12

	res.x12.Add(m.x11.Mul(m.x11, v.x12), m.x12.Mul(m.x12, v.x22))
	m.x11 = tmp.x11
	m.x12 = tmp.x12

	res.x21.Add(m.x21.Mul(m.x21, v.x11), m.x22.Mul(m.x22, v.x21))
	m.x21 = tmp.x21
	m.x22 = tmp.x22

	res.x22.Add(m.x21.Mul(m.x21, v.x12), m.x22.Mul(m.x22, v.x22))

	m.Set(res)
	return res
}

// Init identity matrix.
func IdentityMatrix() Matrix {
	return Matrix{
		x11: big.NewInt(1),
		x12: big.NewInt(0),
		x21: big.NewInt(0),
		x22: big.NewInt(1),
	}
}

// Init null matrix.
func NullMatrix() Matrix {
	return Matrix{
		x11: big.NewInt(0),
		x12: big.NewInt(0),
		x21: big.NewInt(0),
		x22: big.NewInt(0),
	}
}
