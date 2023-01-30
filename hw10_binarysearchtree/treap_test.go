package hw10binarysearchtree

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func initTreapIncreasing(len int) (treap, []int) {
	tree := InitTreap()
	ar := make([]int, len)
	for i := 0; i < len; i++ {
		tree.Insert(i)
		ar[i] = i
	}
	return tree, ar
}

func initTreapRandom(len int) (treap, []int) {
	ar := make([]int, len)
	tree := InitTreap()
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < len; i++ {
		v := int(r1.Int())
		tree.Insert(v)
		ar[i] = v
	}
	return tree, ar
}

// go test -run TestTreap -timeout 600s -v
func TestTreap(t *testing.T) {
	tests := []struct {
		name     string
		initFunc func(int) (treap, []int)
	}{
		{
			name:     "increasing Treap",
			initFunc: initTreapIncreasing,
		},
		{
			name:     "random Treap",
			initFunc: initTreapRandom,
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
func TestSplit(t *testing.T) {
	tree := splitTree
	l, r := tree.split(7)

	require.Less(t, l.findMax().item.key, 7)
	require.GreaterOrEqual(t, r.findMin().item.key, 7)
}

func TestInsert(t *testing.T) {
	tree := treap{
		&treapNode{
			item:     nodeItem{key: 374},
			priority: 8594448610203129222,
			left: &treapNode{
				item:     nodeItem{key: 68},
				priority: 6526000971620458038,
				left:     nil,
				right:    nil,
			},
			right: &treapNode{
				item:     nodeItem{key: 495},
				priority: 8230938549176605455,
				left: &treapNode{
					item:     nodeItem{key: 387},
					priority: 1747550616550721433,
					left:     nil,
					right:    nil,
				},
				right: nil,
			},
		},
	}
	tree.Insert(48)
}
