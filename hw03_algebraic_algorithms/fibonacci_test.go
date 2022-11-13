package hw03algebraicalgorithms

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFib(t *testing.T) {
	tests := []struct {
		name     string
		in       int
		expected int
	}{
		{
			name:     "zero fib",
			in:       0,
			expected: 0,
		},
		{
			name:     "first fib",
			in:       1,
			expected: 1,
		},
		{
			name:     "second fib",
			in:       2,
			expected: 1,
		},
		{
			name:     "ten fib",
			in:       10,
			expected: 55,
		},
		{
			name:     "twenty two fib",
			in:       22,
			expected: 17711,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name+" Recursive", func(t *testing.T) {
			require.Equal(t, tc.expected, FibRecursive(tc.in), "Recursive")
		})
		t.Run(tc.name+" Iterative", func(t *testing.T) {
			require.Equal(t, tc.expected, FibIterative(tc.in), "Iterative")
		})
		t.Run(tc.name+" Golden ratio", func(t *testing.T) {
			require.Equal(t, tc.expected, FibGolden(tc.in), "Golden ratio")
		})
	}
}
