package bridges

import (
	"fmt"
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

		tg := InitGraph(&test)
		bridges := tg.FindBridges()
		require.Equal(t, expect, bridges)
	})
	t.Run("no bridges", func(t *testing.T) {
		test := [][]Vertex{
			{1, 2},       // 0
			{0, 2},       // 1
			{0, 1, 3, 4}, // 2
			{2, 4, 5},    // 3
			{3, 5},       // 4
			{3, 4},       // 5
		}
		expect := []Bridge{{2, 3}}

		tg := InitGraph(&test)
		bridges := tg.FindBridges()
		fmt.Println(tg.tin)
		fmt.Println(tg.tup)
		require.Equal(t, expect, bridges)

	})

}
