package bst

import (
	"fmt"
	"testing"

	"github.com/komarovn654/OTUS_Alg_Hw/hw12_othertrees/node"
)

func TestInsert(t *testing.T) {
	nodes := node.GenerateNodes(10, 100, 100)
	for _, n := range nodes {
		fmt.Println(n)
	}
	bst := InitBST()

	for _, n := range nodes {
		bst.Insert(n)
	}
}
