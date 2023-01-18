package hw10binarysearchtree

import (
	"fmt"
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
		})
	}
}

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		tree   bst
		x      []int
		expect bool
	}{
		{
			name:   "existing items",
			tree:   validBST,
			x:      validBSTItems,
			expect: true,
		},
		{
			name:   "non-existent items",
			tree:   validBST,
			x:      []int{0, 100, 87, 9},
			expect: false,
		},
		{
			name:   "empty tree",
			tree:   bst{},
			x:      []int{0, 100, 87, 9},
			expect: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for _, i := range tc.x {
				require.Equal(t, tc.expect, tc.tree.Search(i))
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tree := validBST
	tree.Insert(12)

	fmt.Println(tree.Search(12))
	tree.Remove(12)
	fmt.Println(tree.IsValid())
	fmt.Println(tree.Search(12))

	// fmt.Println(tree.Search(14))
	// tree.Remove(14)
	// fmt.Println(tree.IsValid())
	// fmt.Println(tree.Search(14))

	tree.Remove(34)
	fmt.Println(tree.IsValid())
	fmt.Println(tree.Search(34))
}

func TestFindMax(t *testing.T) {
	tree := validBST
	tree.Insert(15)
	tree.Insert(17)
	tree.Insert(16)
	tree.Insert(19)
	fmt.Println(tree.findMax(tree.root.left.left))
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
