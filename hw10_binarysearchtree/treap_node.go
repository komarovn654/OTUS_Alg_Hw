package hw10binarysearchtree

import (
	"math/rand"
	"time"
)

type treapNode struct {
	item     nodeItem
	priority int
	left     *treapNode
	right    *treapNode
}

func merge(lTree *treapNode, rTree *treapNode) *treapNode {
	if lTree == nil {
		return rTree
	}

	if rTree == nil {
		return lTree
	}

	if lTree.priority > rTree.priority {
		return &treapNode{
			item:     lTree.item,
			priority: lTree.priority,
			left:     lTree.left,
			right:    merge(lTree.right, rTree),
		}
	}

	// lTree.priority <= rTree.priority
	return &treapNode{
		item:     rTree.item,
		priority: rTree.priority,
		left:     merge(lTree, rTree.left),
		right:    rTree.right,
	}
}

func (tn *treapNode) split(x int) (lTree *treapNode, rTree *treapNode) {
	if tn == nil {
		return nil, nil
	}

	if tn.item.key < x {
		tn.right, rTree = tn.right.split(x)
		lTree = tn
		return
	}

	lTree, tn.left = tn.left.split(x)
	rTree = tn
	return
}

func (tn *treapNode) search(x int) (item *treapNode) {
	if tn == nil {
		return nil
	}

	if tn.item.key == x {
		return tn
	}

	if x < tn.item.key {
		return tn.left.search(x)
	}
	return tn.right.search(x)
}

func newTreapNode(key int) *treapNode {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	return &treapNode{item: nodeItem{key: key}, priority: r1.Int(), left: nil, right: nil}
}

func (tn *treapNode) findMax() (max *treapNode) {
	if tn == nil {
		return nil
	}

	if max = tn.right.findMax(); max == nil {
		return tn
	}
	return max
}

func (tn *treapNode) findMin() (min *treapNode) {
	if tn == nil {
		return nil
	}

	if min = tn.right.findMin(); min == nil {
		return tn
	}
	return min
}

func (tn *treapNode) isValid(prev *int) bool {
	if tn == nil {
		return true
	}

	if !tn.left.isValid(prev) {
		return false
	}

	if tn.item.key <= *prev {
		return false
	}

	*prev = tn.item.key
	return tn.right.isValid(prev)
}
