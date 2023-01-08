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
			array := GenerateArray(tc.len)

			tm := time.Now()
			sorted := BucketSort(array)
			fmt.Println("Sorted time = ", time.Since(tm))
			require.True(t, IsSorted(sorted))
		})
	}
}

func TestStoreValue(t *testing.T) {
	array := []int64{108, 277, 463, 673}
	tests := []struct {
		v      int64
		array  []int64
		expect []int64
	}{
		{
			v:      90,
			array:  array,
			expect: []int64{90, 108, 277, 463, 673},
		},
		{
			v:      150,
			array:  array,
			expect: []int64{108, 150, 277, 463, 673},
		},
		{
			v:      300,
			array:  array,
			expect: []int64{108, 277, 300, 463, 673},
		},
		{
			v:      500,
			array:  array,
			expect: []int64{108, 277, 463, 500, 673},
		},
		{
			v:      700,
			array:  array,
			expect: []int64{108, 277, 463, 673, 700},
		},
	}

	for i, tc := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			buckets := initBuckets(10)
			for _, v := range tc.array {
				buckets.buck[0].PushBack(v)
			}

			buckets.storeValue(tc.v, 0)
			require.Equal(t, tc.expect, buckets.getBucketArray(0))
		})
	}
}
