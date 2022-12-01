package hw04arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddSingleArray(t *testing.T) {
	tests := []struct {
		name   string
		values []Item
		ref    SingleArray
		expect SingleArray
	}{
		{
			name:   "add items",
			values: []Item{15, 34, 425, 44, 345},
			ref:    SingleArray{arr: []Item{}, len: 0},
			expect: SingleArray{arr: []Item{15, 34, 425, 44, 345}, len: 5},
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

func TestInsertSingleArray(t *testing.T) {
	tests := []struct {
		name   string
		value  Item
		index  int
		ref    SingleArray
		expect SingleArray
	}{
		{
			name:   "zero index",
			value:  15,
			index:  0,
			ref:    SingleArray{arr: []Item{20, 21, 22, 23, 24}, len: 5},
			expect: SingleArray{arr: []Item{15, 20, 21, 22, 23, 24}, len: 6},
		},
		{
			name:   "middle index",
			value:  15,
			index:  2,
			ref:    SingleArray{arr: []Item{20, 21, 22, 23, 24}, len: 5},
			expect: SingleArray{arr: []Item{20, 21, 15, 22, 23, 24}, len: 6},
		},
		{
			name:   "last index",
			value:  15,
			index:  4,
			ref:    SingleArray{arr: []Item{20, 21, 22, 23, 24}, len: 5},
			expect: SingleArray{arr: []Item{20, 21, 22, 23, 15, 24}, len: 6},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.ref.Insert(tc.value, tc.index)
			require.Equal(t, tc.expect, tc.ref)
		})
	}
}

func TestRemoveSingleArray(t *testing.T) {
	tests := []struct {
		name    string
		index   int
		ref     SingleArray
		expectA SingleArray
		expectV Item
	}{
		{
			name:    "zero index",
			index:   0,
			ref:     SingleArray{arr: []Item{20, 21, 22, 23, 24}, len: 5},
			expectA: SingleArray{arr: []Item{21, 22, 23, 24}, len: 4},
			expectV: 20,
		},
		{
			name:    "middle index",
			index:   2,
			ref:     SingleArray{arr: []Item{20, 21, 22, 23, 24}, len: 5},
			expectA: SingleArray{arr: []Item{20, 21, 23, 24}, len: 4},
			expectV: 22,
		},
		{
			name:    "last index",
			index:   4,
			ref:     SingleArray{arr: []Item{20, 21, 22, 23, 24}, len: 5},
			expectA: SingleArray{arr: []Item{20, 21, 22, 23}, len: 4},
			expectV: 24,
		},
		{
			name:    "last element",
			index:   0,
			ref:     SingleArray{arr: []Item{20}, len: 1},
			expectA: SingleArray{arr: []Item{}, len: 0},
			expectV: 20,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			value := tc.ref.Remove(tc.index)
			require.Equal(t, tc.expectA, tc.ref)
			require.Equal(t, tc.expectV, value)
		})
	}
}
