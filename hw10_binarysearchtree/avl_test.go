package hw10binarysearchtree

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func initAVLIncreasing(len int) (avl, []int) {
	tree := InitAVL()
	ar := make([]int, len)
	for i := 0; i < len; i++ {
		tree.Insert(i)
		ar[i] = i
	}
	return tree, ar
}

func initAVLRandom(len int) (avl, []int) {
	ar := make([]int, len)
	tree := InitAVL()
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < len; i++ {
		v := int(r1.Int())
		tree.Insert(v)
		ar[i] = v
	}
	return tree, ar
}

// go test -run TestAVL -timeout 600s -v
func TestAVL(t *testing.T) {
	tests := []struct {
		name     string
		initFunc func(int) (avl, []int)
	}{
		{
			name:     "increasing AVL",
			initFunc: initAVLIncreasing,
		},
		{
			name:     "random AVL",
			initFunc: initAVLRandom,
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
				tree.Remove(values[i])
			}
			t.Logf("remove %v; len = %v; time: %v\n", tc.name, i/10, time.Since(start))
			require.True(t, tree.IsValid())
		})
	}
}

func TestSmallRightRotationAVL(t *testing.T) {
	tree := UnbalanceSmallRight

	newRoot := tree.root.smallRightRotation()
	require.Equal(t, 34, newRoot.item.key)
	require.Equal(t, 49, newRoot.right.item.key)
	tree = avl{newRoot}
	require.True(t, tree.IsValid())
}

func TestBigRightRotationAVL(t *testing.T) {
	tree := UnbalanceBigRight

	newRoot := tree.root.bigRightRotation()
	require.Equal(t, 55, newRoot.item.key)
	require.Equal(t, 80, newRoot.right.item.key)
	tree = avl{newRoot}
	require.True(t, tree.IsValid())
}