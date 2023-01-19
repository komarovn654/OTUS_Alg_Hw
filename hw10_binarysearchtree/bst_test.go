package hw10binarysearchtree

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func initBSTIncreasing() (bst, int) {
	tree := InitBST()
	t := time.Now()
	i := 0
	for ; time.Since(t) < time.Second*60; i++ {
		tree.Insert(i)
	}
	return tree, i
}

func initBSTRandom() bst {
	tree := InitBST()
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	t := time.Now()
	for time.Since(t) < 60*time.Duration(t.Second()) {
		tree.Insert(r1.Int())
	}
	return tree
}

// go test -run TestBST -timeout 600s -v
func TestBST(t *testing.T) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	t.Run("increasing BST", func(t *testing.T) {
		tree, len := initBSTIncreasing()
		require.True(t, tree.IsValid())
		t.Log("BST successfully initialized")
		for i := 0; i < len/10; i++ {
			tree.Search(r1.Int())
		}
	})

}

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
			tree := InitBST()
			for _, item := range tc.keys {
				tree.Insert(item)
				require.True(t, tree.Search(item))
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
	items := validBSTItems

	t.Run("remove items", func(t *testing.T) {
		for _, i := range items {
			require.True(t, tree.Search(i))
			tree.Remove(i)
			require.False(t, tree.Search(i))
			require.True(t, tree.IsValid())
		}
	})
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
