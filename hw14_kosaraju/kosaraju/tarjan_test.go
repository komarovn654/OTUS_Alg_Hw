package kosaraju

import (
	"testing"

	hw14 "github.com/komarovn654/OTUS_Alg_Hw/hw14_kosaraju"
	"github.com/stretchr/testify/require"
)

func TestTarjanDFS(t *testing.T) {
	t.Run("directed graph", func(t *testing.T) {
		tg := kosarajuGraph{graph: Graph{
			Vertices{1, -1, -1}, // 0
			Vertices{2, 5, 4},   // 1
			Vertices{3, 6, -1},  // 2
			Vertices{2, 7, -1},  // 3
			Vertices{0, 5, -1},  // 4
			Vertices{6, -1, -1}, // 5
			Vertices{5, 7, -1},  // 6
			Vertices{3, -1, -1}, // 7
			Vertices{3, -1, -1}, // 8 unbound vertex
		},
			visited: make(tarjanVerts), vertices: hw14.New(),
		}
		expected := []Vertex{0, 1, 4, 2, 6, 5, 3, 7}
		res := tg.TarjanDFS(0)

		for i := range tg.graph {
			vertex, ok := res.Pop()
			if i == 8 {
				require.False(t, ok)
				continue
			}
			require.True(t, ok)
			require.Equal(t, expected[i], vertex)
		}

		for _, color := range tg.visited {
			require.Equal(t, BLACK, color)
		}
	})

	t.Run("empty graph", func(t *testing.T) {
		tg := kosarajuGraph{graph: Graph{}, visited: make(tarjanVerts), vertices: hw14.New()}
		require.Empty(t, tg.TarjanDFS(0))
	})

	t.Run("single node loop", func(t *testing.T) {
		tg := kosarajuGraph{graph: Graph{
			Vertices{0, -1, -1}, // 0
		},
			visited: make(tarjanVerts), vertices: hw14.New(),
		}
		res := tg.TarjanDFS(0)
		vertex, ok := res.Pop()
		require.True(t, ok)
		require.Equal(t, Vertex(0), vertex)
	})
}
