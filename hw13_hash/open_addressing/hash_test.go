package openaddressing

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashInit(t *testing.T) {
	h := initHash()

	require.NotEqual(t, h.a, h.b)
}
