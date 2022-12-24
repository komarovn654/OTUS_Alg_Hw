package hw06simplesorts

import (
	"fmt"
	"testing"

	"github.com/komarovn654/OTUS_Alg_Hw/sortutils"
	"github.com/stretchr/testify/require"
)

func TestBubbleSort(t *testing.T) {
	tests := []int64{1, 10, 100, 1000, 10000}

	for _, tc := range tests {
		t.Run("N="+fmt.Sprint(tc), func(t *testing.T) {
			ar := sortutils.Array{}
			ar.InitRandArray(tc)

			st := BubbleSort(ar)
			time := <-st
			require.Equal(t, false, time.Timeout)
			require.True(t, ar.IsSorted())
		})
	}
}

func TestBubbleSortOpt(t *testing.T) {
	tests := []int64{1, 10, 100, 1000, 10000}

	for _, tc := range tests {
		t.Run("N="+fmt.Sprint(tc), func(t *testing.T) {
			ar := sortutils.Array{}
			ar.InitRandArray(tc)

			st := BubbleSortOpt(ar)
			time := <-st
			require.Equal(t, false, time.Timeout)
			require.True(t, ar.IsSorted())
		})
	}
}
