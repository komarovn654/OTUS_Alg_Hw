package chaining

const (
	rehashLoadFactor = 1
)

type key string
type value int

type hashTable struct {
	table []*node
	size  int
	items int
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

func (ht *hashTable) rehash() *hashTable {
	newSize := 2*ht.size + 1
	newHt := InitChainigHashTable(newSize)

	for _, node := range ht.table {
		for ; node != nil; node = node.next {
			newHt.Set(node.item)
		}
	}

	newHt.items = ht.items
	return &newHt
}

func (ht *hashTable) Set(item tableItem) {
	if ht.items/ht.size >= rehashLoadFactor {
		*ht = *ht.rehash()
	}

	index := item.k.getHashCode() % ht.size

	if ok := ht.table[index].replaceIfExist(item); ok {
		return
	}

	ht.table[index] = ht.table[index].add(item)
	ht.items++
}

func (ht *hashTable) Get(k key) (v value, exist bool) {
	index := k.getHashCode() % ht.size

	if item, ok := ht.table[index].isKeyExist(k); ok {
		return item.v, true
	}

	return 0, false
}

func (ht *hashTable) Remove(k key) {
	ok := false
	index := k.getHashCode() % ht.size
	if ht.table[index], ok = ht.table[index].remove(k); ok {
		ht.items--
	}
}

func (n *node) remove(k key) (*node, bool) {
	if n == nil {
		return nil, false
	}

	if n.item.k == k {
		return n.next, true
	}

	for node := n; node != nil; {
		if node.next != nil && node.next.item.k == k {
			node.next = node.next.next
			return n, true
		}
		node = node.next
	}
	return n, false
}

func (n *node) replaceIfExist(item tableItem) bool {
	for ; n != nil; n = n.next {
		if n.item.k == item.k {
			n.item.v = item.v
			return true
		}
	}
	return false
}

func (n *node) add(item tableItem) *node {
	if n == nil {
		return &node{item: item, next: nil}
	}

	return &node{item: item, next: &node{n.item, n.next}}
}

func (n *node) isKeyExist(k key) (tableItem, bool) {
	for ; n != nil; n = n.next {
		if n.item.k == k {
			return n.item, true
		}
	}
	return tableItem{}, false
}

func InitChainigHashTable(size int) hashTable {
	return hashTable{
		table: make([]*node, size),
		size:  size,
		items: 0,
	}
}
