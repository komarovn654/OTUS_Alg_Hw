package hw10binarysearchtree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		name string
		keys []int // node's keys. key == value
	}{
		{
			name: "random array",
			keys: []int{49, 75, 61, 50, 93, 89, 34, 96, 40, 10, 14, 61, 99},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tree := bst{} // TODO Init
			for _, item := range tc.keys {
				tree.Insert(item)
			}
			require.True(t, tree.IsValid())
			// TODO проверять валидность дерева
		})
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name    string
		tree    bst
		isValid bool
	}{
		{
			name:    "valid tree",
			tree:    validBST,
			isValid: true,
		},
		{
			name:    "invalid tree",
			tree:    invalidBST,
			isValid: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.isValid, tc.tree.IsValid())
		})
	}
}
