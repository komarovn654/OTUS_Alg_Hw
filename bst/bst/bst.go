package bst

import (
	"math"
)

type Item interface {
	GetKey() int
}

type bst struct {
	root *bstNode
}

func (b *bst) Insert(x Item) {
	if b.root == nil {
		b.root = b.root.insert(x)
		return
	}
	b.root.insert(x)
}

func (b *bst) Search(key int) (items []Item, exist bool) {
	if item := b.root.search(key); item != nil {
		return item.item.value, true
	}
	return nil, false
}

func (b *bst) Remove(key int) {
	if b.root == nil {
		return
	}
	b.root = b.root.remove(key)
}

func (b *bst) IsValid() bool {
	min := math.MinInt
	return b.root.isValid(&min)
}

func InitBST() bst {
	return bst{}
}
