package hw07_heapsort

import (
	"testing"

	"github.com/komarovn654/OTUS_Alg_Hw/sortutils"
	"github.com/stretchr/testify/require"
)

func TestSelectionSort(t *testing.T) {
	tests := []int64{1, 10, 100, 1000, 10000, 100000}

	for _, tc := range tests {
		ar := sortutils.Array{}
		ar.InitRandArray(tc)

		st := SelectionSort(ar)
		time := <-st
		require.Equal(t, false, time.Timeout)
		require.True(t, ar.IsSorted())
	}
}
