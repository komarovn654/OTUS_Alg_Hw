package hw10binarysearchtree

var (
	validBST = bst{
		&bstNode{
			item: nodeItem{key: 49},
			left: &bstNode{
				item: nodeItem{key: 34},
				left: &bstNode{
					item: nodeItem{key: 10},
					left: nil,
					right: &bstNode{
						item: nodeItem{key: 14},
						left: &bstNode{
							item:  nodeItem{key: 12},
							left:  nil,
							right: nil,
						},
						right: &bstNode{
							item:  nodeItem{key: 15},
							left:  nil,
							right: nil,
						},
					},
				},
				right: &bstNode{
					item:  nodeItem{key: 40},
					left:  nil,
					right: nil,
				},
			},
			right: &bstNode{
				item: nodeItem{key: 75},
				left: &bstNode{
					item: nodeItem{key: 61},
					left: &bstNode{
						item:  nodeItem{key: 50},
						left:  nil,
						right: nil,
					},
					right: nil,
				},
				right: &bstNode{
					item: nodeItem{key: 93},
					left: &bstNode{
						item:  nodeItem{key: 89},
						left:  nil,
						right: nil,
					},
					right: &bstNode{
						item: nodeItem{key: 96},
						left: nil,
						right: &bstNode{
							item:  nodeItem{key: 99},
							left:  nil,
							right: nil,
						},
					},
				},
			},
		},
	}
	validBSTItems = []int{49, 75, 61, 50, 93, 89, 34, 96, 40, 10, 14, 99, 12, 15}

	invalidBST = bst{
		&bstNode{
			item: nodeItem{key: 5},
			left: &bstNode{
				item:  nodeItem{key: 4},
				left:  nil,
				right: nil,
			},
			right: &bstNode{
				item: nodeItem{key: 6},
				left: &bstNode{
					item:  nodeItem{key: 3},
					left:  nil,
					right: nil,
				},
				right: &bstNode{
					item:  nodeItem{key: 7},
					left:  nil,
					right: nil,
				},
			},
		},
	}

	UnbalanceSmallRight = avl{
		&avlNode{
			item:   nodeItem{key: 49},
			height: 4,
			left: &avlNode{
				item:   nodeItem{key: 34},
				height: 3,
				left: &avlNode{
					item:   nodeItem{key: 10},
					height: 2,
					left:   nil,
					right: &avlNode{
						item:   nodeItem{key: 14},
						height: 1,
						left:   nil,
						right:  nil,
					},
				},
				right: &avlNode{
					item:   nodeItem{key: 40},
					height: 1,
					left:   nil,
					right:  nil,
				},
			},
			right: &avlNode{
				item:   nodeItem{key: 60},
				height: 1,
				left:   nil,
				right:  nil,
			},
		},
	}
	UnbalanceBigRight = avl{
		&avlNode{
			item:   nodeItem{key: 80},
			height: 5,
			left: &avlNode{
				item:   nodeItem{key: 50},
				height: 4,
				left: &avlNode{
					item:   nodeItem{key: 40},
					height: 2,
					left: &avlNode{
						item:   nodeItem{key: 30},
						height: 1,
						left:   nil,
						right:  nil,
					},
					right: nil,
				},
				right: &avlNode{
					item:   nodeItem{key: 55},
					height: 3,
					left: &avlNode{
						item:   nodeItem{key: 52},
						height: 2,
						left: &avlNode{
							item:   nodeItem{key: 51},
							height: 1,
							left:   nil,
							right:  nil,
						},
						right: nil,
					},
					right: &avlNode{
						item:   nodeItem{key: 60},
						height: 2,
						left:   nil,
						right: &avlNode{
							item:   nodeItem{key: 70},
							height: 1,
							left:   nil,
							right:  nil,
						},
					},
				},
			},
			right: &avlNode{
				item:   nodeItem{key: 90},
				height: 1,
				left:   nil,
				right:  nil,
			},
		},
	}

	splitTree = treapNode{
		item: nodeItem{key: 3},
		left: &treapNode{
			item: nodeItem{key: 2},
			left: &treapNode{
				item:  nodeItem{key: 1},
				left:  nil,
				right: nil,
			},
			right: nil,
		},
		right: &treapNode{
			item: nodeItem{key: 9},
			left: &treapNode{
				item: nodeItem{key: 6},
				left: &treapNode{
					item: nodeItem{key: 5},
					left: &treapNode{
						item:  nodeItem{key: 4},
						left:  nil,
						right: nil,
					},
					right: nil,
				},
				right: &treapNode{
					item: nodeItem{key: 8},
					left: &treapNode{
						item:  nodeItem{key: 7},
						left:  nil,
						right: nil,
					},
					right: nil,
				},
			},
			right: &treapNode{
				item: nodeItem{key: 13},
				left: &treapNode{
					item: nodeItem{key: 11},
					left: &treapNode{
						item:  nodeItem{key: 10},
						left:  nil,
						right: nil,
					},
					right: &treapNode{
						item:  nodeItem{key: 12},
						left:  nil,
						right: nil,
					},
				},
				right: &treapNode{
					item:  nodeItem{key: 14},
					left:  nil,
					right: nil,
				},
			},
		},
	}
)
