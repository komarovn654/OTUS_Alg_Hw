package sortutils

import (
	"context"
	"sort"
	"testing"
	"time"
)

func sortWrapper(ar Array) <-chan SortTime {
	st := make(chan SortTime)

	go sort.Slice(ar.Ar, func(i, j int) bool {
		return ar.Ar[i] > ar.Ar[j]
	})

	return st
}

func TestRunTest(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 120*time.Second)
	sc := sortConf{
		arrayType:  "Random",
		sortName:   "Default Sort",
		sortMethod: sortWrapper,
		n:          2,
	}

	runSort(ctx, sc)
}
