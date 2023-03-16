package alg2

import (
	"testing"

	node "github.com/komarovn654/OTUS_Alg_Hw/hw12_othertrees/node"
	"github.com/stretchr/testify/require"
)

func TestBSTAlg2(t *testing.T) {
	nodes := node.GenerateNodes(10, 100, 100)

	bst := BuildBST(nodes)

	require.True(t, bst.IsValid())
}
