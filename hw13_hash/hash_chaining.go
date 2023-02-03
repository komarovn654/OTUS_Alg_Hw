package hw13hash

const (
	rehashLoadFactor = 10
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

	return &newHt
}

func (ht *hashTable) Set(item tableItem) {
	if ht.items/ht.size >= rehashLoadFactor {
		*ht = *ht.rehash()
	}

	index := item.k.getHashCode() % ht.size

	if _, ok := ht.table[index].isKeyExist(item.k); ok {
		ht.table[index].replace(item.v)
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
	index := k.getHashCode() % ht.size
	ht.table[index] = ht.table[index].remove(k)
}

func (n *node) remove(k key) *node {
	for ; n != nil; n = n.next {
		if n.next != nil && n.next.item.k == k {
			n.next = n.next.next
		}
	}
	return n
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
