package hw09linearsort

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// go test -run TestBucketSort -timeout 120s -v
func TestBucketSort(t *testing.T) {
	max := uint16(999)
	tests := []struct {
		len int
	}{
		{len: 10},
		{len: 100},
		{len: 1000},
		{len: 10000},
		{len: 100000},
		{len: 1000000},
	}

	for _, tc := range tests {
		t.Run("Length = "+strconv.Itoa(tc.len), func(t *testing.T) {
			array := GenerateArray(tc.len, max)

			tm := time.Now()
			sorted := BucketSort(array, max)
			fmt.Println("Sorted time = ", time.Since(tm))
			require.True(t, IsSorted(sorted))
		})
	}
}

func TestBucket(t *testing.T) {
	array := []uint16{868, 285, 508, 957, 496, 488, 715, 435, 185, 496}
	expect := []uint16{185, 285, 435, 488, 496, 496, 508, 715, 868, 957}
	t.Run("Bucket Insert/Get test", func(t *testing.T) {
		bucket := bucket{}
		for _, v := range array {
			bucket.insert(v)
		}

		require.Equal(t, expect, bucket.getAll())
	})

}
