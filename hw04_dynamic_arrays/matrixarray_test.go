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
