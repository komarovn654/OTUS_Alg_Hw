package hw03alg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBruteforcePrime(t *testing.T) {
	tests := []struct {
		name     string
		in       int64
		expected int64
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
		t.Run(tc.name+" Brutforce", func(t *testing.T) {
			require.Equal(t, tc.expected, PrimeBruteforce(int64(tc.in)))
		})
		t.Run(tc.name+" Brutforce optimized", func(t *testing.T) {
			require.Equal(t, tc.expected, PrimeBFOpt(int64(tc.in)))
		})
		t.Run(tc.name+" Erat", func(t *testing.T) {
			require.Equal(t, tc.expected, PrimeErat(int64(tc.in)))
		})
		t.Run(tc.name+" Erat mem", func(t *testing.T) {
			require.Equal(t, tc.expected, PrimeEratMemOpt(int64(tc.in)))
		})
		t.Run(tc.name+" Erat optimized", func(t *testing.T) {
			require.Equal(t, tc.expected, PrimeEratOpt(int64(tc.in)))
		})
	}
}
