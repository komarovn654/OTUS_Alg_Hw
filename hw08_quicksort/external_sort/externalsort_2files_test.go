package externalsort

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExternalSort2Files(t *testing.T) {
	initTest()
	defer finTest()

	t.Run("common", func(t *testing.T) {
		path := "array_" + strconv.Itoa(20) + ".txt"

		f, _ := InitRandFile(path, 100, 100)

		f.ExternalSort2Files()
		ok, _ := f.IsSorted()
		require.True(t, ok)
	})
}

func TestExternalSortPresort(t *testing.T) {
	initTest()
	defer finTest()

	t.Run("common", func(t *testing.T) {
		path := "array_" + strconv.Itoa(20) + ".txt"

		f, _ := InitRandFile(path, 99, 10)

		f.ExternalSortPresort()
		ok, _ := f.IsSorted()
		require.True(t, ok)
	})
}
