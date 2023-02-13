package hw14_kosaraju

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	g := Graph{
		Vertice{2, 0, 0},
		Vertice{3, 6, 5},
		Vertice{4, 7, 0},
		Vertice{3, 8, 0},
		Vertice{1, 6, 0},
		Vertice{7, 0, 0},
		Vertice{6, 8, 0},
		Vertice{4, 0, 0},
	}

	fmt.Println(g)
}

func TestInvert(t *testing.T) {
	g := Graph{
		Vertice{1, -1, -1},
		Vertice{2, 5, 4},
		Vertice{3, 6, -1},
		Vertice{2, 7, -1},
		Vertice{0, 5, -1},
		Vertice{6, -1, -1},
		Vertice{5, 7, -1},
		Vertice{3, -1, -1},
	}

	fmt.Println(g.reverse())
}
