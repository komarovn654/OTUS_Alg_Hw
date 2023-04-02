package bst

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const TREE_SIZE = 100_000

type testItem struct {
	key int
}

func (this *testItem) GetKey() int {
	return this.key
}

func initBSTIncreasing(len int) (bst, []testItem) {
	tree := InitBST()
	ar := make([]testItem, len)
	for i := 0; i < len; i++ {
		item := testItem{key: i}
		tree.Insert(&item)
		ar[i] = item
	}
	return tree, ar
}

func initBSTRandom(len int) (bst, []testItem) {
	ar := make([]testItem, len)
	tree := InitBST()
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < len; i++ {
		item := testItem{key: r1.Int()}
		tree.Insert(&item)
		ar[i] = item
	}
	return tree, ar
}

func initRandomArray(size int) []int {
	ar := make([]int, size)
	s1 := rand.NewSource(time.Now().UnixNano() + 1)
	r1 := rand.New(s1)

	for i := 0; i < size; i++ {
		ar[i] = int(r1.Int31())
	}
	return ar
}

// go test -run TestBST -timeout 600s -v
func TestBST(t *testing.T) {
	tests := []struct {
		name     string
		initFunc func(int) (bst, []testItem)
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
			t.Log("insert " + tc.name)
			start := time.Now()
			tree, values := tc.initFunc(TREE_SIZE)
			t.Logf("insert %v; len = %v; time = %v\n", tc.name, len(values), time.Since(start))
			require.True(t, tree.IsValid())

			t.Log("search " + tc.name)
			rnd := initRandomArray(len(values) / 10)
			start = time.Now()
			for _, v := range rnd {
				tree.Search(v)
			}
			t.Logf("search %v; len = %v; time: %v\n", tc.name, len(values)/10, time.Since(start))
			require.True(t, tree.IsValid())

			t.Log("remove " + tc.name)
			start = time.Now()
			i := 0
			for ; i < len(values); i += 10 {
				tree.Remove(i)
			}
			t.Logf("remove %v; len = %v; time: %v\n", tc.name, i/10, time.Since(start))
			require.True(t, tree.IsValid())
		})
	}
}

func TestInsertBST(t *testing.T) {
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
			for _, v := range tc.keys {
				item := testItem{key: v}
				tree.Insert(&item)
				_, ok := tree.Search(v)
				require.True(t, ok)
			}
			require.True(t, tree.IsValid())
		})
	}
}

func TestSearchBST(t *testing.T) {
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
				_, ok := tc.tree.Search(i)
				require.Equal(t, tc.expect, ok)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	var tree bst
	items := validBSTItems

	t.Run("remove tree", func(t *testing.T) {
		for _, v := range items {
			item := testItem{key: v}
			tree.Insert(&item)
		}

		for _, i := range items {
			_, ok := tree.Search(i)
			require.True(t, ok)

			tree.Remove(i)

			_, ok = tree.Search(i)
			require.False(t, ok)
			require.True(t, tree.IsValid())
		}
	})

	t.Run("remove items", func(t *testing.T) {
		for _, i := range items {
			for _, v := range items {
				item := testItem{key: v}
				tree.Insert(&item)
			}

			_, ok := tree.Search(i)
			require.True(t, ok)

			tree.Remove(i)

			_, ok = tree.Search(i)
			require.False(t, ok)
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
