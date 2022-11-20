package hw04arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddDynamicArray(t *testing.T) {
	tests := []struct {
		name   string
		values []Item
		ref    DynamicArray
		expect DynamicArray
	}{
		{
			name:   "add items",
			values: []Item{15, 34, 425, 44, 345, 6, 98},
			ref:    DynamicArray{arr: []Item{}, len: 0, cap: 0},
			expect: DynamicArray{arr: []Item{15, 34, 425, 44, 345, 6, 98, nil, nil, nil, nil, nil}, len: 7, cap: 12},
		},
		{
			name:   "add items without new alloc",
			values: []Item{0, 1, 2},
			ref:    DynamicArray{arr: []Item{15, 34, 425, nil, nil, nil}, len: 3, cap: 6},
			expect: DynamicArray{arr: []Item{15, 34, 425, 0, 1, 2}, len: 6, cap: 6},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for _, v := range tc.values {
				tc.ref.Add(v)
			}
			require.Equal(t, tc.expect, tc.ref)
		})
	}
}

func TestInsertDinamycArray(t *testing.T) {
	tests := []struct {
		name   string
		value  Item
		index  int
		ref    DynamicArray
		expect DynamicArray
	}{
		{
			name:   "zero index",
			value:  15,
			index:  0,
			ref:    DynamicArray{arr: []Item{20, 21, 22, 23, 24, nil, nil, nil, nil, nil}, len: 5, cap: 10},
			expect: DynamicArray{arr: []Item{15, 20, 21, 22, 23, 24, nil, nil, nil, nil}, len: 6, cap: 10},
		},
		{
			name:   "middle index",
			value:  15,
			index:  2,
			ref:    DynamicArray{arr: []Item{20, 21, 22, 23, 24, nil, nil, nil, nil, nil}, len: 5, cap: 10},
			expect: DynamicArray{arr: []Item{20, 21, 15, 22, 23, 24, nil, nil, nil, nil}, len: 6, cap: 10},
		},
		{
			name:   "last index",
			value:  15,
			index:  4,
			ref:    DynamicArray{arr: []Item{20, 21, 22, 23, 24, nil, nil, nil, nil, nil}, len: 5, cap: 10},
			expect: DynamicArray{arr: []Item{20, 21, 22, 23, 15, 24, nil, nil, nil, nil}, len: 6, cap: 10},
		},
		{
			name:   "overcap",
			value:  15,
			index:  2,
			ref:    DynamicArray{arr: []Item{20, 21, 22}, len: 3, cap: 3},
			expect: DynamicArray{arr: []Item{20, 21, 15, 22, nil, nil, nil, nil, nil}, len: 4, cap: 9},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.ref.Insert(tc.value, tc.index)
			require.Equal(t, tc.expect, tc.ref)
		})
	}
}

func TestRemoveDinamycArray(t *testing.T) {
	tests := []struct {
		name    string
		index   int
		ref     DynamicArray
		expect  DynamicArray
		expectV Item
	}{
		{
			name:    "zero index",
			index:   0,
			ref:     DynamicArray{arr: []Item{20, 21, 22, 23, 24, nil, nil, nil, nil, nil}, len: 5, cap: 10},
			expect:  DynamicArray{arr: []Item{21, 22, 23, 24, 24, nil, nil, nil, nil, nil}, len: 4, cap: 10},
			expectV: 20,
		},
		{
			name:    "middle index",
			index:   2,
			ref:     DynamicArray{arr: []Item{20, 21, 22, 23, 24, nil, nil, nil, nil, nil}, len: 5, cap: 10},
			expect:  DynamicArray{arr: []Item{20, 21, 23, 24, 24, nil, nil, nil, nil, nil}, len: 4, cap: 10},
			expectV: 22,
		},
		{
			name:    "last index",
			index:   4,
			ref:     DynamicArray{arr: []Item{20, 21, 22, 23, 24, nil, nil, nil, nil, nil}, len: 5, cap: 10},
			expect:  DynamicArray{arr: []Item{20, 21, 22, 23, 24, nil, nil, nil, nil, nil}, len: 4, cap: 10},
			expectV: 24,
		},
		{
			name:    "last index, full array",
			index:   3,
			ref:     DynamicArray{arr: []Item{20, 21, 22, 23}, len: 4, cap: 4},
			expect:  DynamicArray{arr: []Item{20, 21, 22, 23}, len: 3, cap: 4},
			expectV: 23,
		},
		{
			name:    "last element",
			index:   0,
			ref:     DynamicArray{arr: []Item{20, 0}, len: 1, cap: 2},
			expect:  DynamicArray{arr: []Item{20, 0}, len: 0, cap: 2},
			expectV: 20,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.ref.Remove(tc.index)
			require.Equal(t, tc.expect, tc.ref)
		})
	}
}
