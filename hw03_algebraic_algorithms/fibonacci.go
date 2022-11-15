package hw03algebraicalgorithms

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
func FibGolden(n int) int {
	if n == 1 {
		return 1
	}

	p := (1 + math.Sqrt(5)) / 2
	f := math.Abs(math.Pow(p, float64(n))/math.Sqrt(5) + 1/2)

	return int(math.Round(f))
}

// Calculate fibonacci num with matrix. O(LogN)
func FibMatrix(n int) int {
	if n == 0 {
		return 0
	}

	m := Matrix{
		1, 1,
		1, 0,
	}
	res := m.Pwr(n - 1)
	return res.x11
}
