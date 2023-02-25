package iddfs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIDDFS(t *testing.T) {
	//    0
	//    |
	// 1--2--3
	// |     |
	// 4     5
	//       |
	//   10--6--8
	//       |  |
	//       7  9
	//       |
	//       11

	test := Graph{vertices: []Vertices{
		{1, 2, 3},
		{0, 4},
		{0},
		{0, 5},
		{1},
		{3, 6, 8},
		{5, 7, 10},
		{6, 11},
		{5, 9},
		{8},
		{6},
		{7},
	},
	}

	t.Run("common", func(t *testing.T) {
		for target := Vertex(0); target < 11; target++ {
			require.True(t, test.IDDFS(target))
		}

		require.False(t, test.IDDFS(12))
	})

}
