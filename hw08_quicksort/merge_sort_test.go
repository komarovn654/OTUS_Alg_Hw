package hw08quicksort

import (
	"context"
	"testing"

	"github.com/komarovn654/OTUS_Alg_Hw/sortutils"
	"github.com/stretchr/testify/require"
)

func TestMerge(t *testing.T) {
	ar := sortutils.Array{
		Ar: []sortutils.Item{3, 5, 7, 9, 2, 4, 6, 8},
	}
	merge(ar, 0, 7, 3)

	require.True(t, ar.IsSorted())
}

func TestMergeSort(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())
	time := make(chan sortutils.SortTime)

	ar := sortutils.Array{}
	ar.InitRandArray(100)
	go MergeSort(ctx, time, ar)

	st := <-time
	require.Equal(t, false, st.Timeout)
	require.True(t, ar.IsSorted())
}
