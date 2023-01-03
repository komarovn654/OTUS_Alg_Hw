package hw06simplesorts

// func TestInsertionSort(t *testing.T) {
// 	tests := []int64{1, 10, 100, 1000, 10000}

// 	for _, tc := range tests {
// 		t.Run("N="+fmt.Sprint(tc), func(t *testing.T) {
// 			ar := sortutils.Array{}
// 			ar.InitRandArray(tc)

// 			st := InsertionSort(ar)
// 			time := <-st
// 			require.Equal(t, false, time.Timeout)
// 			require.True(t, ar.IsSorted())
// 		})
// 	}
// }

// func TestInsertionSortShift(t *testing.T) {
// 	tests := []int64{1, 10, 100, 1000, 10000}

// 	for _, tc := range tests {
// 		t.Run("N="+fmt.Sprint(tc), func(t *testing.T) {
// 			ar := sortutils.Array{}
// 			ar.InitRandArray(tc)

// 			st := InsertionSortShift(ar)
// 			time := <-st
// 			require.Equal(t, false, time.Timeout)
// 			require.True(t, ar.IsSorted())
// 		})
// 	}
// }

// func TestInsertionSortBinarySearch(t *testing.T) {
// 	tests := []int64{1, 10, 100, 1000, 10000}

// 	for _, tc := range tests {
// 		t.Run("N="+fmt.Sprint(tc), func(t *testing.T) {
// 			ar := sortutils.Array{}
// 			ar.InitRandArray(tc)

// 			st := InsertionSortBinarySearch(ar)
// 			time := <-st
// 			require.Equal(t, false, time.Timeout)
// 			require.True(t, ar.IsSorted())
// 		})
// 	}
// }
