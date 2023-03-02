package bridges

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDFS(t *testing.T) {
	t.Run("bridges exsist", func(t *testing.T) {
		test := [][]Vertex{
			{1, 2},    // 0
			{0, 2},    // 1
			{0, 1, 3}, // 2
			{2, 4, 5}, // 3
			{3, 5},    // 4
			{3, 4},    // 5
		}
		expect := []Bridge{{2, 3}}
		expectTin := map[Vertex]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5, 5: 6}
		expectTup := map[Vertex]int{0: 1, 1: 1, 2: 1, 3: 4, 4: 4, 5: 4}

		tg := InitGraph(&test)
		bridges := tg.FindBridges()
		require.Equal(t, expect, bridges)
		require.Equal(t, tg.tin, expectTin)
		require.Equal(t, tg.tup, expectTup)
	})
	t.Run("no bridges", func(t *testing.T) {
		test := [][]Vertex{
			{1, 2},       // 0
			{0, 2},       // 1
			{0, 1, 3, 4}, // 2
			{2, 4, 5},    // 3
			{2, 3, 5},    // 4
			{3, 4},       // 5
		}
		expect := []Bridge{}
		expectTin := map[Vertex]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5, 5: 6}
		expectTup := map[Vertex]int{0: 1, 1: 1, 2: 1, 3: 3, 4: 3, 5: 4}

		tg := InitGraph(&test)
		bridges := tg.FindBridges()
		require.Equal(t, expect, bridges)
		require.Equal(t, tg.tin, expectTin)
		require.Equal(t, tg.tup, expectTup)
	})

}
