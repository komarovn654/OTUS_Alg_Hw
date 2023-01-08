package externalsort

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExternalSortTFiles(t *testing.T) {
	require.NoError(t, initTest())
	defer finTest()

	t.Run("common", func(t *testing.T) {
		maxNum := 10
		count := 100
		path := "array_" + strconv.Itoa(count) + ".txt"

		file, err := InitRandFile(path, count, maxNum)
		require.NoError(t, err)
		require.NoError(t, file.ExternalSortTFiles())

		isSorted, err := file.IsSorted()
		require.NoError(t, err)
		require.True(t, isSorted)
	})
}

func TestSplitTFiles(t *testing.T) {
	require.NoError(t, initTest())
	defer finTest()

	t.Run("common", func(t *testing.T) {
		maxNum := 10
		count := 20
		path := "array_" + strconv.Itoa(count) + ".txt"

		file, err := InitRandFile(path, count, maxNum)
		require.NoError(t, err)
		require.NoError(t, file.splitTFiles(".splited/"))
	})
}

func TestMergeTFiles(t *testing.T) {
	require.NoError(t, initTest())
	defer finTest()

	t.Run("common", func(t *testing.T) {
		maxNum := 10
		count := 20
		path := "array_" + strconv.Itoa(count) + ".txt"

		file, err := InitRandFile(path, count, maxNum)
		require.NoError(t, err)
		require.NoError(t, file.splitTFiles(".splited/"))
		require.NoError(t, file.mergeTFiles(".splited/"))
	})
}
