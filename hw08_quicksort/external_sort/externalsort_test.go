package externalsort

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func initTest() error {
	if err := os.Mkdir(".tmp", os.ModeDir); err != nil {
		return err
	}
	if err := os.Chdir(".tmp"); err != nil {
		return err
	}

	return nil
}

func finTest() {
	os.Chdir("../")
	os.RemoveAll(".tmp")
}

func TestExternalSorts(t *testing.T) {
	defer finTest()

	tests := []struct {
		lines int
		max   int
	}{
		{lines: 100, max: 10},
		{lines: 1000, max: 10},
		{lines: 10000, max: 10},
		{lines: 100000, max: 10},
		{lines: 100, max: 100},
		{lines: 1000, max: 1000},
		{lines: 10000, max: 10000},
		{lines: 100000, max: 100000},
		{lines: 1000000, max: 1000000},
		{lines: 1000000, max: 10},
	}

	for _, tc := range tests {
		testName := fmt.Sprintf("T Files N = %v, T = %v", tc.lines, tc.max)
		t.Run(testName, func(t *testing.T) {
			initTest()
			f, err := InitRandFile("array.txt", tc.lines, tc.max)
			require.NoError(t, err)

			tm := time.Now()
			require.NoError(t, f.ExternalSortTFiles())
			ok, err := f.IsSorted()
			require.NoError(t, err)
			require.True(t, ok)
			fmt.Println("Sort time: ", time.Since(tm))

			finTest()
		})
		testName = fmt.Sprintf("2 Files N = %v, T = %v", tc.lines, tc.max)
		t.Run(testName, func(t *testing.T) {
			initTest()
			f, err := InitRandFile("array.txt", tc.lines, tc.max)
			require.NoError(t, err)

			tm := time.Now()
			require.NoError(t, f.ExternalSort2Files())
			ok, err := f.IsSorted()
			require.NoError(t, err)
			require.True(t, ok)
			fmt.Println("Sort time: ", time.Since(tm))

			finTest()
		})
		testName = fmt.Sprintf("Presort 2 Files N = %v, T = %v", tc.lines, tc.max)
		t.Run(testName, func(t *testing.T) {
			initTest()
			f, err := InitRandFile("array.txt", tc.lines, tc.max)
			require.NoError(t, err)

			tm := time.Now()
			require.NoError(t, f.ExternalSortPresort())
			ok, err := f.IsSorted()
			require.NoError(t, err)
			require.True(t, ok)
			fmt.Println("Sort time: ", time.Since(tm))

			finTest()
		})
	}
}

func TestInitRandFile(t *testing.T) {
	require.NoError(t, initTest())
	defer finTest()

	t.Run("common", func(t *testing.T) {
		maxNum := 100
		count := 20
		path := "array_" + strconv.Itoa(count) + ".txt"

		file, err := InitRandFile(path, count, maxNum)
		require.NoError(t, err)
		_, err = os.Stat(file.path)
		require.False(t, os.IsNotExist(err))
	})
}
