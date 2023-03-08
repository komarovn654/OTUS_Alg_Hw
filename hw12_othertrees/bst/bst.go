package bst

import (
	node "github.com/komarovn654/OTUS_Alg_Hw/hw12_othertrees/node"
)

type bst struct {
	root *bstNode
}

func (b *bst) Insert(item node.Node) {
	if b.root == nil {
		b.root = b.root.insert(item)
		return
	}

	b.root.insert(item)
}

// func (b *bst) Search(x int) bool {
// 	if item := b.root.search(x); item != nil {
// 		return true
// 	}
// 	return false
// }

// func (b *bst) Remove(x int) {
// 	if b.root == nil {
// 		return
// 	}
// 	b.root = b.root.remove(x)
// }

// func (b *bst) IsValid() bool {
// 	min := math.MinInt
// 	return b.root.isValid(&min)
// }

func InitBST() bst {
	return bst{}
}
