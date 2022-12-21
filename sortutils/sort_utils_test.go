package sortutils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindMax(t *testing.T) {
	s := Array{}
	ar := []Item{18, 12, 3, 14, 7, 9, 18, 0, 18, 13, 4, 17, 19, 18, 4, 11, 12, 5, 10, 18}
	s.Ar = ar
	t.Run("max index", func(t *testing.T) {
		require.Equal(t, 0, s.FindMax(0, 5))
		require.Equal(t, 0, s.FindMax(0, 10))
		require.Equal(t, 12, s.FindMax(0, 19))
		require.Equal(t, 12, s.FindMax(12, 19))
	})
}

func TestIsSorted(t *testing.T) {
	tests := []struct {
		name  string
		array []Item
		res   bool
	}{
		{
			name:  "sorted",
			array: []Item{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			res:   true,
		},
		{
			name:  "equaled",
			array: []Item{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			res:   true,
		},
		{
			name:  "reversed",
			array: []Item{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			res:   false,
		},
	}

	for _, tc := range tests {
		_ = tc
		t.Run(tc.name, func(t *testing.T) {
			a := Array{}
			a.Ar = tc.array
			require.Equal(t, tc.res, a.IsSorted())
		})
	}
}
