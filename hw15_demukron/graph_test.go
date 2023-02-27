package graph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetDegrees(t *testing.T) {
	t.Run("common", func(t *testing.T) {
		test := Graph{
			Vertices{1},          // 0
			Vertices{2},          // 1
			Vertices{3},          // 2
			Vertices{4},          // 3
			Vertices{},           // 4
			Vertices{0, 1, 2, 6}, // 5
			Vertices{2, 4},       // 6
			Vertices{5},          // 7
		}
		expected := []int{1, 2, 3, 1, 2, 1, 1, 0}
		require.Equal(t, expected, test.GetDegrees())
	})

	t.Run("loop", func(t *testing.T) {
		test := Graph{
			Vertices{1},          // 0
			Vertices{2},          // 1
			Vertices{3},          // 2
			Vertices{4, 3},       // 3
			Vertices{},           // 4
			Vertices{0, 1, 2, 6}, // 5
			Vertices{2, 4},       // 6
			Vertices{5},          // 7
		}
		require.Nil(t, test.GetDegrees())
	})
}
