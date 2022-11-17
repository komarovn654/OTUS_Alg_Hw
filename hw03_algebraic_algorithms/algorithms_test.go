package hw03alg

import (
	"testing"
)

func TestBruteforcePrime(t *testing.T) {
	tests := []struct {
		name     string
		in       int
		expected int
	}{
		{
			name:     "range = 0",
			in:       0,
			expected: 0,
		},
		{
			name:     "range = 1",
			in:       1,
			expected: 0,
		},
		{
			name:     "range = 2",
			in:       2,
			expected: 1,
		},
		{
			name:     "range = 10",
			in:       10,
			expected: 4,
		},
		{
			name:     "range = 1000",
			in:       1000,
			expected: 168,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name+" brutforce", func(t *testing.T) {
			// v, err := BruteforcePrime(tc.in)
			// require.Equal(t, tc.err, err)
			// require.Equal(t, tc.expected, v)
		})
		t.Run(tc.name+" with optimizations", func(t *testing.T) {
			// v, err := PrimeOpt(tc.in)
			// require.Equal(t, tc.err, err)
			// require.Equal(t, tc.expected, v)
		})
		t.Run(tc.name+" er", func(t *testing.T) {
			// v, err := PrimeEr(tc.in)
			// require.Equal(t, tc.err, err)
			// require.Equal(t, tc.expected, v)
		})
		t.Run(tc.name+" er optimizations", func(t *testing.T) {
			// v, err := PrimeErOpt(tc.in)
			// require.Equal(t, tc.err, err)
			// require.Equal(t, tc.expected, v)
		})
	}
}
