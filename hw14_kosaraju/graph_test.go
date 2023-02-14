package hw14_kosaraju

import (
	"fmt"
	"testing"
)

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
