package bst

type bstNode struct {
	item  nodeItem
	left  *bstNode
	right *bstNode
}

type nodeItem struct {
	key   int
	value []Item
}

func (bn *bstNode) insert(x Item) *bstNode {
	if bn == nil {
		return &bstNode{nodeItem{key: x.GetKey(), value: []Item{x}}, nil, nil}
	}

	if x.GetKey() == bn.item.key {
		bn.item.value = append(bn.item.value, x)
		return bn
	}

	if x.GetKey() < bn.item.key {
		bn.left = bn.left.insert(x)
		return bn
	}
	bn.right = bn.right.insert(x)
	return bn
}

func (bn *bstNode) search(key int) (item *bstNode) {
	if bn == nil {
		return nil
	}

	if bn.item.key == key {
		return bn
	}

	if key < bn.item.key {
		return bn.left.search(key)
	}
	return bn.right.search(key)
}

func (bn *bstNode) remove(key int) *bstNode {
	if bn == nil {
		return nil
	}

	if key < bn.item.key {
		bn.left = bn.left.remove(key)
	}
	if key > bn.item.key {
		bn.right = bn.right.remove(key)
	}

	if bn.item.key == key {
		if bn.left != nil && bn.right != nil {
			bn = bn.swapWithLeftMax()
			bn.left = bn.left.remove(key)
			return bn
		}

		if bn.left == nil {
			return bn.right
		}
		return bn.left // return left if exist or nil
	}

	return bn
}

// swap with max on left subtree
func (bn *bstNode) swapWithLeftMax() (node *bstNode) {
	max := bn.left.findMax()

	tmp := max.item
	max.item = bn.item
	bn.item = tmp

	return bn
}

func (bn *bstNode) findMax() (max *bstNode) {
	if bn == nil {
		return nil
	}

	if max = bn.right.findMax(); max == nil {
		return bn
	}
	return max
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
