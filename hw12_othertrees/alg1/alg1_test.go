package alg1

import (
	"testing"

	"github.com/komarovn654/OTUS_Alg_Hw/hw12_othertrees/node"
)

func TestBSTAlg1(t *testing.T) {
	nodes := node.GenerateNodes(20, 100, 100)
	BSTAlg1(nodes)
}
