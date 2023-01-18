package hw10binarysearchtree

import (
	"fmt"
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
	ok, _, _ := b.search(b.root, x)
	return ok
}

func (b *bst) search(root *node, x int) (ok bool, item *node, parent *node) {
	if root == nil {
		return false, nil, nil
	}

	if root.item.key == x {
		return true, root, nil
	}

	if root.left != nil {
		if root.left.item.key == x {
			return true, root.left, root
		}
		ok, item, parent = b.search(root.left, x)
	}

	if root.right != nil && !ok {
		if root.right.item.key == x {
			return true, root.right, root
		}
		ok, item, parent = b.search(root.right, x)
	}
	return ok, item, parent
}

func (b *bst) Remove(x int) {
	if b.root == nil {
		return
	}

	if exist, item, parent := b.search(b.root, x); exist {
		if item.left != nil && item.right != nil {
			item, parent = b.swapWithMax(item)
		}

		if parent.left == item {
			b.removeSingleNodeRoot(parent.left, item)
			return
		}
		b.removeSingleNodeRoot(parent.right, item)
		// if item.left != nil {
		// 	parent.right = item.left
		// 	return
		// }
		// if item.right != nil {
		// 	parent.right = item.right
		// 	return
		// }
		// parent.right = nil
		// return
	}
}

func (b *bst) removeSingleNodeRoot(parent *node, item *node) {
	if item.left != nil {
		parent = item.left
		return
	}
	if item.right != nil {
		parent = item.right
		return
	}
	parent = nil
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

func (b *bst) remove(parent *node, x int) {
	if parent == nil {
		return
	}

	fmt.Println(b.search(parent, x))
}

func (b *bst) findMax(root *node) (item *node, parent *node) {
	if root == nil {
		return nil, nil
	}

	if item, parent = b.findMax(root.right); item == nil {
		// if root.right == nil {
		// 	return root, root
		// }
		return root.right, root
	}

	return item, parent
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
