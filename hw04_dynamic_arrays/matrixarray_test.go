package hw04arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddMatrixArray(t *testing.T) {
	tests := []struct {
		name   string
		values []Item
		ref    MatrixArray
		expect MatrixArray
	}{
		{
			name:   "add values, empty array",
			values: []Item{1, 21, 234},
			ref: MatrixArray{
				domains: FactorArray{},
				len:     0,
				maxLen:  10,
			},
			expect: MatrixArray{
				domains: FactorArray{
					arr: []Item{
						FactorArray{arr: []Item{1, 21, 234, nil, nil, nil}, len: 3, cap: 6},
						nil,
					},
					len: 1,
					cap: 2,
				},
				len:    3,
				maxLen: 10,
			},
		},
		{
			name:   "add values, new domain",
			values: []Item{534, 6, 3},
			ref: MatrixArray{
				domains: FactorArray{
					arr: []Item{
						FactorArray{arr: []Item{1, 21, 234, nil, nil, nil}, len: 3, cap: 6},
						FactorArray{arr: []Item{53, 75, 2, nil, nil, nil}, len: 3, cap: 6},
					},
					len: 2,
					cap: 2,
				},
				len:    6,
				maxLen: 3,
			},
			expect: MatrixArray{
				domains: FactorArray{
					arr: []Item{
						FactorArray{arr: []Item{1, 21, 234, nil, nil, nil}, len: 3, cap: 6},
						FactorArray{arr: []Item{53, 75, 2, nil, nil, nil}, len: 3, cap: 6},
						FactorArray{arr: []Item{534, 6, 3, nil, nil, nil}, len: 3, cap: 6},
						nil,
						nil,
						nil,
					},
					len: 3,
					cap: 6,
				},
				len:    9,
				maxLen: 3,
			},
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

func TestInsertMatrixArray(t *testing.T) {
	tests := []struct {
		name   string
		value  Item
		index  int
		ref    MatrixArray
		expect MatrixArray
	}{
		{
			name:  "insert value, middle",
			value: 55,
			index: 4,
			ref: MatrixArray{
				domains: FactorArray{
					arr: []Item{
						FactorArray{arr: []Item{0, 1, 2, nil, nil, nil}, len: 3, cap: 6},
						FactorArray{arr: []Item{3, 4, 5, nil, nil, nil}, len: 3, cap: 6},
						FactorArray{arr: []Item{6, 7, 8, nil, nil, nil}, len: 3, cap: 6},
						FactorArray{arr: []Item{9, 10, 11, nil, nil, nil}, len: 3, cap: 6},
						FactorArray{arr: []Item{12, 13}, len: 2, cap: 2},
						nil,
					},
					len: 5,
					cap: 6,
				},
				len:    13,
				maxLen: 3,
			},
			expect: MatrixArray{
				domains: FactorArray{
					arr: []Item{
						FactorArray{arr: []Item{0, 1, 2, nil, nil, nil}, len: 3, cap: 6},
						FactorArray{arr: []Item{3, 55, 4, 5, nil, nil}, len: 3, cap: 6},
						FactorArray{arr: []Item{5, 6, 7, 8, nil, nil}, len: 3, cap: 6},
						FactorArray{arr: []Item{8, 9, 10, 11, nil, nil}, len: 3, cap: 6},
						FactorArray{arr: []Item{11, 12, 13, nil, nil, nil}, len: 3, cap: 6},
						nil,
					},
					len: 5,
					cap: 6,
				},
				len:    14,
				maxLen: 3,
			},
		},
		{
			name:  "insert value, overcap",
			value: 55,
			index: 2,
			ref: MatrixArray{
				domains: FactorArray{
					arr: []Item{
						FactorArray{arr: []Item{0, 1, 2, nil, nil, nil}, len: 3, cap: 6},
					},
					len: 1,
					cap: 1,
				},
				len:    3,
				maxLen: 3,
			},
			expect: MatrixArray{
				domains: FactorArray{
					arr: []Item{
						FactorArray{arr: []Item{0, 1, 55, 2, nil, nil}, len: 3, cap: 6},
						FactorArray{arr: []Item{2, nil}, len: 1, cap: 2},
						nil,
						nil,
					},
					len: 2,
					cap: 4,
				},
				len:    4,
				maxLen: 3,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.ref.Insert(tc.value, tc.index)
			require.Equal(t, tc.expect, tc.ref)
		})
	}
}

func TestRemoveMatrixArray(t *testing.T) {
	tests := []struct {
		name    string
		index   int
		ref     MatrixArray
		expect  MatrixArray
		exepctV Item
	}{
		{
			name:  "remove value, middle",
			index: 4,
			ref: MatrixArray{
				domains: FactorArray{
					arr: []Item{
						FactorArray{arr: []Item{0, 1, 2, nil, nil, nil}, len: 3, cap: 6},
						FactorArray{arr: []Item{3, 4, 5, nil, nil, nil}, len: 3, cap: 6},
						FactorArray{arr: []Item{6, 7, 8, nil, nil, nil}, len: 3, cap: 6},
						FactorArray{arr: []Item{9, nil}, len: 1, cap: 2},
					},
					len: 4,
					cap: 4,
				},
				len:    10,
				maxLen: 3,
			},
			expect: MatrixArray{
				domains: FactorArray{
					arr: []Item{
						FactorArray{arr: []Item{0, 1, 2, nil, nil, nil}, len: 3, cap: 6},
						FactorArray{arr: []Item{3, 5, 6, nil, nil, nil}, len: 3, cap: 6},
						FactorArray{arr: []Item{7, 8, 9, nil, nil, nil}, len: 3, cap: 6},
						FactorArray{arr: []Item{9, nil}, len: 0, cap: 2},
					},
					len: 4,
					cap: 4,
				},
				len:    9,
				maxLen: 3,
			},
			exepctV: 4,
		},
		{
			name:  "single domain",
			index: 0,
			ref: MatrixArray{
				domains: FactorArray{
					arr: []Item{
						FactorArray{arr: []Item{0, 1, 2, nil, nil, nil}, len: 3, cap: 6},
						nil,
					},
					len: 1,
					cap: 2,
				},
				len:    3,
				maxLen: 3,
			},
			expect: MatrixArray{
				domains: FactorArray{
					arr: []Item{
						FactorArray{arr: []Item{1, 2, 2, nil, nil, nil}, len: 2, cap: 6},
						nil,
					},
					len: 1,
					cap: 2,
				},
				len:    2,
				maxLen: 3,
			},
			exepctV: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			removed := tc.ref.Remove(tc.index)
			require.Equal(t, tc.exepctV, removed)
			require.Equal(t, tc.expect, tc.ref)
		})
	}
}
