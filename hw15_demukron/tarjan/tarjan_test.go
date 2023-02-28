package tarjan

import (
	"testing"

	"golang.org/x/exp/slices"

	hw15 "github.com/komarovn654/OTUS_Alg_Hw/hw15_demukron"
	"github.com/stretchr/testify/require"
)

func TestDFS(t *testing.T) {
	t.Run("common", func(t *testing.T) {
		// Test is directed graph.
		test := hw15.Graph{
			hw15.Vertices{1},          // 0
			hw15.Vertices{2},          // 1
			hw15.Vertices{3},          // 2
			hw15.Vertices{4},          // 3
			hw15.Vertices{},           // 4
			hw15.Vertices{0, 1, 2, 6}, // 5
			hw15.Vertices{2, 4},       // 6
			hw15.Vertices{5},          // 7
		}

		tarGraph := InitTarjan(&test)
		expected := hw15.Vertices{4, 3, 2, 1, 0}
		vertices := tarGraph.DFS(hw15.Vertex(0))
		require.Equal(t, expected, vertices)

		tarGraph = InitTarjan(&test)
		expected = hw15.Vertices{4, 3, 2, 1, 0, 6, 5, 7}
		vertices = tarGraph.DFS(hw15.Vertex(7))
		require.Equal(t, expected, vertices)

		for vertex, color := range tarGraph.visited {
			slices.Contains(tarGraph.vertices, vertex)
			require.Equal(t, BLACK, color)
		}
	})
}
