package hw03algebraicalgorithms

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPwr(t *testing.T) {
	tests := []struct {
		name     string
		in       float64
		pwr      int
		expected float64
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
	}

	for _, tc := range tests {
		t.Run(tc.name+" Multiply", func(t *testing.T) {
			require.Equal(t, tc.expected, PwrSqrMultiply(tc.in, tc.pwr))
		})
	}
}

func TestPwrM(t *testing.T) {
	tests := []struct {
		name     string
		in       float64
		pwr      int
		expected float64
	}{
		{
			name:     "int",
			in:       2.0,
			pwr:      9,
			expected: 1024.0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name+" Multiply", func(t *testing.T) {
			require.Equal(t, tc.expected, PwrBinary(tc.in, tc.pwr))
		})
	}
}