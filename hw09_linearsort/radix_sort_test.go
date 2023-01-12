package hw09linearsort

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRadixSort(t *testing.T) {
	max := uint16(999)
	tests := []struct {
		len int
	}{
		{len: 10},
		{len: 100},
		{len: 1000},
		{len: 10000},
		{len: 100000},
		{len: 1000000},
	}

	for _, tc := range tests {
		t.Run("Length = "+strconv.Itoa(tc.len), func(t *testing.T) {
			array := GenerateArray(tc.len, max)

			tm := time.Now()
			sorted := RadixSort(array, max)
			fmt.Println("Sorted time = ", time.Since(tm))
			require.True(t, IsSorted(sorted))
		})
	}
}

func TestDigitCount(t *testing.T) {
	tests := []struct {
		max    uint16
		expect int
	}{
		{max: 0, expect: 0},
		{max: 1, expect: 1},
		{max: 10, expect: 2},
		{max: 100, expect: 3},
		{max: 1000, expect: 4},
		{max: 12345, expect: 5},
	}

	for _, tc := range tests {
		t.Run("max = "+strconv.Itoa(int(tc.max)), func(t *testing.T) {
			require.Equal(t, tc.expect, countDigits(tc.max))
		})
	}
}
