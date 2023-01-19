package hw10binarysearchtree

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	generateTime = time.Second * 60
)

func initBSTIncreasing() (bst, []int) {
	tree := InitBST()
	ar := make([]int, 0)
	t := time.Now()
	for i := 0; time.Since(t) < generateTime; i++ {
		tree.Insert(i)
		ar = append(ar, i)
	}
	return tree, ar
}

func initBSTRandom() (bst, []int) {
	ar := make([]int, 0)
	tree := InitBST()
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	t := time.Now()
	for i := 0; time.Since(t) < generateTime; i++ {
		tree.Insert(r1.Int())
		ar = append(ar, i)
	}
	return tree, ar
}

func initRandomArray(size int) []int {
	ar := make([]int, size)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < size; i++ {
		ar[i] = r1.Int()
	}
	return ar
}

// go test -run TestBST -timeout 600s -v
func TestBST(t *testing.T) {
	tests := []struct {
		name     string
		initFunc func() (bst, []int)
	}{
		{
			name:     "increasing BST",
			initFunc: initBSTIncreasing,
		},
		{
			name:     "random BST",
			initFunc: initBSTRandom,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Log("generate " + tc.name)
			tree, values := tc.initFunc()
			require.True(t, tree.IsValid())

			t.Log("search " + tc.name)
			rnd := initRandomArray(len(values) / 10)
			for _, v := range rnd {
				tree.Search(v)
			}
			require.True(t, tree.IsValid())

			t.Log("remove " + tc.name)
			for _, v := range rnd {
				tree.Remove(v)
			}
			require.True(t, tree.IsValid())
		})
	}
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
		{
			name:   "random tree",
			tree:   bst{},
			x:      []int{},
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
