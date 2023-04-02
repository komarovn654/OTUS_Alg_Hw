package dijkstra

import (
	"fmt"
	"testing"

	graph "github.com/komarovn654/OTUS_Alg_Hw/hw18_shortestpath"
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

	fmt.Println(FindPath(&g, 0, 0))
}
