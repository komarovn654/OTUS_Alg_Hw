package hw14_kosaraju

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTarjanDFS(t *testing.T) {
	t.Run("directed graph", func(t *testing.T) {
		tg := tarjanGraph{g: Graph{
			Vertices{1, -1, -1}, // 0
			Vertices{2, 5, 4},   // 1
			Vertices{3, 6, -1},  // 2
			Vertices{2, 7, -1},  // 3
			Vertices{0, 5, -1},  // 4
			Vertices{6, -1, -1}, // 5
			Vertices{5, 7, -1},  // 6
			Vertices{3, -1, -1}, // 7
			Vertices{3, -1, -1}, // 8 unbound vertex
		}}
		res := tg.DFS(0)
		for i := range tg.g {
			color, ok := res[Vertex(i)]
			if i == 8 {
				require.False(t, ok)
				continue
			}
			require.True(t, ok)
			require.Equal(t, BLACK, color)
		}
	})

	t.Run("empty graph", func(t *testing.T) {
		tg := tarjanGraph{g: Graph{}}
		require.Empty(t, tg.DFS(0))
	})

	t.Run("single node loop", func(t *testing.T) {
		tg := tarjanGraph{g: Graph{
			Vertices{0, -1, -1}, // 0
		}}
		res := tg.DFS(0)
		color, ok := res[0]
		require.True(t, ok)
		require.Equal(t, BLACK, color)
	})
}
