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

func (b *bst) Search(x int) {

}

func (b *bst) Remove(x int) {

}

func (b *bst) IsValid() bool {
	return b.isValid2(b.root, math.MinInt)
}

func (b *bst) isValid(parent *node, prev int) bool {
	if parent == nil {
		return true
	}

	if parent.left != nil {
		if parent.left.item.key >= parent.item.key {
			return false
		}
		if !b.isValid(parent.left, parent.item.key) {
			return false
		}
	}

	if parent.right != nil {
		if parent.right.item.key <= parent.item.key {
			return false
		}
		if !b.isValid(parent.right, parent.item.key) {
			return false
		}
	}

	return true
}

func (b *bst) isValid2(parent *node, prev int) bool {
	if parent == nil {
		return true
	}

	b.isValid2(parent.left, 0)

	// if parent.left.item.key >= parent.item.key {
	// 	return false
	// }

	b.isValid2(parent.right, 0)

	if parent.right != nil && parent.right.item.key <= parent.item.key {
		return false
	}

	return true
}
