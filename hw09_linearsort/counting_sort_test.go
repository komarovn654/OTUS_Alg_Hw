package hw09linearsort

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// go test -run TestCountingSort -timeout 120s -v
func TestCountingSort(t *testing.T) {
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
			array := GenerateArray(tc.len)

			tm := time.Now()
			sorted := CountingSort(array)
			fmt.Println("Sorted time = ", time.Since(tm))
			require.True(t, IsSorted(sorted))
		})
	}
}
