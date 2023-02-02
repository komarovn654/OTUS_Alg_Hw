package hw13hash

import "fmt"

const initialTableSize = 10

type key string
type value int

type hashTable struct {
	table []*node
	size  int
}

type tableItem struct {
	k key
	v value
}

type node struct {
	item tableItem
	next *node
}

func (k *key) getHashCode() (code int) {
	for _, c := range *k {
		code += int(c)
	}
	return code
}

// return true if key exists
func (ht *hashTable) Set(item tableItem) bool {
	// check for rehash

	index := item.k.getHashCode() % ht.size
	fmt.Println(index)

	if ht.table[index].isKeyExist(item.k) {
		ht.table[index].replace(item.v)
		return true
	}

	ht.table[index] = ht.table[index].add(item)
	return false
}

func (n *node) replace(v value) {
	n.item.v = v
}

func (n *node) add(item tableItem) *node {
	if n == nil {
		return &node{item: item, next: nil}
	}

	return &node{item: item, next: &node{n.item, n.next}}
}

func (n *node) isKeyExist(k key) bool {
	for ; n != nil; n = n.next {
		if n.item.k == k {
			return true
		}
	}
	return false
}

func initChainigHashTable() hashTable {
	return hashTable{make([]*node, initialTableSize), initialTableSize}
}
