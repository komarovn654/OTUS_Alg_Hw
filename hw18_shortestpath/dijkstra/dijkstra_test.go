package dijkstra

import (
	"fmt"
	"testing"

	graph "github.com/komarovn654/OTUS_Alg_Hw/hw18_shortestpath"
	"github.com/stretchr/testify/require"
)

func TestFind(t *testing.T) {
	vectors := []graph.AdjacentVector{
		{{Dst: 1, Weight: 2}, {Dst: 2, Weight: 3}, {Dst: 3, Weight: 6}},                                           // 0
		{{Dst: 0, Weight: 2}, {Dst: 2, Weight: 4}, {Dst: 4, Weight: 9}},                                           // 1
		{{Dst: 0, Weight: 3}, {Dst: 1, Weight: 4}, {Dst: 3, Weight: 1}, {Dst: 4, Weight: 7}, {Dst: 5, Weight: 6}}, // 2
		{{Dst: 0, Weight: 6}, {Dst: 2, Weight: 1}, {Dst: 5, Weight: 4}},                                           // 3
		{{Dst: 1, Weight: 9}, {Dst: 2, Weight: 7}, {Dst: 5, Weight: 1}, {Dst: 6, Weight: 5}},                      // 4
		{{Dst: 2, Weight: 6}, {Dst: 3, Weight: 4}, {Dst: 4, Weight: 1}, {Dst: 6, Weight: 8}},                      // 5
		{{Dst: 4, Weight: 5}, {Dst: 5, Weight: 8}},                                                                // 6
	}

	g := graph.Init()
	for _, v := range vectors {
		g.AddAdjacentVector(v)
	}

	fmt.Println(FindPath(&g, 1, 6))
}

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
	type FromTo struct {
		src int
		dst int
	}

	type edges struct {
		edges  []graph.Edge
		weight int
	}

	tests := []struct {
		vectors  []graph.AdjacentVector
		paths    []FromTo
		shortest edges
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
			paths: []FromTo{
				{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}, {0, 6},
				{1, 0}, {1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6},
				{2, 0}, {2, 1}, {2, 2}, {2, 3}, {2, 4}, {2, 5}, {2, 6},
				{3, 0}, {3, 1}, {3, 2}, {3, 3}, {3, 4}, {3, 5}, {3, 6},
				{4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4}, {4, 5}, {4, 6},
				{5, 0}, {5, 1}, {5, 2}, {5, 3}, {5, 4}, {5, 5}, {5, 6},
				{6, 0}, {6, 1}, {6, 2}, {6, 3}, {6, 4}, {6, 5}, {6, 6},
			},
			shortest: edges{
				{
					edges{weight: 0, edges: []graph.Edge{{Weight: 0, Src: 0, Dst: 0}}},                                                             // 0 -> 0
					edges{weight: -2, edges: []graph.Edge{{Weight: -2, Src: 0, Dst: 1}}},                                                           // 0 -> 1
					edges{weight: 2, edges: []graph.Edge{{Weight: -2, Src: 0, Dst: 1}, {Weight: 8, Src: 1, Dst: 3}, {Weight: -4, Src: 3, Dst: 2}}}, // 0 -> 2
					edges{weight: 6, edges: []graph.Edge{{Weight: -2, Src: 0, Dst: 1}, {Weight: 8, Src: 1, Dst: 3}}},                               // 0 -> 3
				},
				{
					path{weight: 3, edges: []graph.Edge{{Weight: 8, Src: 1, Dst: 3}, {Weight: -4, Src: 3, Dst: 2}, {Weight: -1, Src: 2, Dst: 0}}}, // 1 -> 0
					path{weight: 0, edges: []graph.Edge{{Weight: 0, Src: 1, Dst: 1}}},                                                             // 1 -> 1
					path{weight: 4, edges: []graph.Edge{{Weight: 8, Src: 1, Dst: 3}, {Weight: -4, Src: 3, Dst: 2}}},                               // 1 -> 2
					path{weight: 8, edges: []graph.Edge{{Weight: 8, Src: 1, Dst: 3}}},                                                             // 1 -> 3
				},
				{
					path{weight: -1, edges: []graph.Edge{{Weight: -1, Src: 2, Dst: 0}}},                                                           // 2 -> 0
					path{weight: -3, edges: []graph.Edge{{Weight: -1, Src: 2, Dst: 0}, {Weight: -2, Src: 0, Dst: 1}}},                             // 2 -> 1
					path{weight: 0, edges: []graph.Edge{{Weight: 0, Src: 2, Dst: 2}}},                                                             // 2 -> 2
					path{weight: 5, edges: []graph.Edge{{Weight: -1, Src: 2, Dst: 0}, {Weight: -2, Src: 0, Dst: 1}, {Weight: 8, Src: 1, Dst: 3}}}, // 2 -> 3
				},
				{
					path{weight: -5, edges: []graph.Edge{{Weight: -4, Src: 3, Dst: 2}, {Weight: -1, Src: 2, Dst: 0}}},                               // 3 -> 0
					path{weight: -7, edges: []graph.Edge{{Weight: -4, Src: 3, Dst: 2}, {Weight: -1, Src: 2, Dst: 0}, {Weight: -2, Src: 0, Dst: 1}}}, // 3 -> 1
					path{weight: -4, edges: []graph.Edge{{Weight: -4, Src: 3, Dst: 2}}},                                                             // 3 -> 2
					path{weight: 0, edges: []graph.Edge{{Weight: 0, Src: 3, Dst: 3}}},                                                               // 3 -> 3
				},
			},
		},
	}

	t.Run("common", func(t *testing.T) {
		for _, tc := range tests {
			g := graph.Init()
			for _, v := range tc.vectors {
				g.AddAdjacentVector(v)
			}

			for _, path := range tc.paths {
				edges := FindPath(&g, path.src, path.dst)
				require.Equal(t, tc.shortest[path.src][path.dst].edges, edges)

				weight := 0
				for _, edge := range edges {
					weight += edge.Weight
				}
				require.Equal(t, tc.shortest[path.src][path.dst].weight, weight)
			}
		}
	})
}
