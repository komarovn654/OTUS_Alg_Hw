package main

import (
	"fmt"

	bst "github.com/komarovn654/OTUS_Alg_Hw/bst/bst"
)

type item struct {
	key   int
	value int
}

func (i *item) GetKey() int {
	return i.key
}

func main() {
	tree := bst.InitBST()

	tree.Insert(&item{key: 1, value: 14})
	tree.Insert(&item{key: 1, value: 19})
	tree.Insert(&item{key: 1, value: 4})
	tree.Insert(&item{key: 2, value: 3})
	tree.Insert(&item{key: 3, value: 43})
	tree.Insert(&item{key: 10, value: 53})
	tree.Insert(&item{key: 5, value: 16})
	tree.Insert(&item{key: 0, value: 6})

	items, ok := tree.Search(1)
	fmt.Println(ok)
	for _, item := range items {
		fmt.Println(item)
	}
	tree.Remove(1)
	items, ok = tree.Search(1)
	fmt.Println(items, ok)
}
