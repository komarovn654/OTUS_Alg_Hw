package boruvka

import (
	"testing"

	graph "github.com/komarovn654/OTUS_Alg_Hw/hw17_mst"
)

func TestFindMST(t *testing.T) {
	// vectors := []graph.AdjacentVector{
	// 	{{Num: 1, Weight: 1}, {Num: 3, Weight: 7}, {Num: 4, Weight: 3}},
	// 	{{Num: 0, Weight: 1}, {Num: 2, Weight: 4}, {Num: 3, Weight: 2}},
	// 	{{Num: 1, Weight: 4}, {Num: 3, Weight: 5}},
	// 	{{Num: 0, Weight: 7}, {Num: 1, Weight: 2}, {Num: 2, Weight: 5}, {Num: 4, Weight: 6}},
	// 	{{Num: 0, Weight: 3}, {Num: 3, Weight: 6}},
	// }

	vectors := []graph.AdjacentVector{
		{{Num: 1, Weight: 9}, {Num: 3, Weight: 10}, {Num: 8, Weight: 3}},                                             // 0
		{{Num: 0, Weight: 9}, {Num: 2, Weight: 4}, {Num: 4, Weight: 8}, {Num: 8, Weight: 16}},                        // 1
		{{Num: 1, Weight: 4}, {Num: 4, Weight: 14}, {Num: 5, Weight: 1}},                                             // 2
		{{Num: 0, Weight: 10}, {Num: 4, Weight: 7}, {Num: 6, Weight: 5}, {Num: 7, Weight: 13}, {Num: 8, Weight: 11}}, // 3
		{{Num: 1, Weight: 8}, {Num: 2, Weight: 14}, {Num: 3, Weight: 7}, {Num: 5, Weight: 12}, {Num: 7, Weight: 15}}, // 4
		{{Num: 2, Weight: 1}, {Num: 4, Weight: 12}, {Num: 7, Weight: 2}},                                             // 5
		{{Num: 3, Weight: 5}, {Num: 7, Weight: 6}},                                                                   // 6
		{{Num: 3, Weight: 13}, {Num: 4, Weight: 15}, {Num: 5, Weight: 2}, {Num: 6, Weight: 6}},                       // 7
		{{Num: 0, Weight: 3}, {Num: 1, Weight: 16}, {Num: 3, Weight: 11}},                                            // 8
	}

	g := graph.Init()
	for _, v := range vectors {
		g.AddAdjacentVector(v)
	}

	FindMST(&g)
}
