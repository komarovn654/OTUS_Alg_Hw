package bst

import (
	"container/list"
	"math"

	node "github.com/komarovn654/OTUS_Alg_Hw/hw12_othertrees/node"
)

type BST struct {
	root *bstNode
}

func (b *BST) DFS() []node.Node {
	l := list.New()
	res := make([]node.Node, 0)

	l.PushBack(b.root)

	for l.Len() != 0 {
		node := l.Back()
		if bn, ok := node.Value.(bstNode); ok {
			res = append(res, bn.item)
			l.PushBack(bn.left)
			l.PushBack(bn.right)
		}
	}

	return res
}

func (b *BST) Insert(item node.Node) {
	if b.root == nil {
		b.root = b.root.insert(item)
		return
	}

	b.root.insert(item)
}

func (b *BST) Search(x node.Node) bool {
	if item := b.root.search(x); item != nil {
		return true
	}
	return false
}

func (b *BST) Remove(x node.Node) {
	if b.root == nil {
		return
	}
	b.root = b.root.remove(x)
}

func InitBST() BST {
	return BST{}
}

func (b *BST) IsValid() bool {
	min := math.MinInt
	return b.root.isValid(&min)
}
