package hw09linearsort

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortFile(t *testing.T) {
	filename := "test.txt"
	GenerateFile("test.txt", 1_000, 0xFFFF)

	f, err := os.Open(filename)
	require.NoError(t, err)

	byteArray := make([]byte, 2_000)
	intArray := make([]uint16, 1_000)
	_, err = f.Read(byteArray)
	require.NoError(t, err)

	bi := 0
	for ii := 0; ii < len(intArray); ii++ {
		intArray[ii] = uint16(byteArray[bi])
		intArray[ii] = intArray[ii]<<8 | uint16(byteArray[bi+1])
		bi += 2
	}

	sorted := BucketSort(intArray, 0xFFFF)
	require.True(t, IsSorted(sorted))

}
