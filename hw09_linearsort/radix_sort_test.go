package hw09linearsort

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRadixSort(t *testing.T) {
	tests := []struct {
		len int
	}{
		{len: 10},
		// {len: 100},
		// {len: 1000},
		// {len: 10000},
		// {len: 100000},
		// {len: 1000000},
	}

	for _, tc := range tests {
		t.Run("Length = "+strconv.Itoa(tc.len), func(t *testing.T) {
			array := GenerateArray(tc.len)

			tm := time.Now()
			sorted := RadixSort(array, MAX_VALUE)
			fmt.Println("Sorted time = ", time.Since(tm))
			require.True(t, IsSorted(sorted))
		})
	}
}

func TestDigitCount(t *testing.T) {
	tests := []struct {
		max    int
		expect int
	}{
		{max: 0, expect: 0},
		{max: 1, expect: 1},
		{max: 10, expect: 2},
		{max: 100, expect: 3},
		{max: 1000, expect: 4},
		{max: 1234567890, expect: 10},
	}

	for _, tc := range tests {
		t.Run("max = "+strconv.Itoa(tc.max), func(t *testing.T) {
			require.Equal(t, tc.expect, countDigits(tc.max))
		})
	}
}
