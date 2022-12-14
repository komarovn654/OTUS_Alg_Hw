package hw06simplesorts

import (
	"context"
	"math"
	"time"

	"github.com/komarovn654/OTUS_Alg_Hw/sortutils"
)

func ShellSort(ctx context.Context, sTime chan<- sortutils.SortTime, array sortutils.Array) {
	start := time.Now()
	for gap := len(array.Ar) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(array.Ar); i++ {
			for j := i; j >= gap && (array.Ar)[j-gap] > (array.Ar)[j]; j -= gap {
				array.Swap(j, j-gap)

				select {
				case <-ctx.Done():
					sTime <- sortutils.SortTime{Timeout: true}
					return
				default:
				}
			}
		}
	}

	sTime <- sortutils.SortTime{Time: time.Since(start)}
}

func ShellSortFrankLazarus(ctx context.Context, sTime chan<- sortutils.SortTime, array sortutils.Array) {
	start := time.Now()
	gap := 2*(len(array.Ar)/4) + 1
	for k := 4; gap > 1; k *= 2 {
		gap = (2 * (len(array.Ar) / k)) + 1
		for i := gap; i < len(array.Ar); i++ {
			for j := i; j >= gap && (array.Ar)[j-gap] > (array.Ar)[j]; j -= gap {
				array.Swap(j, j-gap)

				select {
				case <-ctx.Done():
					sTime <- sortutils.SortTime{Timeout: true}
					return
				default:
				}
			}
		}
	}

	sTime <- sortutils.SortTime{Time: time.Since(start)}
}

func ShellSortHibbard(ctx context.Context, sTime chan<- sortutils.SortTime, array sortutils.Array) {
	start := time.Now()
	stepsNum := getSteps(len(array.Ar))
	for gap := int(math.Pow(2, float64(stepsNum))) - 1; gap > 0; gap /= 2 {
		for i := gap; i < len(array.Ar); i++ {
			for j := i; j >= gap && (array.Ar)[j-gap] > (array.Ar)[j]; j -= gap {
				array.Swap(j, j-gap)

				select {
				case <-ctx.Done():
					sTime <- sortutils.SortTime{Timeout: true}
					return
				default:
				}
			}
		}
	}

	sTime <- sortutils.SortTime{Time: time.Since(start)}
}

func getSteps(N int) (count int) {
	for i := 2; N > i; count++ {
		i = i * 2
	}
	return count
}
