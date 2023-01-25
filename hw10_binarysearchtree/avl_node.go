package hw10binarysearchtree

import "math"

type avlNode struct {
	item   nodeItem
	height int
	left   *avlNode
	right  *avlNode
}

func (an *avlNode) insert(key int) *avlNode {
	if an == nil {
		return &avlNode{nodeItem{key: key, value: []int{key}}, 1, nil, nil}
	}

	if key == an.item.key {
		an.item.value = append(an.item.value, key)
		return an
	}

	if key < an.item.key {
		an.left = an.left.insert(key)
		an = an.rebalance()
		return an
	}
	an.right = an.right.insert(key)

	return an.rebalance()
}

func (an *avlNode) smallRightRotation() *avlNode {
	tmp := an.left.right
	res := an.left

	an.left.right = an
	an.left = tmp

	res.right.calcHeight()
	res.calcHeight()
	return res
}

func (an *avlNode) smallLeftRotation() *avlNode {
	tmp := an.right.left
	res := an.right

	an.right.left = an
	an.right = tmp

	res.left.calcHeight()
	res.calcHeight()
	return res
}

func (an *avlNode) bigLeftRotation() *avlNode {
	an.right = an.right.smallRightRotation()
	return an.smallLeftRotation()
}

func (an *avlNode) bigRightRotation() *avlNode {
	an.left = an.left.smallLeftRotation()
	return an.smallRightRotation()
}

func (an *avlNode) getHeight() int {
	if an == nil {
		return 0
	}
	return an.height
}

func (an *avlNode) calcHeight() {
	hl := 0
	hr := 0
	if an.left != nil {
		hl = an.left.height
	}
	if an.right != nil {
		hr = an.right.height
	}
	an.height = int(math.Max(float64(hl), float64(hr))) + 1
}

func (an *avlNode) rebalance() *avlNode {
	an.calcHeight()

	if an.left.getHeight()-an.right.getHeight() > 1 {
		if an.left.left.getHeight() < an.left.right.getHeight() {
			return an.bigRightRotation()
		}
		return an.smallRightRotation()
	}
	if an.right.getHeight()-an.left.getHeight() > 1 {
		if an.right.right.getHeight() < an.right.left.getHeight() {
			return an.bigLeftRotation()
		}
		return an.smallLeftRotation()
	}
	return an
}

func (an *avlNode) search(x int) *avlNode {
	if an == nil {
		return nil
	}

	if an.item.key == x {
		return an
	}

	if x < an.item.key {
		return an.left.search(x)
	}
	return an.right.search(x)
}

func (an *avlNode) remove(x int) *avlNode {

	if an.item.key == x {
		if an.left != nil && an.right != nil {
			// swap
		}
		if an.left == nil {
			return an.right
		}
		if an.right == nil {
			return an.left
		}
		return nil
	}

	if an.item.key < x {
		an.left = an.left.remove(x)
	}
	an.right = an.right.remove(x)
	return an
}

func (an *avlNode) removeMax() *avlNode {
	if an.right == nil {
		return nil
	}

	if an.right.removeMax() == nil {
		an.right = nil
	}
	return an.rebalance()
}

func (an *avlNode) swapWithLeftMax() *avlNode {
	max := an.left.findMax()

	tmp := max.item
	max.item = an.item
	an.item = tmp
	an.calcHeight()

	return max
}

func (an *avlNode) findMax() (max *avlNode) {
	if an == nil {
		return nil
	}

	if max = an.right.findMax(); max == nil {
		return an
	}
	return max
}

func (an *avlNode) isValid(prev *int) bool {
	if an == nil {
		return true
	}

	if !an.left.isValid(prev) {
		return false
	}

	if an.item.key <= *prev {
		return false
	}

	*prev = an.item.key
	return an.right.isValid(prev)
}
