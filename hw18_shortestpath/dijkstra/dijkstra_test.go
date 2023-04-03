package dijkstra

import (
	"testing"

	graph "github.com/komarovn654/OTUS_Alg_Hw/hw18_shortestpath"
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

func TestFindPath(t *testing.T) {
	type fromTo struct {
		src int
		dst int
	}

	type pathways struct {
		path   fromTo
		edges  []graph.Edge
		weight int
	}

	tests := []struct {
		vectors  []graph.AdjacentVector
		shortest []pathways
	}{
		{
			vectors: []graph.AdjacentVector{
				{{Dst: 1, Weight: 2}, {Dst: 2, Weight: 3}, {Dst: 3, Weight: 6}},                                           // 0
				{{Dst: 0, Weight: 2}, {Dst: 2, Weight: 4}, {Dst: 4, Weight: 9}},                                           // 1
				{{Dst: 0, Weight: 3}, {Dst: 1, Weight: 4}, {Dst: 3, Weight: 1}, {Dst: 4, Weight: 7}, {Dst: 5, Weight: 6}}, // 2
				{{Dst: 0, Weight: 6}, {Dst: 2, Weight: 1}, {Dst: 5, Weight: 4}},                                           // 3
				{{Dst: 1, Weight: 9}, {Dst: 2, Weight: 7}, {Dst: 5, Weight: 1}, {Dst: 6, Weight: 5}},                      // 4
				{{Dst: 2, Weight: 6}, {Dst: 3, Weight: 4}, {Dst: 4, Weight: 1}, {Dst: 6, Weight: 8}},                      // 5
				{{Dst: 4, Weight: 5}, {Dst: 5, Weight: 8}},                                                                // 6
			},
			shortest: []pathways{
				{path: fromTo{src: 0, dst: 0}, edges: []graph.Edge{{Weight: 0, Src: 0, Dst: 0}}, weight: 0},
				{path: fromTo{src: 0, dst: 1}, edges: []graph.Edge{{Weight: 2, Src: 0, Dst: 1}}, weight: 2},
				{path: fromTo{src: 0, dst: 2}, edges: []graph.Edge{{Weight: 3, Src: 0, Dst: 2}}, weight: 3},
				{path: fromTo{src: 0, dst: 3}, edges: []graph.Edge{{Weight: 3, Src: 0, Dst: 2}, {Weight: 1, Src: 2, Dst: 3}}, weight: 4},
				{path: fromTo{src: 0, dst: 4}, edges: []graph.Edge{{Weight: 3, Src: 0, Dst: 2}, {Weight: 7, Src: 2, Dst: 4}}, weight: 10},
				{path: fromTo{src: 0, dst: 5}, edges: []graph.Edge{{Weight: 3, Src: 0, Dst: 2}, {Weight: 1, Src: 2, Dst: 3}, {Weight: 4, Src: 3, Dst: 5}}, weight: 8},
				{path: fromTo{src: 0, dst: 6}, edges: []graph.Edge{{Weight: 3, Src: 0, Dst: 2}, {Weight: 7, Src: 2, Dst: 4}, {Weight: 5, Src: 4, Dst: 6}}, weight: 15},
				{path: fromTo{src: 1, dst: 0}, edges: []graph.Edge{{Weight: 2, Src: 1, Dst: 0}}, weight: 2},
				{path: fromTo{src: 1, dst: 1}, edges: []graph.Edge{{Weight: 0, Src: 1, Dst: 1}}, weight: 0},
				{path: fromTo{src: 1, dst: 2}, edges: []graph.Edge{{Weight: 4, Src: 1, Dst: 2}}, weight: 4},
				{path: fromTo{src: 1, dst: 3}, edges: []graph.Edge{{Weight: 4, Src: 1, Dst: 2}, {Weight: 1, Src: 2, Dst: 3}}, weight: 5},
				{path: fromTo{src: 1, dst: 4}, edges: []graph.Edge{{Weight: 9, Src: 1, Dst: 4}}, weight: 9},
				{path: fromTo{src: 1, dst: 5}, edges: []graph.Edge{{Weight: 4, Src: 1, Dst: 2}, {Weight: 1, Src: 2, Dst: 3}, {Weight: 4, Src: 3, Dst: 5}}, weight: 9},
				{path: fromTo{src: 1, dst: 6}, edges: []graph.Edge{{Weight: 9, Src: 1, Dst: 4}, {Weight: 5, Src: 4, Dst: 6}}, weight: 14},
				{path: fromTo{src: 2, dst: 0}, edges: []graph.Edge{{Weight: 3, Src: 2, Dst: 0}}, weight: 3},
				{path: fromTo{src: 2, dst: 1}, edges: []graph.Edge{{Weight: 4, Src: 2, Dst: 1}}, weight: 4},
				{path: fromTo{src: 2, dst: 2}, edges: []graph.Edge{{Weight: 0, Src: 2, Dst: 2}}, weight: 0},
				{path: fromTo{src: 2, dst: 3}, edges: []graph.Edge{{Weight: 1, Src: 2, Dst: 3}}, weight: 1},
				{path: fromTo{src: 2, dst: 4}, edges: []graph.Edge{{Weight: 7, Src: 2, Dst: 4}}, weight: 7},
				{path: fromTo{src: 2, dst: 5}, edges: []graph.Edge{{Weight: 1, Src: 2, Dst: 3}, {Weight: 4, Src: 3, Dst: 5}}, weight: 5},
				{path: fromTo{src: 2, dst: 6}, edges: []graph.Edge{{Weight: 7, Src: 2, Dst: 4}, {Weight: 5, Src: 4, Dst: 6}}, weight: 12},
				{path: fromTo{src: 3, dst: 0}, edges: []graph.Edge{{Weight: 6, Src: 3, Dst: 0}}, weight: 6},
				{path: fromTo{src: 3, dst: 1}, edges: []graph.Edge{{Weight: 1, Src: 3, Dst: 2}, {Weight: 4, Src: 2, Dst: 1}}, weight: 5},
				{path: fromTo{src: 3, dst: 2}, edges: []graph.Edge{{Weight: 1, Src: 3, Dst: 2}}, weight: 1},
				{path: fromTo{src: 3, dst: 3}, edges: []graph.Edge{{Weight: 0, Src: 3, Dst: 3}}, weight: 0},
				{path: fromTo{src: 3, dst: 4}, edges: []graph.Edge{{Weight: 4, Src: 3, Dst: 5}, {Weight: 1, Src: 5, Dst: 4}}, weight: 5},
				{path: fromTo{src: 3, dst: 5}, edges: []graph.Edge{{Weight: 4, Src: 3, Dst: 5}}, weight: 4},
				{path: fromTo{src: 3, dst: 6}, edges: []graph.Edge{{Weight: 4, Src: 3, Dst: 5}, {Weight: 1, Src: 5, Dst: 4}, {Weight: 5, Src: 4, Dst: 6}}, weight: 10},
				{path: fromTo{src: 4, dst: 0}, edges: []graph.Edge{{Weight: 7, Src: 4, Dst: 2}, {Weight: 3, Src: 2, Dst: 0}}, weight: 10},
				{path: fromTo{src: 4, dst: 1}, edges: []graph.Edge{{Weight: 9, Src: 4, Dst: 1}}, weight: 9},
				{path: fromTo{src: 4, dst: 2}, edges: []graph.Edge{{Weight: 7, Src: 4, Dst: 2}}, weight: 7},
				{path: fromTo{src: 4, dst: 3}, edges: []graph.Edge{{Weight: 1, Src: 4, Dst: 5}, {Weight: 4, Src: 5, Dst: 3}}, weight: 5},
				{path: fromTo{src: 4, dst: 4}, edges: []graph.Edge{{Weight: 0, Src: 4, Dst: 4}}, weight: 0},
				{path: fromTo{src: 4, dst: 5}, edges: []graph.Edge{{Weight: 1, Src: 4, Dst: 5}}, weight: 1},
				{path: fromTo{src: 4, dst: 6}, edges: []graph.Edge{{Weight: 5, Src: 4, Dst: 6}}, weight: 5},
				{path: fromTo{src: 4, dst: 0}, edges: []graph.Edge{{Weight: 7, Src: 4, Dst: 2}, {Weight: 3, Src: 2, Dst: 0}}, weight: 10},
				{path: fromTo{src: 5, dst: 1}, edges: []graph.Edge{{Weight: 6, Src: 5, Dst: 2}, {Weight: 4, Src: 2, Dst: 1}}, weight: 10},
				{path: fromTo{src: 5, dst: 2}, edges: []graph.Edge{{Weight: 6, Src: 5, Dst: 2}}, weight: 6},
				{path: fromTo{src: 5, dst: 3}, edges: []graph.Edge{{Weight: 4, Src: 5, Dst: 3}}, weight: 4},
				{path: fromTo{src: 5, dst: 4}, edges: []graph.Edge{{Weight: 1, Src: 5, Dst: 4}}, weight: 1},
				{path: fromTo{src: 5, dst: 5}, edges: []graph.Edge{{Weight: 0, Src: 5, Dst: 5}}, weight: 0},
				{path: fromTo{src: 5, dst: 6}, edges: []graph.Edge{{Weight: 1, Src: 5, Dst: 4}, {Weight: 5, Src: 4, Dst: 6}}, weight: 6},
				{path: fromTo{src: 6, dst: 0}, edges: []graph.Edge{{Weight: 5, Src: 6, Dst: 4}, {Weight: 7, Src: 4, Dst: 2}, {Weight: 3, Src: 2, Dst: 0}}, weight: 15},
				{path: fromTo{src: 6, dst: 1}, edges: []graph.Edge{{Weight: 5, Src: 6, Dst: 4}, {Weight: 9, Src: 4, Dst: 1}}, weight: 14},
				{path: fromTo{src: 6, dst: 2}, edges: []graph.Edge{{Weight: 5, Src: 6, Dst: 4}, {Weight: 7, Src: 4, Dst: 2}}, weight: 12},
				{path: fromTo{src: 6, dst: 3}, edges: []graph.Edge{{Weight: 5, Src: 6, Dst: 4}, {Weight: 1, Src: 4, Dst: 5}, {Weight: 4, Src: 5, Dst: 3}}, weight: 10},
				{path: fromTo{src: 6, dst: 4}, edges: []graph.Edge{{Weight: 5, Src: 6, Dst: 4}}, weight: 5},
				{path: fromTo{src: 6, dst: 5}, edges: []graph.Edge{{Weight: 5, Src: 6, Dst: 4}, {Weight: 1, Src: 4, Dst: 5}}, weight: 6},
				{path: fromTo{src: 6, dst: 6}, edges: []graph.Edge{{Weight: 0, Src: 6, Dst: 6}}, weight: 0},
			},
		},
	}

	t.Run("common", func(t *testing.T) {
		for _, tc := range tests {
			g := graph.Init()
			for _, v := range tc.vectors {
				g.AddAdjacentVector(v)
			}

			for _, sh := range tc.shortest {
				edges := FindPath(&g, sh.path.src, sh.path.dst)
				require.Equal(t, sh.edges, edges)

				weight := 0
				for _, edge := range edges {
					weight += edge.Weight
				}
				// fmt.Printf("%+v, weight: %v},\n", edges, weight)
				require.Equal(t, sh.weight, weight)
			}
		}
	})
}
