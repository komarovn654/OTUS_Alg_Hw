package hw14_iddfs

import "testing"

func TestIDDFS(t *testing.T) {
	test := Graph{vertices: []Vertices{
		{1, 2, 3},
		{0, 4},
		{0},
		{0, 5},
		{1},
		{3, 6, 8},
		{5, 7, 10},
		{6},
		{5, 9},
		{8},
		{6},
	},
	}

	test.DLS()
}
