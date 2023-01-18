package hw10binarysearchtree

import (
	"math"
)

type bst struct {
	root *node
}

type node struct {
	item  nodeItem
	left  *node
	right *node
}

type nodeItem struct {
	key   int
	value []int
}

func (b *bst) Insert(x int) {
	b.insert(b.root, x)
}

func (b *bst) insert(parent *node, key int) {
	newNode := &node{nodeItem{key: key, value: []int{key}}, nil, nil}
	if parent == nil {
		b.root = newNode
		return
	}

	if key == parent.item.key {
		parent.item.value = append(parent.item.value, key)
		return
	}

	if key < parent.item.key {
		if parent.left == nil {
			parent.left = newNode
			return
		}
		b.insert(parent.left, key)
		return
	}

	if parent.right == nil {
		parent.right = newNode
		return
	}
	b.insert(parent.right, key)
}

func (b *bst) Search(x int) bool {
	ok, _ := b.search(b.root, x)
	return ok
}

func (b *bst) search(parent *node, x int) (bool, *node) {
	if parent == nil {
		return false, parent
	}

	if parent.item.key == x {
		return true, parent
	}
	ok, res := b.search(parent.left, x)
	if !ok {
		return b.search(parent.right, x)
	}
	return true, res
}

func (b *bst) Remove(x int) {
	b.remove(b.root, x)
}

func (b *bst) remove(parent *node, x int) {
	if parent == nil {
		return
	}

	if ok, rNode := b.search(b.root, x); ok {
		max := b.findMax(rNode)
		max.left = rNode.left
		max.right = rNode.right
	}
}

func (b *bst) findMax(parent *node) *node {
	if parent == nil {
		return nil
	}
	if parent.right == nil {
		return parent
	}
	return b.findMax(parent.right)
}

func (b *bst) findMin(parent *node) *node {
	if parent == nil {
		return nil
	}
	if parent.left == nil {
		return parent
	}
	return b.findMin(parent.left)
}

func (b *bst) IsValid() bool {
	return b.isValid(b.root, math.MinInt)
}

func (b *bst) isValid(parent *node, prev int) bool {
	if parent == nil {
		return true
	}

	if !b.isValid(parent.left, prev) {
		return false
	}

	if parent.item.key <= prev {
		return false
	}

	prev = parent.item.key
	return b.isValid(parent.right, prev)
}
