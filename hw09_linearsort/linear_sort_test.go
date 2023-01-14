package hw09linearsort

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSortFile(t *testing.T) {
	filename := "test.txt"
	defer os.Remove(filename)
	GenerateFile("test.txt", 1_000_000_000, 0xFFFF)

	f, err := os.Open(filename)
	require.NoError(t, err)
	defer f.Close()

	t.Run("Bucket Sort, N = 10^9", func(t *testing.T) {
		array, err := GetArray(filename)
		require.NoError(t, err)

		tm := time.Now()
		sorted := BucketSort(array, 0xFFFF)
		fmt.Println("Sorted time = ", time.Since(tm))
		require.True(t, IsSorted(sorted))
	})
	t.Run("Counting Sort, N = 10^9", func(t *testing.T) {
		array, err := GetArray(filename)
		require.NoError(t, err)

		tm := time.Now()
		sorted := CountingSort(array, 0xFFFF)
		fmt.Println("Sorted time = ", time.Since(tm))
		require.True(t, IsSorted(sorted))
	})
	t.Run("Radix Sort, N = 10^9", func(t *testing.T) {
		array, err := GetArray(filename)
		require.NoError(t, err)

		tm := time.Now()
		sorted := RadixSort(array, 0xFFFF)
		fmt.Println("Sorted time = ", time.Since(tm))
		require.True(t, IsSorted(sorted))
	})
}
