package bst

import (
	node "github.com/komarovn654/OTUS_Alg_Hw/hw12_othertrees/node"
)

type bstNode struct {
	item  node.Node
	left  *bstNode
	right *bstNode
}

// type nodeItem struct {
// 	key   node.Node
// 	value []int
// }

func (bn *bstNode) insert(item node.Node) *bstNode {
	if bn == nil {
		return &bstNode{item, nil, nil}
	}

	if item.GetKey() == bn.item.GetKey() {
		return bn
	}

	if item.GetKey() < bn.item.GetKey() {
		bn.left = bn.left.insert(item)
		return bn
	}

	bn.right = bn.right.insert(item)
	return bn
}

// func (bn *bstNode) search(x int) (item *bstNode) {
// 	if bn == nil {
// 		return nil
// 	}

// 	if bn.item.key == x {
// 		return bn
// 	}

// 	if x < bn.item.key {
// 		return bn.left.search(x)
// 	}
// 	return bn.right.search(x)
// }

// func (bn *bstNode) remove(x int) *bstNode {
// 	if bn == nil {
// 		return nil
// 	}

// 	if x < bn.item.key {
// 		bn.left = bn.left.remove(x)
// 	}
// 	if x > bn.item.key {
// 		bn.right = bn.right.remove(x)
// 	}

// 	if bn.item.key == x {
// 		if bn.left != nil && bn.right != nil {
// 			bn = bn.swapWithLeftMax()
// 			bn.left = bn.left.remove(x)
// 			return bn
// 		}

// 		if bn.left == nil {
// 			return bn.right
// 		}
// 		return bn.left // return left if exist or nil
// 	}

// 	return bn
// }

// // swap with max on left subtree
// func (bn *bstNode) swapWithLeftMax() (node *bstNode) {
// 	max := bn.left.findMax()

// 	tmp := max.item
// 	max.item = bn.item
// 	bn.item = tmp

// 	return bn
// }

// func (bn *bstNode) findMax() (max *bstNode) {
// 	if bn == nil {
// 		return nil
// 	}

// 	if max = bn.right.findMax(); max == nil {
// 		return bn
// 	}
// 	return max
// }

// func (bn *bstNode) isValid(prev *int) bool {
// 	if bn == nil {
// 		return true
// 	}

// 	if !bn.left.isValid(prev) {
// 		return false
// 	}

// 	if bn.item.key <= *prev {
// 		return false
// 	}

// 	*prev = bn.item.key
// 	return bn.right.isValid(prev)
// }
