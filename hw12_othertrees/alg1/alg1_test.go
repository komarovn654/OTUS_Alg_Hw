package alg1

import (
	"testing"

	"github.com/komarovn654/OTUS_Alg_Hw/hw12_othertrees/node"
	"github.com/stretchr/testify/require"
)

func TestBSTAlg1(t *testing.T) {
	nodes := node.GenerateNodes(100, 100, 100)
	BuildBST(nodes)

	bst := BuildBST(nodes)

	require.True(t, bst.IsValid())
}