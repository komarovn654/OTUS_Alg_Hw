package hw03alg

import (
	"fmt"
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
		fmt.Println(m.x11, m.x12, m.x21, m.x22)
		m.Set(Multiply(*m, d))

		d.Set(*m)
		if pwr%2 == 1 {
			res.Set(Multiply(res, d))
		}

	}
	return res
}

func (m *Matrix) Set(v Matrix) Matrix {
	m.x11.Set(v.x11)
	m.x12.Set(v.x12)
	m.x21.Set(v.x21)
	m.x22.Set(v.x22)

	return *m
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

// Square matrix multiplication 2x2.
func Multiply(a Matrix, b Matrix) Matrix {
	res := Matrix{
		x11: big.NewInt(0),
		x12: big.NewInt(0),
		x21: big.NewInt(0),
		x22: big.NewInt(0),
	}
	res.x11.Add(a.x11.Mul(a.x11, b.x11), a.x12.Mul(a.x12, b.x21))
	res.x12.Add(a.x11.Mul(a.x11, b.x12), a.x12.Mul(a.x12, b.x22))
	res.x21.Add(a.x21.Mul(a.x21, b.x11), a.x22.Mul(a.x22, b.x21))
	res.x22.Add(a.x21.Mul(a.x21, b.x12), a.x22.Mul(a.x22, b.x22))
	fmt.Println(res.x11, res.x12, res.x21, res.x22)
	return res
}
