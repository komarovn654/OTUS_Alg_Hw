package hw06simplesorts

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsSorted(t *testing.T) {
	tests := []struct {
		name string
		ar   []Item
		res  bool
	}{
		{
			name: "unsorted",
			ar:   []Item{1, 2, 3, 2, 1, 54, 62, 62, 3},
			res:  false,
		},
		{
			name: "equal",
			ar:   []Item{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			res:  true,
		},
		{
			name: "sorted",
			ar:   []Item{1, 2, 3, 4, 5, 7, 8, 9, 10, 10, 10},
			res:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.res, IsSorted(&tc.ar))
		})
	}
}

func TestArrayGen(t *testing.T) {
	fmt.Println(RandArray(20))
	fmt.Println(RandDigits(20))
	fmt.Println(SortedArray(20))
	fmt.Println(ReversArray(20))
}
