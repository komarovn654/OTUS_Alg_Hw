package ucs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUCS(t *testing.T) {
	//    -------0-------
	//    |      |	    |
	// (1)|   (3)|   (5)|
	//    |	     |      |
	//    1      2      3
	//    |      |	    |
	// (3)|   (5)|   (2)|
	//    |	     |      |
	//    4      5      6
	//                  |
	//               (1)|
	//                  |
	//                  7
	//    vertex : 0, 1, 2, 3, 4, 5, 6, 7
	costs := []int{0, 1, 3, 5, 4, 8, 7, 8}
	test := Graph{vertices: []Vertices{
		{Vertex{num: 1, cost: 1}, Vertex{num: 2, cost: 3}, Vertex{num: 3, cost: 5}}, // 0
		{Vertex{num: 0, cost: 1}, Vertex{num: 4, cost: 3}},                          // 1
		{Vertex{num: 0, cost: 3}, Vertex{num: 5, cost: 5}},                          // 2
		{Vertex{num: 0, cost: 5}, Vertex{num: 6, cost: 2}},                          // 3
		{Vertex{num: 1, cost: 3}},                                                   // 4
		{Vertex{num: 2, cost: 5}},                                                   // 5
		{Vertex{num: 3, cost: 2}, Vertex{num: 7, cost: 1}},                          // 6
		{Vertex{num: 6, cost: 1}},                                                   // 7
	},
	}

	t.Run("common", func(t *testing.T) {
		for vertexNum := range test.vertices {
			cost, ok := test.UCS(vertexNum)
			require.True(t, ok)
			require.Equal(t, costs[vertexNum], cost)
		}

		cost, ok := test.UCS(8)
		require.False(t, ok)
		require.Equal(t, 8, cost)
	})

}
