package demukron

import (
	"testing"

	hw15 "github.com/komarovn654/OTUS_Alg_Hw/hw15_demukron"
	"github.com/stretchr/testify/require"
)

func TestDemukron(t *testing.T) {
	t.Run("common", func(t *testing.T) {
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
		expected := []hw15.Vertices{{7}, {5}, {0, 6}, {1}, {2}, {3}, {4}}

		demGraph := InitDemukron(&test)
		require.Equal(t, expected, demGraph.Sort())
	})

	t.Run("loop", func(t *testing.T) {
		test := hw15.Graph{
			hw15.Vertices{1},          // 0
			hw15.Vertices{2},          // 1
			hw15.Vertices{3, 2},       // 2
			hw15.Vertices{4},          // 3
			hw15.Vertices{},           // 4
			hw15.Vertices{0, 1, 2, 6}, // 5
			hw15.Vertices{2, 4},       // 6
			hw15.Vertices{5},          // 7
		}

		demGraph := InitDemukron(&test)
		require.Nil(t, demGraph.Sort())
	})
}
