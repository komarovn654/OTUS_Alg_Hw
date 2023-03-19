package kruskal

import (
	"fmt"
	"testing"

	graph "github.com/komarovn654/OTUS_Alg_Hw/hw17_mst"
)

func TestFindMST(t *testing.T) {
	vectors := []graph.AdjacentVector{
		{{Num: 1, Weight: 1}, {Num: 3, Weight: 7}, {Num: 4, Weight: 3}},
		{{Num: 0, Weight: 1}, {Num: 2, Weight: 4}, {Num: 3, Weight: 2}},
		{{Num: 1, Weight: 4}, {Num: 3, Weight: 5}},
		{{Num: 0, Weight: 7}, {Num: 1, Weight: 2}, {Num: 2, Weight: 5}, {Num: 4, Weight: 6}},
		{{Num: 0, Weight: 3}, {Num: 3, Weight: 6}},
	}

	g := graph.Init()
	for _, v := range vectors {
		g.AddAdjacentVector(v)
	}

	edges := FindMST(&g)
	for _, edge := range edges {
		fmt.Printf("%+v\n", edge)
	}
}
