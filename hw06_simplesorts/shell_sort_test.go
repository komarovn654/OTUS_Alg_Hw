package hw06simplesorts

// func TestShellSort(t *testing.T) {
// 	tests := []int64{1, 10, 100, 1000, 10000}

// 	for _, tc := range tests {
// 		t.Run("N="+fmt.Sprint(tc), func(t *testing.T) {
// 			ar := sortutils.Array{}
// 			ar.InitRandArray(tc)

// 			st := ShellSort(ar)
// 			time := <-st
// 			require.Equal(t, false, time.Timeout)
// 			require.True(t, ar.IsSorted())
// 		})
// 	}
// }

// func TestShellSortFrankLazarus(t *testing.T) {
// 	tests := []int64{1, 10, 100, 1000, 10000}

// 	for _, tc := range tests {
// 		t.Run("N="+fmt.Sprint(tc), func(t *testing.T) {
// 			ar := sortutils.Array{}
// 			ar.InitRandArray(tc)

// 			st := ShellSortFrankLazarus(ar)
// 			time := <-st
// 			require.Equal(t, false, time.Timeout)
// 			require.True(t, ar.IsSorted())
// 		})
// 	}
// }

// func TestShellSortHibbard(t *testing.T) {
// 	tests := []int64{1, 10, 100, 1000, 10000}

// 	for _, tc := range tests {
// 		t.Run("N="+fmt.Sprint(tc), func(t *testing.T) {
// 			ar := sortutils.Array{}
// 			ar.InitRandArray(tc)

// 			st := ShellSortHibbard(ar)
// 			time := <-st
// 			require.Equal(t, false, time.Timeout)
// 			require.True(t, ar.IsSorted())
// 		})
// 	}
// }
