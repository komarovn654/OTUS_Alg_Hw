package hw03alg

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFib(t *testing.T) {
	tests := []struct {
		name     string
		in       int
		expected *big.Int
	}{
		{
			name:     "zero fib",
			in:       0,
			expected: big.NewInt(0),
		},
		{
			name:     "first fib",
			in:       1,
			expected: big.NewInt(1),
		},
		{
			name:     "second fib",
			in:       2,
			expected: big.NewInt(1),
		},
		{
			name:     "ten fib",
			in:       10,
			expected: big.NewInt(55),
		},
		{
			name:     "twenty two fib",
			in:       22,
			expected: big.NewInt(17711),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name+" Recursive", func(t *testing.T) {
			require.Equal(t, tc.expected, FibRecursive(tc.in))
		})
		t.Run(tc.name+" Iterative", func(t *testing.T) {
			require.Equal(t, tc.expected, FibIterative(tc.in))
		})
		t.Run(tc.name+" Golden ratio", func(t *testing.T) {
			require.Equal(t, tc.expected, FibGolden(tc.in))
		})
		t.Run(tc.name+" Matrix", func(t *testing.T) {
			require.Equal(t, tc.expected, FibMatrix(tc.in))
		})
	}
}
