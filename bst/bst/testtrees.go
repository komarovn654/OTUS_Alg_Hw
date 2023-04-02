package bst

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
)
