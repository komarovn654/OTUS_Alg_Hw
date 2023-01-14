package hw09linearsort

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortFile(t *testing.T) {
	filename := "test.txt"
	GenerateFile("test.txt", 1_000_000_000, 0xFFFF)

	f, err := os.Open(filename)
	require.NoError(t, err)
	defer f.Close()

	array, err := GetArray(filename)
	require.NoError(t, err)

	sorted := RadixSort(array, 0xFFFF)
	require.True(t, IsSorted(sorted))

}
