package alg2

import (
	"sort"

	bst "github.com/komarovn654/OTUS_Alg_Hw/hw12_othertrees/bst"
	node "github.com/komarovn654/OTUS_Alg_Hw/hw12_othertrees/node"
)

func BuildBST(nodes []node.Node) bst.BST {
	// sort nodes by key
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].GetKey() < nodes[j].GetKey()
	})

	bst := bst.InitBST()

	buildBST(&bst, nodes, 0, len(nodes)-1)
	return bst
}

func buildBST(tree *bst.BST, nodes []node.Node, left int, right int) {
	r, pos := findRoot(nodes, left, right)
	if r == nil {
		return
	}

	tree.Insert(r)
	buildBST(tree, nodes, left, pos-1)
	buildBST(tree, nodes, pos+1, right)
}

func findRoot(nodes []node.Node, left int, right int) (root node.Node, pos int) {
	sum := 0
	leftSum := 0

	for i := left; i <= right; i++ {
		sum += nodes[i].GetWeight()
	}

	for i := left; i <= right; i++ {
		if (leftSum < sum/2) && ((leftSum + nodes[i].GetWeight()) >= sum/2) {
			return nodes[i], i
		}
		leftSum += nodes[i].GetWeight()
	}

	return nil, 0
}
