package hw10binarysearchtree

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func initAVLIncreasing(len int) (avl, []int) {
	tree := InitAVL()
	ar := make([]int, len)
	for i := 0; i < len; i++ {
		tree.Insert(i)
		ar[i] = i
	}
	return tree, ar
}

func initAVLRandom(len int) (avl, []int) {
	ar := make([]int, len)
	tree := InitAVL()
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < len; i++ {
		v := int(r1.Int31n(500))
		tree.Insert(v)
		ar[i] = v
	}
	return tree, ar
}

func TestInsertAVL(t *testing.T) {
	tree, keys := initAVLIncreasing(100)
	fmt.Println(keys)
	t.Log(tree.IsValid(), tree.root.height)
}

func TestCustomInsertAVL(t *testing.T) {
	keys := []int{321, 64, 342, 83, 75, 272, 464, 311, 97, 164}
	tree := InitAVL()

	for _, k := range keys {
		tree.Insert(k)
	}
}

func TestSmallRightRotationAVL(t *testing.T) {
	tree := UnbalanceSmallRight

	newRoot := tree.root.smallRightRotation()
	fmt.Println(newRoot)
}

func TestBigRightRotationAVL(t *testing.T) {
	tree := UnbalanceBigRight

	newRoot := tree.root.bigRightRotation()
	fmt.Println(newRoot)
}

func TestSwapWithLeftMax(t *testing.T) {
	tree := avl{
		root: &avlNode{
			item:   nodeItem{key: 321},
			height: 4,
			left: &avlNode{
				item:   nodeItem{key: 83},
				height: 3,
				left: &avlNode{
					item:   nodeItem{key: 75},
					height: 2,
					left: &avlNode{
						item:   nodeItem{key: 64},
						height: 1,
						left:   nil,
						right:  nil,
					},
					right: nil,
				},
				right: &avlNode{
					item:   nodeItem{key: 272},
					height: 3,
					left: &avlNode{
						item:   nodeItem{key: 97},
						height: 2,
						left:   nil,
						right: &avlNode{
							item:   nodeItem{key: 164},
							height: 1,
							left:   nil,
							right:  nil,
						},
					},
					right: &avlNode{
						item:   nodeItem{key: 311},
						height: 1,
						left:   nil,
						right:  nil,
					},
				},
			},
			right: &avlNode{
				item:   nodeItem{key: 342},
				height: 2,
				left:   nil,
				right: &avlNode{
					item:   nodeItem{key: 464},
					height: 1,
					left:   nil,
					right:  nil,
				},
			},
		},
	}

	// tree.root = tree.root.swapWithLeftMax()
	fmt.Println(tree.root)
}

func TestSearchAVL(t *testing.T) {
	tree := avl{
		root: &avlNode{
			item:   nodeItem{key: 321},
			height: 4,
			left: &avlNode{
				item:   nodeItem{key: 83},
				height: 3,
				left: &avlNode{
					item:   nodeItem{key: 75},
					height: 2,
					left: &avlNode{
						item:   nodeItem{key: 64},
						height: 1,
						left:   nil,
						right:  nil,
					},
					right: nil,
				},
				right: &avlNode{
					item:   nodeItem{key: 272},
					height: 3,
					left: &avlNode{
						item:   nodeItem{key: 97},
						height: 2,
						left:   nil,
						right: &avlNode{
							item:   nodeItem{key: 164},
							height: 1,
							left:   nil,
							right:  nil,
						},
					},
					right: &avlNode{
						item:   nodeItem{key: 311},
						height: 1,
						left:   nil,
						right:  nil,
					},
				},
			},
			right: &avlNode{
				item:   nodeItem{key: 342},
				height: 2,
				left:   nil,
				right: &avlNode{
					item:   nodeItem{key: 464},
					height: 1,
					left:   nil,
					right:  nil,
				},
			},
		},
	}

	for _, v := range []int{64, 75, 83, 321, 272, 97, 164, 311, 342, 464} {
		fmt.Println(tree.root.search(v))
	}
	for _, v := range []int{1, 2, 3, 4, 5, 6, 42, 7, 8, 9} {
		fmt.Println(tree.root.search(v))
	}
}

func TestRemoveAVL(t *testing.T) {
	tree := avl{
		root: &avlNode{
			item:   nodeItem{key: 321},
			height: 4,
			left: &avlNode{
				item:   nodeItem{key: 83},
				height: 3,
				left: &avlNode{
					item:   nodeItem{key: 75},
					height: 2,
					left: &avlNode{
						item:   nodeItem{key: 64},
						height: 1,
						left:   nil,
						right:  nil,
					},
					right: nil,
				},
				right: &avlNode{
					item:   nodeItem{key: 272},
					height: 3,
					left: &avlNode{
						item:   nodeItem{key: 97},
						height: 2,
						left:   nil,
						right: &avlNode{
							item:   nodeItem{key: 164},
							height: 1,
							left:   nil,
							right:  nil,
						},
					},
					right: &avlNode{
						item:   nodeItem{key: 311},
						height: 1,
						left:   nil,
						right:  nil,
					},
				},
			},
			right: &avlNode{
				item:   nodeItem{key: 342},
				height: 2,
				left:   nil,
				right: &avlNode{
					item:   nodeItem{key: 464},
					height: 1,
					left:   nil,
					right:  nil,
				},
			},
		},
	}

	tree.root = tree.root.remove(321)
}
