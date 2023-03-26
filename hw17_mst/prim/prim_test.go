package prim

import (
	"testing"

	graph "github.com/komarovn654/OTUS_Alg_Hw/hw17_mst"
	"github.com/stretchr/testify/require"
)

func isExist(src []graph.Edge, edge graph.Edge) bool {
	reverse := graph.Edge{Weight: edge.Weight, Src: edge.Dst, Dst: edge.Src}

	for _, e := range src {
		if e == edge || e == reverse {
			return true
		}
	}

	return false
}

func TestFindMST(t *testing.T) {
	tests := []struct {
		vectors []graph.AdjacentVector
		mst     []graph.Edge
	}{
		{
			vectors: []graph.AdjacentVector{
				{{Num: 1, Weight: 1}, {Num: 3, Weight: 7}, {Num: 4, Weight: 3}},                      // 0
				{{Num: 0, Weight: 1}, {Num: 2, Weight: 4}, {Num: 3, Weight: 2}},                      // 1
				{{Num: 1, Weight: 4}, {Num: 3, Weight: 5}},                                           // 2
				{{Num: 0, Weight: 7}, {Num: 1, Weight: 2}, {Num: 2, Weight: 5}, {Num: 4, Weight: 6}}, // 3
				{{Num: 0, Weight: 3}, {Num: 3, Weight: 6}},                                           // 4
			},
			mst: []graph.Edge{
				{Weight: 4, Src: 2, Dst: 1},
				{Weight: 2, Src: 3, Dst: 1},
				{Weight: 3, Src: 4, Dst: 0},
				{Weight: 1, Src: 0, Dst: 1},
			},
		},
		{
			vectors: []graph.AdjacentVector{
				{{Num: 1, Weight: 9}, {Num: 3, Weight: 10}, {Num: 8, Weight: 3}},                                             // 0
				{{Num: 0, Weight: 9}, {Num: 2, Weight: 4}, {Num: 4, Weight: 8}, {Num: 8, Weight: 16}},                        // 1
				{{Num: 1, Weight: 4}, {Num: 4, Weight: 14}, {Num: 5, Weight: 1}},                                             // 2
				{{Num: 0, Weight: 10}, {Num: 4, Weight: 7}, {Num: 6, Weight: 5}, {Num: 7, Weight: 13}, {Num: 8, Weight: 11}}, // 3
				{{Num: 1, Weight: 8}, {Num: 2, Weight: 14}, {Num: 3, Weight: 7}, {Num: 5, Weight: 12}, {Num: 7, Weight: 15}}, // 4
				{{Num: 2, Weight: 1}, {Num: 4, Weight: 12}, {Num: 7, Weight: 2}},                                             // 5
				{{Num: 3, Weight: 5}, {Num: 7, Weight: 6}},                                                                   // 6
				{{Num: 3, Weight: 13}, {Num: 4, Weight: 15}, {Num: 5, Weight: 2}, {Num: 6, Weight: 6}},                       // 7
				{{Num: 0, Weight: 3}, {Num: 1, Weight: 16}, {Num: 3, Weight: 11}},                                            // 8
			},
			mst: []graph.Edge{
				{Weight: 6, Src: 7, Dst: 6},
				{Weight: 7, Src: 4, Dst: 3},
				{Weight: 5, Src: 6, Dst: 3},
				{Weight: 3, Src: 8, Dst: 0},
				{Weight: 2, Src: 7, Dst: 5},
				{Weight: 4, Src: 1, Dst: 2},
				{Weight: 1, Src: 5, Dst: 2},
				{Weight: 9, Src: 0, Dst: 1},
			},
		},
	}

	for _, tc := range tests {
		t.Run("common", func(t *testing.T) {
			g := graph.Init()
			for _, v := range tc.vectors {
				g.AddAdjacentVector(v)
			}

			mst := FindMST(&g)
			for _, edge := range mst {
				require.True(t, isExist(tc.mst, edge), "%+v doesn't match", edge)
			}
		})
	}
}
