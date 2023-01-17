package hw10binarysearchtree

import "math"

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
	return b.search(b.root, x)
}

func (b *bst) search(parent *node, x int) bool {
	if parent == nil {
		return false
	}

	if parent.item.key == x {
		return true
	}
	if !b.search(parent.left, x) {
		return b.search(parent.right, x)
	}
	return true
}

func (b *bst) Remove(x int) {

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
