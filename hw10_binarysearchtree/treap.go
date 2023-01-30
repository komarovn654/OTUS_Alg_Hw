package hw10binarysearchtree

import (
	"math"
)

type treap struct {
	root *treapNode
}

func (tn *treap) Insert(x int) {
	l, r := tn.root.split(x)
	l = merge(l, newTreapNode(x))
	tn.root = merge(l, r)
}

func (tn *treap) Remove(x int) {
	less, greater := tn.root.split(x)
	_, greater = greater.split(x)
	tn.root = merge(less, greater)
}

func (tn *treap) Search(x int) bool {
	if item := tn.root.search(x); item != nil {
		return true
	}
	return false
}

func (tn *treap) IsValid() bool {
	min := math.MinInt
	return tn.root.isValid(&min)
}

func InitTreap() treap {
	return treap{}
}
