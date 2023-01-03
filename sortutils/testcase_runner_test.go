package sortutils

import (
	"context"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func SortTest(ctx context.Context, st chan<- SortTime, ar Array) {
	start := time.Now()

	sort.Slice(ar.Ar, func(i, j int) bool {
		return ar.Ar[i] < ar.Ar[j]
	})

	st <- SortTime{Time: time.Since(start)}
}

func TestRunSort(t *testing.T) {
	ctx := context.Background()
	sc := taskConf{
		sortFunc:  SortTest,
		arrayType: "Random",
		sortName:  "Default Sort",
	}
	tc := testCase{
		Name:     "Random",
		Array:    []Item{5, 9, 8, 15, 10, 3, 6, 4, 18, 2, 19, 1, 17, 12, 7, 14, 13, 0, 11, 16},
		Expected: []Item{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19},
	}

	tr := runSort(ctx, sc, tc)
	require.NoError(t, tr.err)
	require.Equal(t, false, tr.time.Timeout)
}

func TestRunTest(t *testing.T) {
	t.Run("successful test", func(t *testing.T) {
		defer goleak.VerifyNone(t)
		conf := SortConf{
			SortFuncs: SortFunctions{"Default Sort": SortTest},
			TestsDir:  []string{"testcases/random"},
			SizeCount: 2,
		}

		rt, err := RunTest(conf)
		require.NoError(t, err)
		require.Equal(t, false, rt["testcases/random"]["Default Sort"][0].Timeout)
		require.Equal(t, false, rt["testcases/random"]["Default Sort"][1].Timeout)
	})

	t.Run("error test", func(t *testing.T) {
		defer goleak.VerifyNone(t)
		conf := SortConf{
			SortFuncs: SortFunctions{"Default Sort": SortTest},
			TestsDir:  []string{"random"},
			SizeCount: 2,
		}

		_, err := RunTest(conf)
		require.Error(t, err)
	})
}
