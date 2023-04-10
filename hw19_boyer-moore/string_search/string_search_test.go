package stringsearch

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFullScan(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		haystack string
		needle   string
	}{
		{
			name:     "common",
			haystack: "ASdadfasdAF{{P<rQWEd32F}}",
			needle:   "P<rQWE",
			expected: 13,
		},
		{
			name:     "needle length greater then haystack",
			haystack: "A",
			needle:   "P<rQWE",
			expected: -1,
		},
		{
			name:     "haystack out of range",
			haystack: "aaa",
			needle:   "aaaa",
			expected: -1,
		},
		{
			name:     "same strings",
			haystack: "a",
			needle:   "a",
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, FullScan(tc.haystack, tc.needle))
		})
	}
}

func TestPrefixScan(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		haystack string
		needle   string
	}{
		{
			haystack: "ASdadfasdAF{{P<rQWEd32F}}",
			needle:   "P<rQWE",
			expected: 13,
		},
		{
			haystack: "A",
			needle:   "P<rQWE",
			expected: -1,
		},
		{
			haystack: "mississippi",
			needle:   "issip",
			expected: 4,
		},
		{
			haystack: "mississippi",
			needle:   "sipp",
			expected: 4,
		},
		{
			haystack: "aaa",
			needle:   "aaaa",
			expected: -1,
		},
		{
			haystack: "a",
			needle:   "a",
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s->%s", tc.needle, tc.haystack), func(t *testing.T) {
			require.Equal(t, tc.expected, PrefixScan(tc.haystack, tc.needle))
		})
	}
}

func TestScan(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		haystack string
		needle   string
	}{
		{
			haystack: "ASdadfasdAF{{P<rQWEd32F}}",
			needle:   "P<rQWE",
			expected: 13,
		},
		{
			haystack: "A",
			needle:   "P<rQWE",
			expected: -1,
		},
		{
			haystack: "mississippi",
			needle:   "issip",
			expected: 4,
		},
		{
			haystack: "mississippi",
			needle:   "sipp",
			expected: 4,
		},
		{
			haystack: "aaa",
			needle:   "aaaa",
			expected: -1,
		},
		{
			haystack: "a",
			needle:   "a",
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s->%s", tc.needle, tc.haystack), func(t *testing.T) {
			require.Equal(t, tc.expected, Scan(tc.haystack, tc.needle))
		})
	}
}
