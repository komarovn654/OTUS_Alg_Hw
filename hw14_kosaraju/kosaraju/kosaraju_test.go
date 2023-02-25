package kosaraju

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverse(t *testing.T) {
	t.Run("common", func(t *testing.T) {
		direct := Graph{
			Vertices{1, -1, -1},
			Vertices{2, 5, 4},
			Vertices{3, 6, -1},
			Vertices{2, 7, -1},
			Vertices{0, 5, -1},
			Vertices{6, -1, -1},
			Vertices{5, 7, -1},
			Vertices{3, -1, -1},
		}
		expected := Graph{
			Vertices{4},
			Vertices{0},
			Vertices{1, 3},
			Vertices{2, 7},
			Vertices{1},
			Vertices{1, 4, 6},
			Vertices{2, 5},
			Vertices{3, 6},
		}

		reverse := direct.reverse()

		for i := range *reverse {
			require.Equal(t, expected[i], (*reverse)[i])
		}
	})
}

func TestKosaraju(t *testing.T) {
	t.Run("common", func(t *testing.T) {
		kg := Graph{
			Vertices{1, -1, -1},
			Vertices{2, 5, 4},
			Vertices{3, 6, -1},
			Vertices{2, 7, -1},
			Vertices{0, 5, -1},
			Vertices{6, -1, -1},
			Vertices{5, -1, -1},
			Vertices{3, 6, -1},
		}
		expected := [][]Vertex{
			{5, 6},
			{2, 3, 7},
			{0, 1, 4},
		}

		res := kg.Kosaraju()

		for i, components := range res {
			res := make([]Vertex, 0)
			components.ForEach(func(a any) {
				if vertex, ok := a.(Vertex); ok {
					res = append(res, vertex)
				}
			})
			require.Equal(t, expected[i], res)
		}
	})
}
