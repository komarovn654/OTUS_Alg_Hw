package sortutils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseTestCase(t *testing.T) {
	tc, err := ParseTestCase("testcases/random", "testcases/random/test.1.in", "testcases/random/test.1.out")
	require.NoError(t, err)
	require.Equal(t, "testcases/random", tc.Name)
	require.Equal(t, int64(10), tc.N)
	require.Equal(t, []Item{7, 0, 6, 1, 3, 2, 8, 5, 4, 9}, tc.Array)
	require.Equal(t, []Item{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, tc.Expected)

	tc, err = ParseTestCase("testcases/random", "test.1.in", "test.1.out")
	require.Error(t, err)
}
