package hw03alg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPwr(t *testing.T) {
	tests := []struct {
		name     string
		in       float64
		pwr      int64
		expected float32
	}{
		{
			name:     "int",
			in:       4,
			pwr:      4,
			expected: 256,
		},
		{
			name:     "int",
			in:       2,
			pwr:      17,
			expected: 131072,
		},
		{
			name:     "zero degree",
			in:       5,
			pwr:      0,
			expected: 1,
		},
		{
			name:     "zero value",
			in:       0,
			pwr:      5,
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name+" Iterative", func(t *testing.T) {
			require.Equal(t, tc.expected, PwrIterative(tc.in, tc.pwr))
		})
		t.Run(tc.name+" Multiply", func(t *testing.T) {
			require.Equal(t, tc.expected, PwrSqrMultiply(tc.in, tc.pwr))
		})
		t.Run(tc.name+" Binary", func(t *testing.T) {
			require.Equal(t, tc.expected, PwrBinary(tc.in, tc.pwr))
		})
	}
}
