package hw03alg

import (
	"math"
	"math/big"
)

// Calculate fibonacci num recursive. O(2^N).
func FibRecursive(n int) *big.Int {
	if n == 0 {
		return big.NewInt(0)
	}
	if n == 1 || n == 2 {
		return big.NewInt(1)
	}

	result := big.NewInt(0)
	return result.Add(FibRecursive(n-1), FibRecursive(n-2))
}

// Calculate fibonacci num iterative. O(N).
func FibIterative(n int) *big.Int {
	if n == 0 {
		return big.NewInt(0)
	}
	prevprev := big.NewInt(0)
	prev := big.NewInt(1)
	f := big.NewInt(1)
	for i := 1; i < n; i++ {
		f.Add(prevprev, prev)
		prevprev.Set(prev)
		prev.Set(f)

	}
	return f
}

// Calculate fibonacci num with golden ratio method. O(1)
func FibGolden(n int) *big.Int {
	if n == 0 {
		return big.NewInt(0)
	}

	if n == 1 {
		return big.NewInt(1)
	}

	p := (1 + math.Sqrt(5)) / 2
	f := math.Abs(math.Pow(p, float64(n))/math.Sqrt(5) + 1.0/2.0)

	return big.NewInt(int64(math.Floor(f)))
}

// Calculate fibonacci num with matrix. O(LogN)
func FibMatrix(n int) *big.Int {
	if n == 0 {
		return big.NewInt(0)
	}

	m := Matrix{
		x11: big.NewInt(1),
		x12: big.NewInt(1),
		x21: big.NewInt(1),
		x22: big.NewInt(0),
	}

	res := m.Pwr(n - 1)
	return res.x11
}
