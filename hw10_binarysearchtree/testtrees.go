package hw10binarysearchtree

var (
	validBST = bst{
		&node{
			item: nodeItem{key: 49},
			left: &node{
				item: nodeItem{key: 34},
				left: &node{
					item: nodeItem{key: 10},
					left: nil,
					right: &node{
						item:  nodeItem{key: 14},
						left:  nil,
						right: nil,
					},
				},
				right: &node{
					item:  nodeItem{key: 40},
					left:  nil,
					right: nil,
				},
			},
			right: &node{
				item: nodeItem{key: 75},
				left: &node{
					item: nodeItem{key: 61},
					left: &node{
						item:  nodeItem{key: 50},
						left:  nil,
						right: nil,
					},
					right: nil,
				},
				right: &node{
					item: nodeItem{key: 93},
					left: &node{
						item:  nodeItem{key: 89},
						left:  nil,
						right: nil,
					},
					right: &node{
						item: nodeItem{key: 96},
						left: nil,
						right: &node{
							item:  nodeItem{key: 99},
							left:  nil,
							right: nil,
						},
					},
				},
			},
		},
	}

	invalidBST = bst{
		&node{
			item: nodeItem{key: 5},
			left: &node{
				item:  nodeItem{key: 4},
				left:  nil,
				right: nil,
			},
			right: &node{
				item: nodeItem{key: 6},
				left: &node{
					item:  nodeItem{key: 3},
					left:  nil,
					right: nil,
				},
				right: &node{
					item:  nodeItem{key: 7},
					left:  nil,
					right: nil,
				},
			},
		},
	}
)
