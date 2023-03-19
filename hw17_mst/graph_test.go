package graph

import (
	"testing"
)

func TestGraph(t *testing.T) {
	vectors := []AdjacentVector{
		{{Num: 1, Weight: 1}, {Num: 3, Weight: 7}, {Num: 4, Weight: 3}},
		{{Num: 0, Weight: 1}, {Num: 2, Weight: 4}, {Num: 3, Weight: 2}},
		{{Num: 1, Weight: 4}, {Num: 3, Weight: 5}},
		{{Num: 0, Weight: 7}, {Num: 1, Weight: 2}, {Num: 2, Weight: 5}, {Num: 4, Weight: 6}},
		{{Num: 0, Weight: 3}, {Num: 3, Weight: 6}},
	}

	g := Init()
	for _, v := range vectors {
		g.AddAdjacentVector(v)
	}

}
