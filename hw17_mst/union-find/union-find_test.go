package unionfind

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	// 0->1 1->1 2->3 3->1 4->4 5->4 9->0 6->7
	vertices := []int{0, 1, 2, 3, 4, 6, 7, 8, 9}

	pairs := Init(vertices)
	for _, vertex := range vertices {
		require.Equal(t, vertex, pairs.Find(vertex))
	}
}

func TestAddPair(t *testing.T) {
	// 0->1 1->1 2->3 3->1 4->4 5->4 9->0 6->7
	pairs := []struct {
		vert int
		root int
	}{
		{0, 1},
		{1, 1},
		{2, 3},
		{3, 1},
		{4, 4},
		{5, 4},
		{9, 0},
		{6, 7},
	}

	uf := Init(nil)
	for _, p := range pairs {
		uf.AddPair(p.vert, p.root)
	}

	res := uf.GetPairs()
	for _, p := range pairs {
		root, ok := res[p.vert]
		require.True(t, ok)
		require.Equal(t, p.root, root)
	}
}

func TestFind(t *testing.T) {
	// 0->1 1->1 2->3 3->1 4->4 5->4 9->0 6->7
	pairs := []struct {
		vert      int
		root      int
		superroot int
	}{
		{0, 1, 1},
		{1, 1, 1},
		{2, 3, 1},
		{3, 1, 1},
		{4, 4, 4},
		{5, 4, 4},
		{9, 0, 1},
		{6, 7, -1},
	}

	uf := Init(nil)
	for _, p := range pairs {
		uf.AddPair(p.vert, p.root)
	}

	for _, p := range pairs {
		require.Equal(t, p.superroot, uf.Find(p.vert))
	}

	for _, p := range pairs {
		require.Equal(t, p.superroot, uf.Find(p.vert))
	}
}

func TestUnion(t *testing.T) {
	// 0->1 1->1 2->3 3->1 4->4 5->4 9->0 6->7
	pairs := []struct {
		vert int
		root int
	}{
		{0, 1},
		{1, 1},
		{2, 3},
		{3, 1},
		{4, 4},
		{5, 4},
		{9, 0},
		{6, 4},
	}

	uf := Init(nil)
	for _, p := range pairs {
		uf.AddPair(p.vert, p.root)
	}

	root := uf.Union(2, 5)
	require.Equal(t, uf.Find(2), root)
	require.Equal(t, root, uf.Find(5))

	root = uf.Union(0, 0)
	require.Equal(t, uf.Find(0), root)

	require.Equal(t, uf.Find(6), uf.Find(4))
	uf.Union(1, 4)
	require.Equal(t, uf.Find(4), uf.Find(1))
	require.Equal(t, uf.Find(6), uf.Find(1))
}
