package hw10binarysearchtree

import "math"

type avl struct {
	root *avlNode
}

func (a *avl) Insert(x int) {
	a.root = a.root.insert(x)
}

func (a *avl) Search(x int) bool {
	if item := a.root.search(x); item != nil {
		return true
	}
	return false
}

func (a *avl) Remove(x int) {
	if a.root == nil {
		return
	}
	a.root = a.root.remove(x)
}

func (a *avl) IsValid() bool {
	min := math.MinInt
	return a.root.isValid(&min)
}

func InitAVL() avl {
	return avl{}
}
