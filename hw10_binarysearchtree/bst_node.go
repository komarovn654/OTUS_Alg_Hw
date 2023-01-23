package hw10binarysearchtree

type bstNode struct {
	item  nodeItem
	left  *bstNode
	right *bstNode
}

type nodeItem struct {
	key   int
	value []int
}

func (bn *bstNode) insert(key int) *bstNode {
	if bn == nil {
		return &bstNode{nodeItem{key: key, value: []int{key}}, nil, nil}
	}

	if key == bn.item.key {
		bn.item.value = append(bn.item.value, key)
		return bn
	}

	if key < bn.item.key {
		bn.left = bn.left.insert(key)
		return bn
	}
	bn.right = bn.right.insert(key)
	return bn
}

func (bn *bstNode) search(x int) (item *bstNode, parent *bstNode) {
	if bn == nil {
		return nil, nil
	}

	if bn.item.key == x {
		return bn, nil
	}

	if x < bn.item.key {
		item, parent = bn.left.search(x)
		if item != nil && parent == nil {
			parent = bn
		}
		return item, parent
	}
	item, parent = bn.right.search(x)
	if item != nil && parent == nil {
		parent = bn
	}
	return item, parent
}

func (bn *bstNode) remove(x int) *bstNode {
	if item, parent := bn.search(x); item != nil {
		if item.left != nil && item.right != nil {
			item, parent = item.swapWithMax()
		}

		if parent == nil {
			return nil
		}

		if parent.left == item {
			parent.left = item.singleNodeChild()
			return bn
		}
		parent.right = item.singleNodeChild()
	}
	return bn
}

func (bn *bstNode) singleNodeChild() *bstNode {
	if bn.left != nil {
		return bn.left
	}
	if bn.right != nil {
		return bn.right
	}
	return nil
}

// swap with max on left subtree
func (bn *bstNode) swapWithMax() (node *bstNode, parent *bstNode) {
	max, maxParent := bn.left.findMax()
	if max == nil {
		max = bn.left
		maxParent = bn
	}

	bn.item = max.item
	parent = maxParent
	node = max
	return node, parent
}

func (bn *bstNode) findMax() (item *bstNode, parent *bstNode) {
	if bn == nil {
		return nil, nil
	}

	if item, parent = bn.right.findMax(); item == nil {
		return bn.right, bn
	}
	return item, parent
}

func (bn *bstNode) isValid(prev *int) bool {
	if bn == nil {
		return true
	}

	if !bn.left.isValid(prev) {
		return false
	}

	if bn.item.key <= *prev {
		return false
	}

	*prev = bn.item.key
	return bn.right.isValid(prev)
}
