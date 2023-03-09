package bst

import (
	"fmt"
	"testing"

	"github.com/komarovn654/OTUS_Alg_Hw/hw12_othertrees/node"
)

func TestDFS(t *testing.T) {
	bst := InitBST()
	nodes := node.GenerateNodes(10, 100, 100)
	for _, n := range nodes {
		bst.Insert(n)
	}

	fmt.Println(bst.DFS())
}
