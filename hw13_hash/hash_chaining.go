package hw13hash

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
func (ht *hashTable) set(item tableItem) bool {
	// check for rehash

	index := item.k.getHashCode() % ht.size

	if !ht.table[index].isKeyExist(item.k) {
		newNode := node{item: item, next: ht.table[index]}
		ht.table[index] = &newNode
		return false
	}

	return true
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
	return hashTable{make([]*node, 100), 100}
}
