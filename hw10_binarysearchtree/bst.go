package hw10binarysearchtree

import (
	"math"
)

type bst struct {
	root *bstNode
}

func (b *bst) Insert(x int) {
	if b.root == nil {
		b.root = b.root.insert(x)
		return
	}
	b.root.insert(x)
}

func (b *bst) Search(x int) bool {
	if item := b.root.search(x); item != nil {
		return true
	}
	return false
}

func (b *bst) Remove(x int) {
	if b.root == nil {
		return
	}
	b.root = b.root.remove(x)
}

func (b *bst) IsValid() bool {
	min := math.MinInt
	return b.root.isValid(&min)
}

func InitBST() bst {
	return bst{}
}
