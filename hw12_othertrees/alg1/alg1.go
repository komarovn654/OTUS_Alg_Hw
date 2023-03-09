package alg1

import (
	"sort"

	bst "github.com/komarovn654/OTUS_Alg_Hw/hw12_othertrees/bst"
	node "github.com/komarovn654/OTUS_Alg_Hw/hw12_othertrees/node"
)

func BSTAlg1(nodes []node.Node) bst.BST {
	// sort nodes by weight
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].GetWeight() > nodes[j].GetWeight()
	})

	bst := bst.InitBST()
	for _, node := range nodes {
		bst.Insert(node)
	}

	return bst
}
