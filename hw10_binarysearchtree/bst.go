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
	if item, _ := b.search(b.root, x); item != nil {
		return true
	}
	return false
}

func (b *bst) search(root *node, x int) (item *node, parent *node) {
	if root == nil {
		return nil, nil
	}

	if root.item.key == x {
		return root, nil
	}

	if root.left != nil {
		if root.left.item.key == x {
			return root.left, root
		}
		item, parent = b.search(root.left, x)
	}

	if root.right != nil && item == nil {
		if root.right.item.key == x {
			return root.right, root
		}
		item, parent = b.search(root.right, x)
	}
	return item, parent
}

func (b *bst) Remove(x int) {
	if b.root == nil {
		return
	}

	if item, parent := b.search(b.root, x); item != nil {
		if item.left != nil && item.right != nil {
			item, parent = b.swapWithMax(item)
		}

		if parent == nil {
			b.root = b.singleNodeChild(b.root)
			return
		}

		if parent.left == item {
			parent.left = b.singleNodeChild(item)
			return
		}
		parent.right = b.singleNodeChild(item)
	}
}

func (b *bst) singleNodeChild(item *node) *node {
	if item.left != nil {
		return item.left

	}
	if item.right != nil {
		return item.right
	}
	return nil
}

// swap with max on left subtree
func (b *bst) swapWithMax(item *node) (sItem *node, sParent *node) {
	max, maxParent := b.findMax(item.left)
	if max == nil {
		max = item.left
		maxParent = item
	}

	item.item = max.item
	sParent = maxParent
	sItem = max
	return sItem, sParent
}

func (b *bst) findMax(root *node) (item *node, parent *node) {
	if root == nil {
		return nil, nil
	}

	if item, parent = b.findMax(root.right); item == nil {
		return root.right, root
	}

	return item, parent
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

func InitBST() bst {
	return bst{}
}
