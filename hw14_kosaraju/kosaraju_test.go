package hw14_kosaraju

import (
	"fmt"
	"testing"
)

func TestInvert(t *testing.T) {
	g := Graph{
		Vertices{1, -1, -1},
		Vertices{2, 5, 4},
		Vertices{3, 6, -1},
		Vertices{2, 7, -1},
		Vertices{0, 5, -1},
		Vertices{6, -1, -1},
		Vertices{5, 7, -1},
		Vertices{3, -1, -1},
	}

	fmt.Println(g.reverse())
}

func TestSearch(t *testing.T) {
	g := Graph{
		Vertices{1, -1, -1},
		Vertices{2, 5, 4},
		Vertices{3, 6, -1},
		Vertices{2, 7, -1},
		Vertices{0, 5, -1},
		Vertices{6, -1, -1},
		Vertices{5, 7, -1},
		Vertices{3, -1, -1},
	}

	g.search()

}
