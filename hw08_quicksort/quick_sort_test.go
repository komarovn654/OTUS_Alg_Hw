package hw08quicksort

import (
	"context"
	"testing"

	"github.com/komarovn654/OTUS_Alg_Hw/sortutils"
	"github.com/stretchr/testify/require"
)

func TestPartition(t *testing.T) {
	ar := sortutils.Array{
		Ar: []sortutils.Item{3, 1, 7, 9, 2, 4, 6, 8, 5},
	}

	partition(ar, 0, 8)
	require.Equal(t, []sortutils.Item{3, 1, 2, 4, 5, 9, 6, 8, 7}, ar.Ar)
}

func TestQuickSort(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())
	time := make(chan sortutils.SortTime)

	ar := sortutils.Array{}
	ar.InitRandArray(100)
	go QuickSort(ctx, time, ar)

	st := <-time
	require.Equal(t, false, st.Timeout)
	require.True(t, ar.IsSorted())
}
