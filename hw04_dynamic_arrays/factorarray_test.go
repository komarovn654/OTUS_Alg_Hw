package hw04arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddFactorArray(t *testing.T) {
	tests := []struct {
		name   string
		values []Item
		ref    FactorArray
		expect FactorArray
	}{
		{
			name:   "add items",
			values: []Item{15, 34, 425, 44, 345},
			ref:    FactorArray{arr: []Item{}, len: 0, cap: 0},
			expect: FactorArray{arr: []Item{15, 34, 425, 44, 345, nil}, len: 5, cap: 6},
		},
		{
			name:   "add items without new alloc",
			values: []Item{0, 1, 2},
			ref:    FactorArray{arr: []Item{15, 34, 425, 0, 0, 0}, len: 3, cap: 6},
			expect: FactorArray{arr: []Item{15, 34, 425, 0, 1, 2}, len: 6, cap: 6},
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

func TestInsertFactorArray(t *testing.T) {
	tests := []struct {
		name   string
		value  Item
		index  int
		ref    FactorArray
		expect FactorArray
	}{
		{
			name:   "zero index",
			value:  15,
			index:  0,
			ref:    FactorArray{arr: []Item{20, 21, 22, 23, 24, nil, nil, nil, nil, nil}, len: 5, cap: 10},
			expect: FactorArray{arr: []Item{15, 20, 21, 22, 23, 24, nil, nil, nil, nil}, len: 6, cap: 10},
		},
		{
			name:   "middle index",
			value:  15,
			index:  2,
			ref:    FactorArray{arr: []Item{20, 21, 22, 23, 24, nil, nil, nil, nil, nil}, len: 5, cap: 10},
			expect: FactorArray{arr: []Item{20, 21, 15, 22, 23, 24, nil, nil, nil, nil}, len: 6, cap: 10},
		},
		{
			name:   "last index",
			value:  15,
			index:  4,
			ref:    FactorArray{arr: []Item{20, 21, 22, 23, 24, nil, nil, nil, nil, nil}, len: 5, cap: 10},
			expect: FactorArray{arr: []Item{20, 21, 22, 23, 15, 24, nil, nil, nil, nil}, len: 6, cap: 10},
		},
		{
			name:   "overcap",
			value:  15,
			index:  2,
			ref:    FactorArray{arr: []Item{20, 21, 22}, len: 3, cap: 3},
			expect: FactorArray{arr: []Item{20, 21, 15, 22, nil, nil, nil, nil}, len: 4, cap: 8},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.ref.Insert(tc.value, tc.index)
			require.Equal(t, tc.expect, tc.ref)
		})
	}
}

func TestRemoveFactorArray(t *testing.T) {
	tests := []struct {
		name    string
		index   int
		ref     FactorArray
		expect  FactorArray
		expectV Item
	}{
		{
			name:    "zero index",
			index:   0,
			ref:     FactorArray{arr: []Item{20, 21, 22, 23, 24, nil, nil, nil, nil, nil}, len: 5, cap: 10},
			expect:  FactorArray{arr: []Item{21, 22, 23, 24, 24, nil, nil, nil, nil, nil}, len: 4, cap: 10},
			expectV: 20,
		},
		{
			name:    "middle index",
			index:   2,
			ref:     FactorArray{arr: []Item{20, 21, 22, 23, 24, nil, nil, nil, nil, nil}, len: 5, cap: 10},
			expect:  FactorArray{arr: []Item{20, 21, 23, 24, 24, nil, nil, nil, nil, nil}, len: 4, cap: 10},
			expectV: 22,
		},
		{
			name:    "last index",
			index:   4,
			ref:     FactorArray{arr: []Item{20, 21, 22, 23, 24, nil, nil, nil, nil, nil}, len: 5, cap: 10},
			expect:  FactorArray{arr: []Item{20, 21, 22, 23, 24, nil, nil, nil, nil, nil}, len: 4, cap: 10},
			expectV: 24,
		},
		{
			name:    "last index, full array",
			index:   3,
			ref:     FactorArray{arr: []Item{20, 21, 22, 23}, len: 4, cap: 4},
			expect:  FactorArray{arr: []Item{20, 21, 22, 23}, len: 3, cap: 4},
			expectV: 23,
		},
		{
			name:    "last element",
			index:   0,
			ref:     FactorArray{arr: []Item{20, nil}, len: 1, cap: 2},
			expect:  FactorArray{arr: []Item{20, nil}, len: 0, cap: 2},
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
