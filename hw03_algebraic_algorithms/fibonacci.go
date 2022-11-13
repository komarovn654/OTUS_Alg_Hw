package hw03algebraicalgorithms

import (
	"math"
)

// Calculate fibonacci num recursive. O(2^N).
func FibRecursive(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return FibRecursive(n-1) + FibRecursive(n-2)
}

// Calculate fibonacci num iterative. O(N).
func FibIterative(n int) int {
	if n == 0 {
		return 0
	}
	prevprev := 0
	prev := 1
	f := 1
	for i := 1; i < n; i++ {
		f = prevprev + prev
		prevprev = prev
		prev = f
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

	return int(math.Floor(f))
}

// Calculate fibonacci num with matrix. O(LogN)
func FibMatrix(n int) int {
	m := Matrix{
		1, 1,
		1, 0,
	}
	res := m.Pwr(n - 1)
	return res.x11
}
