package hw06simplesorts

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		array  []Item
		keypos []struct {
			key Item
			pos int
		}
	}{
		{array: []Item{6, 9, 40, 55, 74, 76, 78, 81, 86, 98},
			keypos: []struct {
				key Item
				pos int
			}{
				{key: 5, pos: 0}, {key: 7, pos: 1}, {key: 10, pos: 2}, {key: 45, pos: 3}, {key: 65, pos: 4},
				{key: 75, pos: 5}, {key: 77, pos: 6}, {key: 80, pos: 7}, {key: 85, pos: 8}, {key: 90, pos: 9},
				{key: 100, pos: 10},
			},
		},
		{array: []Item{6, 8, 8, 8, 74, 76, 78, 78, 86, 98},
			keypos: []struct {
				key Item
				pos int
			}{
				{key: 8, pos: 4}, {key: 78, pos: 8},
			},
		},
	}

	for _, tc := range tests {
		for _, kp := range tc.keypos {
			t.Run(strconv.Itoa(kp.pos), func(t *testing.T) {
				require.Equal(t, kp.pos, binarySearch(&tc.array, kp.key, 0, 9))
			})
		}
	}
}

func TestInsertionSort(t *testing.T) {
	t.Run("random array", func(t *testing.T) {
		ar := RandArray(100)
		done := InsertionSort(&ar)
		<-done
		require.True(t, IsSorted(&ar))
	})
	t.Run("random digits", func(t *testing.T) {
		ar := RandDigits(100)
		done := InsertionSort(&ar)
		<-done
		require.True(t, IsSorted(&ar))
	})
	t.Run("sorted array", func(t *testing.T) {
		ar := SortedArray(100)
		done := InsertionSort(&ar)
		<-done
		require.True(t, IsSorted(&ar))
	})
	t.Run("reverse array", func(t *testing.T) {
		ar := ReversArray(100)
		done := InsertionSort(&ar)
		<-done
		require.True(t, IsSorted(&ar))
	})
}

func TestInsertionSortShift(t *testing.T) {
	t.Run("random array", func(t *testing.T) {
		ar := RandArray(100)
		done := InsertionSortShift(&ar)
		<-done
		require.True(t, IsSorted(&ar))
	})
	t.Run("random digits", func(t *testing.T) {
		ar := RandDigits(100)
		done := InsertionSortShift(&ar)
		<-done
		require.True(t, IsSorted(&ar))
	})
	t.Run("sorted array", func(t *testing.T) {
		ar := SortedArray(100)
		done := InsertionSortShift(&ar)
		<-done
		require.True(t, IsSorted(&ar))
	})
	t.Run("reverse array", func(t *testing.T) {
		ar := ReversArray(100)
		done := InsertionSortShift(&ar)
		<-done
		require.True(t, IsSorted(&ar))
	})
}

func TestInsertionSortBinarySearch(t *testing.T) {
	t.Run("random array", func(t *testing.T) {
		ar := RandArray(100)
		done := InsertionSortBinarySearch(&ar)
		<-done
		require.True(t, IsSorted(&ar))
	})
	t.Run("random digits", func(t *testing.T) {
		ar := RandDigits(100)
		done := InsertionSortBinarySearch(&ar)
		<-done
		require.True(t, IsSorted(&ar))
	})
	t.Run("sorted array", func(t *testing.T) {
		ar := SortedArray(100)
		done := InsertionSortBinarySearch(&ar)
		<-done
		require.True(t, IsSorted(&ar))
	})
	t.Run("reverse array", func(t *testing.T) {
		ar := ReversArray(100)
		done := InsertionSortBinarySearch(&ar)
		<-done
		require.True(t, IsSorted(&ar))
	})
}
